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

func (o PersonName) GetFirstName() string {
	return o.firstName
}

func (o PersonName) GetLastName() string {
	return o.lastName
}
