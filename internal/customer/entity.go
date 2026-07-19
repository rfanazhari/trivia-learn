package customer

type Customer struct {
}

func NewCustomer(nik Nik, name PersonName) (*Customer, error) {
	return &Customer{}, nil
}
