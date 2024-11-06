package enums

//go:generate go run github.com/dmarkham/enumer -type=StatusState -trimprefix=StatusState -transform=snake-upper -json -sql -values -gqlgen
type StatusState int

const (
	StatusInactive StatusState = iota
	StatusActive
)
