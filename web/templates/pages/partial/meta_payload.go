package partial

import "RouteHub.Service.Dashboard/ent/schema/types"

type MetaPayload struct {
	Title       string `json:"meta_description_title" form:"meta_description_title"`
	FavIcon     string `json:"meta_description_favicon" form:"meta_description_favicon"`
	Description string `json:"meta_description_description" form:"meta_description_description"`

	// og:description, twitter:description
	OGDescription string `json:"meta_description_og_description" form:"meta_description_og_description"`

	// og:title, twitter:title
	OGTitle string `json:"meta_description_og_title" form:"meta_description_og_title"`

	Locale string `json:"meta_description_locale" form:"meta_description_locale"`

	TracingScript string `json:"meta_description_tracing_script" form:"meta_description_tracing_script"`
}

// https://developer.x.com/en/docs/x-for-websites/cards/overview/markup
// https://ogp.me
type MetaDescriptionPayload struct {
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
	OGCreator string `json:"meta_description_og_creator" form:"meta_description_og_creator"`
}

func (mdp *MetaDescriptionPayload) FromModel(linkMeta *types.MetaDescription) {
	if linkMeta != nil {
		mdp.Title = linkMeta.Title
		mdp.Description = linkMeta.Description
		mdp.FavIcon = linkMeta.FavIcon
		mdp.Locale = linkMeta.Locale
		mdp.OGTitle = linkMeta.OGTitle
		mdp.OGDescription = linkMeta.OGDescription
		mdp.OGURL = linkMeta.OGURL
		mdp.OGSiteName = linkMeta.OGSiteName
		mdp.OGLocale = linkMeta.OGLocale
		mdp.OGBigImage = linkMeta.OGBigImage
		mdp.OGSmallImage = linkMeta.OGSmallImage
		mdp.OGCard = linkMeta.OGCard
		mdp.OGType = linkMeta.OGType
		mdp.OGCreator = linkMeta.OGCreator
	}
}

func (mdp *MetaDescriptionPayload) UpdateModel(linkMeta *types.MetaDescription) {
	if linkMeta == nil {
		linkMeta = new(types.MetaDescription)
	}

	linkMeta.Title = mdp.Title
	linkMeta.Description = mdp.Description
	linkMeta.FavIcon = mdp.FavIcon
	linkMeta.Locale = mdp.Locale
	linkMeta.OGTitle = mdp.OGTitle
	linkMeta.OGDescription = mdp.OGDescription
	linkMeta.OGURL = mdp.OGURL
	linkMeta.OGSiteName = mdp.OGSiteName
	linkMeta.OGLocale = mdp.OGLocale
	linkMeta.OGBigImage = mdp.OGBigImage
	linkMeta.OGSmallImage = mdp.OGSmallImage
	linkMeta.OGCard = mdp.OGCard
	linkMeta.OGType = mdp.OGType
	linkMeta.OGCreator = mdp.OGCreator
}

func (mdp *MetaDescriptionPayload) AsModel() *types.MetaDescription {
	linkMeta := new(types.MetaDescription)
	mdp.UpdateModel(linkMeta)

	return linkMeta
}
