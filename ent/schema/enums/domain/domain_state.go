package domain

//go:generate go run github.com/dmarkham/enumer -type=DomainState -trimprefix=DomainState -transform=snake-upper -json -sql -values -gqlgen

type DomainState int

const (
	Passive DomainState = iota
	DNSStatusPending
	DNSStatusVerified
	DNSStatusFailed
	Active
)
