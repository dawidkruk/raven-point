package main

import (
	"fmt"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	log "github.com/sirupsen/logrus"
)

const (
	// Broker specifies the IP address/FQDN of a MQTT broker
	Broker = "broker.hivemq.com"
	// Port specifies which port to connect to MQTT broker
	Port = 1883
	// DelayInMs specifies how much time Publish function will wait to send another message
	DelayInMs = 1000 // in milliseconds
)

var connectHandler mqtt.OnConnectHandler = func(client mqtt.Client) {
	log.Infof("Connected")
}

var connectLostHandler mqtt.ConnectionLostHandler = func(client mqtt.Client, err error) {
	log.Errorf("Connection lost: %v", err)
}

// CreateMqttClient create the MQTT client with additional configuration
func CreateMqttClient() mqtt.Client {
	opts := mqtt.NewClientOptions()
	opts.AddBroker(fmt.Sprintf("tcp://%s:%d", Broker, Port))
	opts.OnConnect = connectHandler
	opts.OnConnectionLost = connectLostHandler

	return mqtt.NewClient(opts)
}
