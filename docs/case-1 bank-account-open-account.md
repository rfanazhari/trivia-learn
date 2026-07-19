# Case 1: Buka Akun Baru (Open Account)

## Deskripsi Flow

User (nasabah/sistem) ingin membuka akun bank baru dengan memberikan data pemilik
akun. Sistem akan membuat akun baru dengan saldo awal 0 dan status belum aktif.

## Input

| Field       | Wajib?   | Aturan                              |
|-------------|----------|--------------------------------------|
| `FirstName` | Wajib    | Minimum 3 karakter, tidak boleh kosong |
| `LastName`  | Opsional | Tidak ada validasi khusus            |

## Proses

1. Validasi `FirstName` (tidak kosong, minimum 3 karakter)
2. Generate `AccountID` unik (UUID)
3. Set `Balance` awal = 0
4. Set `Status` = `Pending`

## Output (Happy Path)

Akun baru berhasil dibuat dengan:
- `ID` — unik per akun
- `FirstName`, `LastName` — sesuai input
- `Balance` = 0
- `Status` = `Pending`

## Edge Case / Error

| Kondisi                        | Perilaku                          |
|---------------------------------|------------------------------------|
| `FirstName` kosong              | Ditolak — `ErrInvalidFirstName`    |
| `FirstName` < 3 karakter        | Ditolak — `ErrInvalidFirstName`    |
| `FirstName` tepat 3 karakter    | Diterima (boundary case)           |

## Business Rules / Invariant

- Akun baru **selalu** dibuat dengan saldo ≥ 0 (di Case 1 selalu tepat 0)
- Setiap akun **harus** memiliki `ID` unik begitu dibuat
- Akun dengan saldo 0 → status awal selalu `Pending` (bukan `Active`)
- Status `Pending` hanya berubah menjadi `Active` melalui setoran pertama —
  **di luar scope Case 1**, akan diimplementasikan di Case 2 (Deposit)

## Ubiquitous Language

| Istilah         | Arti                                                              |
|------------------|--------------------------------------------------------------------|
| `Account`        | Entity/Aggregate Root — representasi akun bank                    |
| `Owner`          | Value Object — pemilik akun (`FirstName`, `LastName`)              |
| `Money`/`Balance`| Value Object — saldo akun, non-negatif                            |
| `AccountStatus`  | Value Object/enum — state akun: `Pending`, `Active`, `Dormant`, `Inactive` |
| `OpenAccount`    | Use case — operasi membuka akun baru                              |

## Status Akun (State Machine — didefinisikan untuk seluruh domain)

| Status      | Arti                                                        |
|-------------|---------------------------------------------------------------|
| `Pending`   | Akun baru dibuat, saldo masih 0, belum ada transaksi          |
| `Active`    | Sudah pernah setor (deposit pertama), bisa transaksi normal    |
| `Dormant`   | Akun aktif tapi tidak ada aktivitas dalam periode tertentu      |
| `Inactive`  | Akun ditutup/dinonaktifkan (oleh user atau admin)              |

Transisi status:
```
Pending --(deposit pertama)--> Active
Active --(tidak ada aktivitas X lama)--> Dormant
Dormant --(ada transaksi lagi)--> Active
Active/Dormant --(tutup akun, oleh user atau admin)--> Inactive
```

> **Catatan scope Case 1:** hanya status `Pending` yang diimplementasikan.
> Transisi lain didefinisikan sebagai bagian dari domain model (enum), tapi
> logic transisinya menyusul di case-case berikutnya:
> - `Pending → Active`: Case 2 (Deposit)
> - `Active/Dormant → Inactive`: Case terpisah (Tutup Akun) — dipicu user atau admin
> - `Active → Dormant`: **di-skip**, tidak ada scheduler untuk versi ini

## Domain Model

| Nama            | Tipe                     | Alasan                                          |
|-----------------|---------------------------|--------------------------------------------------|
| `Owner`         | Value Object               | Tidak ada ID, dibandingkan by value               |
| `Money`         | Value Object               | Tidak ada ID, immutable, dibandingkan by value    |
| `AccountStatus` | Value Object / enum        | Representasi state, tanpa identitas               |
| `Account`       | Entity (Aggregate Root)    | Punya `ID` unik, state (`Balance`, `Status`) berubah seiring waktu |

## Struktur Kode (implementasi)

Pola: **Package-by-feature**

```
internal/account/
├── entity.go       # Account (Aggregate Root)
├── entity_test.go
├── owner.go        # Value Object: Owner
├── owner_test.go
├── money.go        # Value Object: Money
├── money_test.go
├── status.go        # Value Object: AccountStatus
├── status_test.go
└── errors.go        # domain errors
```

## Status Implementasi

- [x] `Owner` (Value Object) — TDD selesai, 5 test
- [x] `Money` (Value Object) — TDD selesai, 3 test
- [x] `AccountStatus` (Value Object) — TDD selesai, 1 test
- [x] `Account` (Entity/Aggregate Root) — TDD selesai, 2 test
- [ ] `AccountRepository` (interface) — belum
- [ ] `OpenAccountUseCase` (interface + implementasi) — belum
- [ ] Wiring (`main.go`) — belum

**Total test saat ini: 11, semua PASS.**