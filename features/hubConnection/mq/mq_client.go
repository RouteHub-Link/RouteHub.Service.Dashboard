package mq

import (
	"os"
	"strings"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

type MQTTClient struct {
	tcpAddr  string
	Timeout  time.Duration
	Client   *mqtt.Client
	ClientID string
}

func NewMQTTClient(tcpAddr string, clientPrefix string) (*MQTTClient, error) {
	hostName := os.Getenv("HOSTNAME")
	if hostName == "" {
		hostName = "localhost"
	}
	clientID := strings.Join([]string{clientPrefix, hostName}, "/")

	rc := MQTTClient{
		tcpAddr:  tcpAddr,
		Timeout:  30 * time.Second,
		ClientID: clientID,
	}

	options := mqtt.NewClientOptions().AddBroker(rc.tcpAddr)
	options.SetClientID(rc.ClientID)

	client := mqtt.NewClient(options)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		return nil, token.Error()
	}

	rc.Client = &client

	return &rc, nil
}
