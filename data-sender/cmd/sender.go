package main

import (
	"os"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

func GetHostname() string {
	name, err := os.Hostname()
	if err != nil {
		panic(err)
	}
	return name
}

// MessageBuilder - Payload is still in development, placeholder used
func MessageBuilder() string {

	return "Your hostname is: " + GetHostname()
}

func Publish(client mqtt.Client, topic string) {
	for {
		token := client.Publish(topic, 0, false, MessageBuilder())
		token.Wait()
		time.Sleep(delay * time.Millisecond)
	}
}
