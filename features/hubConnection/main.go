package hubConnection

import "RouteHub.Service.Dashboard/features/hubConnection/mq"

func TryConnectToHub(address string) (mqc *mq.MQTTClient, err error) {
	return mq.NewMQTTClient(address, "dashboard")
}
