package main

import (
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	log "github.com/sirupsen/logrus"
)

const (
	broker = "broker.hivemq.com"
	port = 1883
)

var messagePubHandler mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
	if msg.Topic() == topic {
		log.Infof("Received message on topic %s with content:\n%s", msg.Topic(), msg.Payload())
	}
}

var connectHandler mqtt.OnConnectHandler = func(client mqtt.Client) {
	fmt.Println("Connected")
}

var connectLostHandler mqtt.ConnectionLostHandler = func(client mqtt.Client, err error) {
	fmt.Printf("Connect lost: %v", err)
}

func CreateMqttClient() mqtt.Client {
	opts := mqtt.NewClientOptions()
	opts.AddBroker(fmt.Sprintf("tcp://%s:%d", broker, port))
	opts.SetDefaultPublishHandler(messagePubHandler)
	opts.OnConnect = connectHandler
	opts.OnConnectionLost = connectLostHandler

	return mqtt.NewClient(opts)
}

var subscribeHandler mqtt.MessageHandler = func(client mqtt.Client, message mqtt.Message) {
	log.Infof("Topic: %s\n", message.Topic())
	log.Info("Payload: %s\n", message.Payload())
}

func Subscribe(client mqtt.Client, topic string) {
	token := client.Subscribe(topic, 1, subscribeHandler)
	token.Wait()
	fmt.Printf("Successfully subscribed to topic: %s \n", topic)
}
