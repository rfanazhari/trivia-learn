package account

type Owner struct {
	FirstName string
	LastName  string
}

func NewOwner(firstName, lastName string) (Owner, error) {
	if firstName == "" {
		return Owner{}, ErrInvalidFirstName
	}

	if len(firstName) < 3 {
		return Owner{}, ErrFirstNameTooShort
	}

	return Owner{
		FirstName: firstName,
		LastName:  lastName,
	}, nil
}

func (o Owner) GetFirstName() string {
	return o.FirstName
}

func (o Owner) GetLastName() string {
	return o.LastName
}
