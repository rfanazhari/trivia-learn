package customer

type PersonName struct {
	firstName string
	lastName  string
}

func NewPersonName(firstName, lastName string) (PersonName, error) {
	if firstName == "" {
		return PersonName{}, ErrInvalidFirstName
	}

	if len(firstName) < 3 {
		return PersonName{}, ErrFirstNameTooShort
	}

	return PersonName{
		firstName: firstName,
		lastName:  lastName,
	}, nil
}

func (o PersonName) FirstName() string {
	return o.firstName
}

func (o PersonName) LastName() string {
	return o.lastName
}
