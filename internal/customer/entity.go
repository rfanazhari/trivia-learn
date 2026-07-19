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

func (c *Customer) GetID() string {
	return c.id
}

func (c *Customer) GetNik() Nik {
	return c.nik
}

func (c *Customer) GetPersonName() PersonName {
	return c.personName
}
