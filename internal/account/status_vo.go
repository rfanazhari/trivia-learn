package account

type AccountStatus int

const (
	StatusPending AccountStatus = iota
	StatusActive
	StatusDormant
	StatusInactive
)

func (s AccountStatus) String() string {
	switch s {
	case StatusPending:
		return "Pending"
	case StatusActive:
		return "Active"
	case StatusDormant:
		return "Dormant"
	case StatusInactive:
		return "Inactive"
	default:
		return "Unknown"
	}
}
