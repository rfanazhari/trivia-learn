package customer

import "github.com/google/uuid"

type Customer struct {
	id         string
	personName PersonName
	nik        Nik
}

func NewCustomer(nik Nik, name PersonName) (*Customer, error) {
	return &Customer{
		id:         uuid.New().String(),
		nik:        nik,
		personName: name,
	}, nil
}

func (c *Customer) ID() string {
	return c.id
}

func (c *Customer) Nik() Nik {
	return c.nik
}

func (c *Customer) Name() PersonName {
	return c.personName
}
