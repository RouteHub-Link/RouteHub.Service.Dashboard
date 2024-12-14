package hubConnection

import (
	"RouteHub.Service.Dashboard/features/hubConnection/mq"
	"RouteHub.Service.Dashboard/features/hubConnection/publish"
)

func TryConnectToHub(address string) (mqc *mq.MQTTClient, err error) {
	return mq.NewMQTTClient(address, "dashboard")
}

func NewMQTTPublisher(address string) (p *publish.Publisher, err error) {
	mqc, err := TryConnectToHub(address)
	if err != nil {
		return nil, err
	}

	return publish.NewPublisher(mqc), nil
}
