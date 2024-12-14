package publish

import (
	"encoding/json"

	"RouteHub.Service.Dashboard/ent"
	"RouteHub.Service.Dashboard/features/hubConnection/mq"
)

func (p *Publisher) PublishPlatformSet(platform *ent.Hub) error {
	jsonPlatform, err := json.Marshal(platform)
	client := p.GetClient()
	if err != nil {
		return err
	}

	token := client.Publish(mq.MQE_PLATFORM_SET.AsTopic(), 1, false, jsonPlatform)
	token.WaitTimeout(p.GetTimeout())

	return token.Error()
}
