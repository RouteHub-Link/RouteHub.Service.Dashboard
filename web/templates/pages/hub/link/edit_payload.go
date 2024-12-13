package link

import (
	"strings"

	"RouteHub.Service.Dashboard/ent"
	"RouteHub.Service.Dashboard/ent/schema/enums"
	"RouteHub.Service.Dashboard/ent/schema/types"
	"RouteHub.Service.Dashboard/web/extensions"
	"RouteHub.Service.Dashboard/web/templates/pages/partial"
	"github.com/labstack/echo/v4"
)

type EditLinkPayload struct {
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
}

func (p *EditLinkPayload) Validate(c echo.Context) error {
	return extensions.BindAndValidate(c, p)
}

func (p *EditLinkPayload) FromModel(link *ent.Link) {
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
}

func (p *EditLinkPayload) UpdateModel(link *ent.Link) {
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
}

// templates\layouts\components\meta.templ Edit this file after the fields
// https://developer.x.com/en/docs/x-for-websites/cards/overview/markup
// https://ogp.me
type EditLinkMetaDescriptionPayload struct {
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

func (mdp *EditLinkMetaDescriptionPayload) FromModel(linkMeta *types.MetaDescription) {
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

func (mdp *EditLinkMetaDescriptionPayload) UpdateModel(linkMeta *types.MetaDescription) {
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

func (mdp *EditLinkMetaDescriptionPayload) AsModel() *types.MetaDescription {
	linkMeta := new(types.MetaDescription)
	mdp.UpdateModel(linkMeta)

	return linkMeta
}

func OGURLOptions(selected string, ep EditLinkPayload, hub *ent.Hub) partial.SelectOptions {
	pathLink := strings.Join([]string{hub.Edges.Domain.URL, ep.Path}, "/")
	optionsArray := []partial.SelectOption{
		{Value: "target", Label: strings.Join([]string{"Target", ep.Target}, " : ")},
		{Value: "path", Label: strings.Join([]string{"Path", pathLink}, " : ")},
	}

	if selected == "" {
		selected = "target"
	}

	options := partial.NewSelectOptions(optionsArray)
	options.Select(selected)

	return options
}

func OGTwitterCardOptions(selected string) partial.SelectOptions {
	if selected == "" {
		selected = "summary_large_image"
	}

	return partial.NewSelectOptions([]partial.SelectOption{
		{Value: "summary", Label: "Summary"},
		{Value: "summary_large_image", Label: "Summary Large Image"},
		{Value: "app", Label: "App"},
		{Value: "player", Label: "Player"},
	}).Select(selected)
}
