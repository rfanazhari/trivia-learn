package customer

import "regexp"

var (
	nikFormatLength = regexp.MustCompile(`^\d{16}$`)
	nikFormatType   = regexp.MustCompile(`^\d+$`)
)

type Nik struct {
	value string
}

func NewNik(nik string) (Nik, error) {
	if nik == "" {
		return Nik{}, ErrInvalidNik
	}

	if !nikFormatLength.MatchString(nik) {
		return Nik{}, ErrInvalidLength
	}

	if !nikFormatType.MatchString(nik) {
		return Nik{}, ErrInvalidType
	}

	return Nik{value: nik}, nil
}

func (n Nik) String() string {
	return n.value
}
