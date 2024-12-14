package publish

import (
	"time"

	"RouteHub.Service.Dashboard/features/hubConnection/mq"
	mqtt "github.com/eclipse/paho.mqtt.golang"
)

type Publisher struct {
	client *mq.MQTTClient
}

func NewPublisher(client *mq.MQTTClient) *Publisher {
	return &Publisher{
		client: client,
	}
}

func (p *Publisher) GetClient() mqtt.Client {
	return *p.client.Client
}

func (p *Publisher) GetTimeout() time.Duration {
	return p.client.Timeout
}
