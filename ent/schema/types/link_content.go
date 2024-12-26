package types

type LinkContent struct {
	Title              string
	Subtitle           string
	ContentContainer   string
	RedirectionURLText string
	RedirectionDelay   *int
	MetaDescription    *MetaDescription
	AdditionalHead     *string
	AdditionalFooter   *string
}

type MetaDescription struct {
	Title         string
	FavIcon       string
	Description   string
	Locale        string
	OGTitle       string
	OGDescription string
	OGURL         string
	OGSiteName    string
	OGLocale      string
	OGBigImage    string
	OGBigWidth    string
	OGBigHeight   string
	OGSmallImage  string
	OGSmallWidth  string
	OGSmallHeight string
	OGCard        string
	OGType        string
	OGCreator     string
	TracingScript string
}
