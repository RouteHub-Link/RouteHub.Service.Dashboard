package enums

//go:generate go run github.com/dmarkham/enumer -type=StatusState -trimprefix=StatusState -transform=snake-upper -json -sql -values -gqlgen
type StatusState int

const (
	StatusInactive StatusState = iota
	StatusActive
)

func (s StatusState) Humanize() string {
	switch s {
	case StatusInactive:
		return "Inactive"
	case StatusActive:
		return "Active"
	default:
		return "Unknown"
	}
}

func (s StatusState) HumanizeAll() []string {
	return []string{
		StatusInactive.Humanize(),
		StatusActive.Humanize(),
	}
}

func (s StatusState) HumanizeWithValue() []struct {
	Value int
	Label string
} {
	return []struct {
		Value int
		Label string
	}{
		{int(StatusInactive), StatusInactive.Humanize()},
		{int(StatusActive), StatusActive.Humanize()},
	}
}
