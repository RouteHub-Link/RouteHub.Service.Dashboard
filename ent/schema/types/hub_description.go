package types

type HubDetails struct {
	MetaDescription   MetaDescription
	NavbarDescription NavbarDescription
	FooterDescription FooterDescription
}

type NavbarDescription struct {
	BrandImg   *ImageDescription
	BrandURL   string
	BrandName  string
	Target     string
	StartItems *[]NavbarItem
	EndButtons *[]NavbarButton
}

type NavbarItem struct {
	Text     string
	URL      string
	Target   string
	Icon     string
	Dropdown *[]NavbarItem
}

type NavbarButton struct {
	Text       string
	URL        string
	Icon       string
	Target     string
	ColorClass string
}

type ImageDescription struct {
	SRC    string
	Alt    string
	Height string
	Width  string
}

type FooterDescription struct {
	ShowRouteHubBranding bool
	CompanyBrandingHTML  string
	SocialMediaContainer *SocialMediaContainer
}

type SocialMediaContainer struct {
	SocialMediaPeddingClass string
	SocialMediaColorClass   string
	SocialMediaSizeClass    string
	SocialMediaLinks        *[]ASocialMedia
}

type ASocialMedia struct {
	Icon   string `json:"icon" form:"icon" validate:"required"`
	Link   string `json:"link" form:"link" validate:"required,url"`
	Target string `json:"target" form:"target" validate:"required"`
}
