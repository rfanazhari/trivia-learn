# Case 0: Customer Registration (& Auto Open Account)

> Studi kasus lanjutan — mendahului Case 1 (Buka Akun Baru) dalam urutan bisnis, tapi ditulis setelahnya karena kebutuhan model `Customer` baru muncul saat diskusi lanjutan.
>
> Bounded Context baru: `customer` — berdampingan dengan `account` yang sudah ada.

---

## 1. Requirement

### Input

| Field | Wajib? | Aturan |
|---|---|---|
| `NIK` | Wajib | 16 digit, numerik, unik (belum pernah terdaftar) |
| `FirstName` | Wajib | Minimum 3 karakter (reuse aturan dari `PersonName`, dulu `Owner`) |
| `LastName` | Opsional | Tidak ada validasi khusus |

### Proses (Satu Alur — Registrasi + Buka Akun Pertama)

1. Validasi `NIK` (format 16 digit numerik)
2. Validasi keunikan `NIK` (belum terdaftar di `CustomerRepository`)
3. Validasi `FirstName`
4. Generate `CustomerID` unik (UUID)
5. Simpan `Customer`
6. **Trigger otomatis**: panggil `OpenAccountUseCase` (reuse dari Case 1) untuk `customerID` yang baru dibuat
7. Simpan `Account` dengan `customerID` terhubung + snapshot nama pemilik

### Edge Case

- `NIK` kurang/lebih dari 16 digit
- `NIK` mengandung karakter non-digit
- `NIK` sudah terdaftar (duplicate) → `ErrDuplicateNIK`
- `FirstName` kosong / < 3 karakter → `ErrInvalidFirstName` (reuse dari Case 1)
- **Partial failure**: `Customer` berhasil disimpan, tapi `Account` gagal dibuat

### Business Rule

- `NIK` adalah identifier bisnis utama seorang `Customer` — wajib unik
- Satu `Customer` bisa punya banyak `Account` (alasan utama `Owner` naik kelas jadi Entity, lihat §5)
- **Partial failure ditoleransi**: `Customer` tetap valid meski `Account` gagal dibuat. Tidak ada rollback lintas aggregate. Retry buka akun dilakukan terpisah (di luar scope Case 0 — kemungkinan Case lanjutan: "Buka Akun Tambahan untuk Customer Existing")

---

## 2. Domain Discovery

### Bounded Context Baru: `customer`

Dipisah dari `account` karena kriteria Context Mapping (Vernon):
- **Lifecycle beda** — `Customer` hidup independen dari `Account` manapun (satu customer, banyak akun)
- **Rate of change beda** — data profil customer (nama, NIK) berubah dengan pola berbeda dari saldo/status akun
- **Bahasa beda** — "Customer mendaftar" vs "Account dibuka" adalah proses bisnis yang berbeda meski sering terjadi berurutan

### Dampak ke Bounded Context `account` (Existing)

`Account` yang sebelumnya berdiri sendiri (Case 1) sekarang **wajib** terhubung ke `Customer`:

- Tambah field `customerID` (source of truth identitas → `Customer`)
- `Owner` (VO) **di-rename** jadi `PersonName` (lebih generik, dipakai juga di `Customer`)
- `PersonName` di `Account` **tetap dipertahankan**, tapi perannya berubah jadi **snapshot histori** saat akun dibuka — bukan lagi sumber kebenaran identitas

**Preseden desain:** pola ini sama seperti `OrderItem.unitPrice` di Case Checkout — snapshot sengaja dipertahankan supaya data historis merepresentasikan kondisi *saat kejadian terjadi*, bukan kondisi terkini.

---

## 3. Domain Model

| Nama | Tipe | Alasan |
|---|---|---|
| `NIK` | Value Object | Tidak ada identitas sendiri, immutable — keunikan dicek di level use case/repository, bukan di VO |
| `PersonName` | Value Object | Rename dari `Owner`; reuse di `Customer` maupun sebagai snapshot di `Account` |
| `Customer` | Entity (Aggregate Root) | Punya `CustomerID` unik, lifecycle independen dari `Account` |
| `Account` (revisi) | Entity (Aggregate Root) | Tetap Aggregate Root sendiri, sekarang mereferensikan `Customer` via ID (bukan embed) |

### 3.1 `NIK` (Value Object)

```go
// customer/nik.go
package customer

import (
	"regexp"
)

var nikPattern = regexp.MustCompile(`^\d{16}$`)

type NIK struct {
	value string
}

func NewNIK(value string) (NIK, error) {
	if !nikPattern.MatchString(value) {
		return NIK{}, ErrInvalidNIK
	}
	return NIK{value: value}, nil
}

func (n NIK) String() string { return n.value }
```

Test mencakup: happy path (16 digit valid), kurang dari 16 digit, lebih dari 16 digit, mengandung huruf/karakter non-digit, string kosong.

