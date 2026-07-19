package customer

import "regexp"

var (
	nikFormatType = regexp.MustCompile(`^\d+$`)
)

type Nik struct {
	value string
}

func NewNik(nik string) (Nik, error) {
	if nik == "" {
		return Nik{}, ErrInvalidNik
	}

	if len(nik) != 16 {
		return Nik{}, ErrInvalidLength
	}

	if !nikFormatType.MatchString(nik) {
		return Nik{}, ErrInvalidType
	}

	return Nik{value: nik}, nil
}

func (n Nik) Value() string {
	return n.value
}
