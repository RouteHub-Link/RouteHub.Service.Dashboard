package link

import (
	"RouteHub.Service.Dashboard/ent"
	"RouteHub.Service.Dashboard/ent/schema/enums"
	"RouteHub.Service.Dashboard/ent/schema/types"
	"RouteHub.Service.Dashboard/web/extensions"
	"github.com/labstack/echo/v4"
)

type EditPayload struct {
	Target             string                  `json:"target_address" form:"target_address" validate:"required,url"`
	Path               string                  `json:"path" form:"path" validate:"required"`
	Title              string                  `json:"title" form:"title"`
	Subtitle           string                  `json:"subtitle" form:"subtitle"`
	ContentContainer   string                  `json:"content_container" form:"content_container"`
	RedirectionChoice  enums.RedirectionChoice `json:"redirection_choice" form:"redirection_choice"`
	RedirectionURLText string                  `json:"redirection_url_text" form:"redirection_url_text"`
	RedirectionDelay   *int                    `json:"redirection_delay" form:"redirection_delay"`
	AdditionalHead     *string                 `json:"additional_head" form:"additional_head"`
	AdditionalFooter   *string                 `json:"additional_foot" form:"additional_foot"`
	MetaDescription    MetaDescriptionPayload  `json:"meta_description" form:"meta_description"`
}

func (p *EditPayload) Validate(c echo.Context) error {
	return extensions.BindAndValidate(c, p)
}

func (p *EditPayload) FromModel(link *ent.Link) {
	p.Target = link.Target
	p.Path = link.Path
	p.Title = link.LinkContent.Title
	p.Subtitle = link.LinkContent.Subtitle

	p.ContentContainer = link.LinkContent.ContentContainer
	p.RedirectionChoice = link.RedirectionChoice
	p.RedirectionURLText = link.LinkContent.RedirectionURLText
	p.RedirectionDelay = link.LinkContent.RedirectionDelay

	p.AdditionalHead = link.LinkContent.AdditionalHead
	p.AdditionalFooter = link.LinkContent.AdditionalFooter

	linkMeta := link.LinkContent.MetaDescription
	if linkMeta != nil {
		p.MetaDescription.FromModel(linkMeta)
	}
}

func (p *EditPayload) UpdateModel(link *ent.Link) {
	link.Target = p.Target
	link.Path = p.Path

	link.LinkContent.Title = p.Title
	link.LinkContent.Subtitle = p.Subtitle

	link.LinkContent.ContentContainer = p.ContentContainer
	link.RedirectionChoice = p.RedirectionChoice
	link.LinkContent.RedirectionURLText = p.RedirectionURLText
	link.LinkContent.RedirectionDelay = p.RedirectionDelay

	link.LinkContent.AdditionalHead = p.AdditionalHead
	link.LinkContent.AdditionalFooter = p.AdditionalFooter

	p.MetaDescription.UpdateModel(link.LinkContent.MetaDescription)
}

type MetaDescriptionPayload struct {
	Title         string `json:"title" form:"title"`
	Description   string `json:"description" form:"description"`
	FavIcon       string `json:"favicon" form:"favicon"`
	Locale        string `json:"locale" form:"locale"`
	OGTitle       string `json:"og_title" form:"og_title"`
	OGDescription string `json:"og_description" form:"og_description"`
	OGURL         string `json:"og_url" form:"og_url"`
	OGSiteName    string `json:"og_site_name" form:"og_site_name"`
	OGMetaType    string `json:"og_meta_type" form:"og_meta_type"`
	OGLocale      string `json:"og_locale" form:"og_locale"`
	OGBigImage    string `json:"og_big_image" form:"og_big_image"`
	OGSmallImage  string `json:"og_small_image" form:"og_small_image"`
	OGCard        string `json:"og_card" form:"og_card"`
	OGSite        string `json:"og_site" form:"og_site"`
	OGType        string `json:"og_type" form:"og_type"`
	OGCreator     string `json:"og_creator" form:"og_creator"`
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
		mdp.OGMetaType = linkMeta.OGMetaType
		mdp.OGLocale = linkMeta.OGLocale
		mdp.OGBigImage = linkMeta.OGBigImage
		mdp.OGSmallImage = linkMeta.OGSmallImage
		mdp.OGCard = linkMeta.OGCard
		mdp.OGSite = linkMeta.OGSite
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
	linkMeta.OGMetaType = mdp.OGMetaType
	linkMeta.OGLocale = mdp.OGLocale
	linkMeta.OGBigImage = mdp.OGBigImage
	linkMeta.OGSmallImage = mdp.OGSmallImage
	linkMeta.OGCard = mdp.OGCard
	linkMeta.OGSite = mdp.OGSite
	linkMeta.OGType = mdp.OGType
	linkMeta.OGCreator = mdp.OGCreator
}

func (mdp *MetaDescriptionPayload) AsModel() *types.MetaDescription {
	linkMeta := new(types.MetaDescription)
	mdp.UpdateModel(linkMeta)

	return linkMeta
}
