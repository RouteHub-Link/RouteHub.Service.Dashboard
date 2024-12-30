package publish

import (
	"encoding/json"

	"RouteHub.Service.Dashboard/ent"
	"RouteHub.Service.Dashboard/ent/schema/enums"
	"RouteHub.Service.Dashboard/ent/schema/types"
	"RouteHub.Service.Dashboard/features/hubConnection/mq"
)

func (p *Publisher) PublishPageSet(page *ent.Page) error {

	jsonLink, err := pageToLinkJSON(page)
	client := p.GetClient()
	if err != nil {
		return err
	}

	token := client.Publish(mq.MQE_LINK_SET.AsTopic(), 1, false, jsonLink)
	token.WaitTimeout(p.GetTimeout())

	return token.Error()
}

func (p *Publisher) PublishPageDel(page *ent.Page) error {
	client := p.GetClient()

	linkTarget := page.Slug

	token := client.Publish(mq.MQE_LINK_DEL.AsTopic(), 1, false, linkTarget)
	token.WaitTimeout(p.GetTimeout())

	return token.Error()
}

func pageToLinkJSON(page *ent.Page) (string, error) {

	link := ent.Link{
		Target: "",
		Path:   page.Slug,
		LinkContent: &types.LinkContent{
			Title:            page.Name,
			Subtitle:         page.PageDescription,
			ContentContainer: page.PageContentHTML,
			MetaDescription:  page.MetaDescription,
		},
		Status:            page.Status,
		RedirectionChoice: enums.RedirectionChoice(enums.RedirectionChoiceCustom),
		CreatedAt:         page.CreatedAt,
	}

	json, err := json.Marshal(link)
	if err != nil {
		return "", err
	}

	return string(json), nil

}