### 3.2 `PersonName` (Value Object — Rename dari `Owner`)

```go
// customer/person_name.go
package customer

const minFirstNameLength = 3

type PersonName struct {
	firstName string
	lastName  string
}

func NewPersonName(firstName, lastName string) (PersonName, error) {
	if len(firstName) < minFirstNameLength {
		return PersonName{}, ErrInvalidFirstName
	}
	return PersonName{firstName: firstName, lastName: lastName}, nil
}

func (p PersonName) FirstName() string { return p.firstName }
func (p PersonName) LastName() string  { return p.lastName }
```

> Catatan migrasi: ini secara fungsional identik dengan `Owner` di Case 1. Kalau `PersonName` dipindah ke package terpisah (mis. `internal/shared` atau tetap di `internal/customer` dan diimpor oleh `account`), perlu diputuskan strategi sharing VO lintas bounded context — lihat §6.

### 3.3 `Customer` (Entity / Aggregate Root)

```go
// customer/entity.go
package customer

import "github.com/google/uuid"

type Customer struct {
	id   string
	nik  NIK
	name PersonName
}

func NewCustomer(nik NIK, name PersonName) Customer {
	return Customer{
		id:   uuid.NewString(),
		nik:  nik,
		name: name,
	}
}

func (c Customer) ID() string         { return c.id }
func (c Customer) NIK() NIK           { return c.nik }
func (c Customer) Name() PersonName   { return c.name }
```

### 3.4 `Account` (Revisi dari Case 1)

```go
// account/entity.go
package account

import "github.com/google/uuid"

type Account struct {
	id         string
	customerID string     // source of truth identitas → Customer
	ownerName  PersonName // snapshot nama saat akun dibuka, untuk histori/tampilan
	balance    Money
	status     AccountStatus
}

func NewAccount(customerID string, ownerName PersonName) (Account, error) {
	initialBalance, err := NewMoney(0)
	if err != nil {
		return Account{}, err
	}
	return Account{
		id:         uuid.NewString(),
		customerID: customerID,
		ownerName:  ownerName,
		balance:    initialBalance,
		status:     StatusPending,
	}, nil
}

func (a Account) ID() string             { return a.id }
func (a Account) CustomerID() string     { return a.customerID }
func (a Account) OwnerName() PersonName  { return a.ownerName }
func (a Account) Balance() Money         { return a.balance }
func (a Account) Status() AccountStatus  { return a.status }
```

**Breaking change dari Case 1:** `NewAccount(owner Owner)` → `NewAccount(customerID string, ownerName PersonName)`. Konsekuensi lanjutan ada di §6.

---

## 4. Use Case Layer

### 4.1 `errors.go` — Tambahan Domain Errors

```go
// customer/errors.go
package customer

import "errors"

var (
	ErrInvalidNIK       = errors.New("NIK must be exactly 16 numeric digits")
	ErrDuplicateNIK      = errors.New("NIK is already registered")
	ErrInvalidFirstName  = errors.New("first name must be at least 3 characters")
	ErrCustomerNotFound  = errors.New("customer not found")
)
```

```go
// Tambahan di account/errors.go
var ErrAccountCreationFailed = errors.New("customer registered but account creation failed")
```

### 4.2 `repository.go` — Port Baru

```go
// customer/repository.go
package customer

type CustomerRepository interface {
	Save(c Customer) error
	FindByID(id string) (Customer, error)
	FindByNIK(nik NIK) (Customer, error) // dipakai untuk cek keunikan
}
```

### 4.3 `usecase.go` — Port Baru & Revisi

```go
// customer/usecase.go
package customer

type RegisterCustomerUseCase interface {
	Execute(nikValue, firstName, lastName string) (Customer, error)
}
```

```go
// account/usecase.go — signature berubah
type OpenAccountUseCase interface {
	Execute(customerID, firstName, lastName string) (Account, error)
}
```

### 4.4 `service.go` — Orkestrasi Lintas Bounded Context

```go
// customer/service.go
package customer

var _ RegisterCustomerUseCase = (*RegisterCustomerService)(nil)

type RegisterCustomerService struct {
	repo          CustomerRepository
	openAccountUC account.OpenAccountUseCase // dependency ke context lain via interface
}

func NewRegisterCustomerService(repo CustomerRepository, openAccountUC account.OpenAccountUseCase) *RegisterCustomerService {
	return &RegisterCustomerService{repo: repo, openAccountUC: openAccountUC}
}

func (s *RegisterCustomerService) Execute(nikValue, firstName, lastName string) (Customer, error) {
	nik, err := NewNIK(nikValue)
	if err != nil {
		return Customer{}, err
	}

	if _, err := s.repo.FindByNIK(nik); err == nil {
		return Customer{}, ErrDuplicateNIK
	}

	name, err := NewPersonName(firstName, lastName)
	if err != nil {
		return Customer{}, err
	}

	cust := NewCustomer(nik, name)
	if err := s.repo.Save(cust); err != nil {
		return Customer{}, err
	}

	// Partial failure ditoleransi: Account gagal TIDAK rollback Customer
	if _, err := s.openAccountUC.Execute(cust.ID(), firstName, lastName); err != nil {
		return cust, account.ErrAccountCreationFailed
	}

	return cust, nil
}
```

