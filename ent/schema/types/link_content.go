package types

import "encoding/json"

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

func (lc LinkContent) UnmarshalJSON(data []byte) error {
	type Alias LinkContent
	aux := &struct {
		*Alias
	}{
		Alias: (*Alias)(&lc),
	}
	err := json.Unmarshal(data, &aux)

	return err
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
}

func (og *MetaDescription) ParseFromJSON(jsonOpenGraph string) {
	json.Unmarshal([]byte(jsonOpenGraph), &og)
}

func (og *MetaDescription) AsJSON() (string, error) {
	jsonOpengraph, err := json.Marshal(og)
	if err != nil {
		return "", err
	}
	return string(jsonOpengraph), nil
}
