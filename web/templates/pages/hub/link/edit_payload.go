package link

import (
	"strings"

	"RouteHub.Service.Dashboard/ent"
	"RouteHub.Service.Dashboard/ent/schema/enums"
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

func (p *EditLinkPayload) GetRedirectionDelay() int {
	if p.RedirectionDelay == nil {
		return 10
	}

	return *p.RedirectionDelay
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