> Catatan penting: `RegisterCustomerService` mengembalikan `cust` yang **valid** bersamaan dengan error `ErrAccountCreationFailed` — pemanggil (caller) harus tahu membedakan ini dari error fatal (mis. `ErrDuplicateNIK`), karena registrasi customer-nya sendiri tetap sukses.

---

## 5. Ringkasan Keputusan Desain (Trade-off yang Dibahas)

| Keputusan | Pilihan | Alasan Utama |
|---|---|---|
| `Owner`/`PersonName` di `Account` | **Tetap ada, sebagai snapshot** (bukan dihapus) | Konsisten dengan pola snapshot `OrderItem.unitPrice`; `Account` tetap self-contained untuk keperluan tampilan/histori tanpa query lintas context |
| Partial failure (Customer sukses, Account gagal) | **Customer tetap valid**, retry Account terpisah | Selaras dengan keputusan loose coupling — kalau bukan satu transaksi, aneh kalau baca data harus sinkron lintas context |
| Orchestration `RegisterCustomer` vs `OpenAccountUseCase` | **Reuse (compose)**, bukan duplikasi logic | DRY & OCP — `OpenAccountUseCase` yang sudah teruji (14 test di Case 1) dipakai ulang, bukan ditulis ulang |

---

## 6. Dampak & Migrasi ke Case 1 (Existing)

Perubahan ini **bukan penambahan murni** — ada breaking change ke kode Case 1 yang sudah selesai:

1. **`NewAccount` signature berubah** — `(owner Owner)` → `(customerID string, ownerName PersonName)`. Semua test `entity_test.go` Case 1 yang manggil `NewAccount` perlu update.
2. **`Owner` di-rename `PersonName`** dan pindah ke bertanggung jawab siapa? Dua opsi:
    - Tetap didefinisikan di `account` package, di-reuse (import) oleh `customer`
    - Pindah ke `customer` package (karena secara konseptual milik domain customer), di-reuse (import) oleh `account`

   **Belum diputuskan** — ini soal *siapa yang punya definisi VO ini*, bukan cuma soal penempatan file. Perlu didiskusikan terpisah kalau mau lanjut implementasi, karena berkaitan dengan arah dependency antar bounded context (mana yang boleh import mana).
3. **`OpenAccountUseCase.Execute` signature berubah** — nambah parameter `customerID`. `OpenAccountService` di Case 1 perlu direvisi.
4. **14 test Case 1 yang sudah PASS perlu di-review ulang**, bukan otomatis rusak, tapi assertion soal `Owner`/`NewAccount` perlu disesuaikan dengan signature baru.

---

## 7. Rencana Test (TDD — Belum Ditulis)

- `NIK`: happy path, < 16 digit, > 16 digit, non-digit, kosong (5 test)
- `PersonName`: sama seperti `Owner` di Case 1 — happy path, kosong, kurang dari minimum, boundary case (4 test)
- `Customer`: pembuatan valid, ID ter-generate unik (2 test)
- `RegisterCustomerService` (pakai fake `CustomerRepository` + fake `OpenAccountUseCase`):
    - Happy path — Customer & Account berhasil dibuat
    - NIK duplicate → `ErrDuplicateNIK`, `Save` tidak pernah dipanggil
    - NIK invalid → error sebelum sampai ke repository
    - FirstName invalid → error sebelum sampai ke repository
    - **Account gagal dibuat** → `Customer` tetap tersimpan, return `ErrAccountCreationFailed` (bukan generic error)

---

## 8. Status & Langkah Selanjutnya

**Sudah selesai (desain, belum coding):**
- [x] Domain discovery `customer` context
- [x] Requirement & edge case Case 0
- [x] Domain model (`NIK`, `PersonName`, `Customer`)
- [x] Revisi model `Account` (customerID + snapshot)
- [x] Use case layer & orkestrasi lintas context

**Belum dikerjakan:**
- [ ] Keputusan §6 poin 2 — kepemilikan definisi `PersonName` (di `account` atau `customer`)
- [ ] TDD implementasi (§7)
- [ ] Migrasi test Case 1 existing ke signature baru
- [ ] Implementasi `CustomerRepository` konkret
- [ ] Update `main.go` wiring untuk include `customer` context
- [ ] Case lanjutan: "Buka Akun Tambahan untuk Customer Existing" (menangani skenario retry dari partial failure)