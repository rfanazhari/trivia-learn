package customer

type Nik struct {
	value string
}

func NewNik(nik string) (*Nik, error) {
	return &Nik{value: nik}, nil
}
