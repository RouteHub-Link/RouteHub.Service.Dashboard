package hub

//go:generate go run github.com/dmarkham/enumer -type=RedirectionOption -trimprefix=RedirectionOption -transform=snake-upper -json -sql -values -gqlgen
type RedirectionOption int

const (
	Timed RedirectionOption = iota
	NotAutoRedirect
	DirectHTTPRedirect
	ConfirmRedirect
	Custom
)
