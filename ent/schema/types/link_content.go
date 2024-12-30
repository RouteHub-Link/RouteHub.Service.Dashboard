package types

type LinkContent struct {
	Title              string           `json:"title" form:"title"`
	Subtitle           string           `json:"subtitle" form:"subtitle"`
	ContentContainer   string           `json:"content_container" form:"content_container"`
	RedirectionURLText string           `json:"redirection_url_text" form:"redirection_url_text"`
	RedirectionDelay   *int             `json:"redirection_delay" form:"redirection_delay"`
	MetaDescription    *MetaDescription `json:"meta_description" form:"meta_description"`
	AdditionalHead     *string          `json:"additional_head" form:"additional_head"`
	AdditionalFooter   *string          `json:"additional_foot" form:"additional_foot"`
}

type MetaDescription struct {
	Title       string `json:"meta_description_title" form:"meta_description_title"`
	FavIcon     string `json:"meta_description_favicon" form:"meta_description_favicon"`
	Description string `json:"meta_description_description" form:"meta_description_description"`

	// og:description, twitter:description
	OGDescription string `json:"meta_description_og_description" form:"meta_description_og_description"`
	OGType        string `json:"meta_description_og_type" form:"meta_description_og_type"`

	// og:url, twitter:url
	OGURL string `json:"meta_description_og_url" form:"meta_description_og_url"`
	// og:title, twitter:title
	OGTitle string `json:"meta_description_og_title" form:"meta_description_og_title"`

	OGBigImage   string `json:"meta_description_og_big_image" form:"meta_description_og_big_image"`
	OGSmallImage string `json:"meta_description_og_small_image" form:"meta_description_og_small_image"`

	Locale     string `json:"meta_description_locale" form:"meta_description_locale"`
	OGSiteName string `json:"meta_description_og_site_name" form:"meta_description_og_site_name"`
	OGLocale   string `json:"meta_description_og_locale" form:"meta_description_og_locale"`
	// twitter:card
	OGCard string `json:"meta_description_og_card" form:"meta_description_og_card"`
	/*
		twitter:creator
		@username of content creator

		Used with summary_large_image cards
	*/
	OGCreator     string `json:"meta_description_og_creator" form:"meta_description_og_creator"`
	TracingScript string `json:"meta_description_tracing_script" form:"meta_description_tracing_script"`
}
