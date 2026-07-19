package customer

type PersonName struct {
	FirstName string
	LastName  string
}

func NewOwner(firstName, lastName string) (PersonName, error) {
	if firstName == "" {
		return PersonName{}, ErrInvalidFirstName
	}

	if len(firstName) < 3 {
		return PersonName{}, ErrFirstNameTooShort
	}

	return PersonName{
		FirstName: firstName,
		LastName:  lastName,
	}, nil
}

func (o PersonName) GetFirstName() string {
	return o.FirstName
}

func (o PersonName) GetLastName() string {
	return o.LastName
}
