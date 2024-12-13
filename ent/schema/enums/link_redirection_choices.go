package enums

//go:generate go run github.com/dmarkham/enumer -type=RedirectionChoice -trimprefix=RedirectionChoice -transform=snake-upper -json -sql -values -gqlgen
type RedirectionChoice int

const (
	RedirectionChoiceTimed RedirectionChoice = iota
	RedirectionChoiceNotAuto
	RedirectionChoiceDirectHTTP
	RedirectionChoiceConfirm
	RedirectionChoiceCustom
)

func (r RedirectionChoice) Humanize() string {
	switch r {
	case RedirectionChoiceTimed:
		return "Timed"
	case RedirectionChoiceNotAuto:
		return "Not Auto"
	case RedirectionChoiceDirectHTTP:
		return "Direct HTTP"
	case RedirectionChoiceConfirm:
		return "Confirm"
	case RedirectionChoiceCustom:
		return "Custom"
	default:
		return "Unknown"
	}
}

func (r RedirectionChoice) HumanizeAll() []string {
	return []string{
		RedirectionChoiceTimed.Humanize(),
		RedirectionChoiceNotAuto.Humanize(),
		RedirectionChoiceDirectHTTP.Humanize(),
		RedirectionChoiceConfirm.Humanize(),
		RedirectionChoiceCustom.Humanize(),
	}
}

func (r RedirectionChoice) HumanizeWithValue() []struct {
	Value int
	Label string
} {
	return []struct {
		Value int
		Label string
	}{
		{int(RedirectionChoiceTimed), RedirectionChoiceTimed.Humanize()},
		{int(RedirectionChoiceNotAuto), RedirectionChoiceNotAuto.Humanize()},
		{int(RedirectionChoiceDirectHTTP), RedirectionChoiceDirectHTTP.Humanize()},
		{int(RedirectionChoiceConfirm), RedirectionChoiceConfirm.Humanize()},
		{int(RedirectionChoiceCustom), RedirectionChoiceCustom.Humanize()},
	}
}
