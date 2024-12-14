package publish

import (
	"encoding/json"

	"RouteHub.Service.Dashboard/ent"
	"RouteHub.Service.Dashboard/features/hubConnection/mq"
)

func (p *Publisher) PublishLinkSet(link *ent.Link) error {
	jsonLink, err := json.Marshal(link)
	client := p.GetClient()
	if err != nil {
		return err
	}

	token := client.Publish(mq.MQE_LINK_SET.AsTopic(), 1, false, jsonLink)
	token.WaitTimeout(p.GetTimeout())

	return token.Error()
}

func (p *Publisher) PublishLinkDel(link *ent.Link) error {
	client := p.GetClient()

	linkTarget := link.Path

	token := client.Publish(mq.MQE_LINK_DEL.AsTopic(), 1, false, linkTarget)
	token.WaitTimeout(p.GetTimeout())

	return token.Error()
}
