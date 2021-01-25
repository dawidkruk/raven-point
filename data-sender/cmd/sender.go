package main

import (
	"fmt"
	"os"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	log "github.com/sirupsen/logrus"
)

const (
	// QosByte specifies at which level of QoS message publishing will take place
	// 0 - At most once delivery
	// 1 - At least once delivery
	// 2 - Exactly once delivery
	QosByte = 2
)

// MessageBuilder will build the payload used in the Publish function
func MessageBuilder() string {
	name, err := os.Hostname()
	if err != nil {
		panic(err)
	}
	return fmt.Sprintf("Your hostname is: %s", name)
}

// Publish will send a message to the desired topic
func Publish(client mqtt.Client, topic string) {
	for {
		token := client.Publish(topic, QosByte, false, MessageBuilder())
		log.Infof("TOPIC: %s, MESSAGE: %s", topic, MessageBuilder())
		token.Wait()
		time.Sleep(DelayInMs * time.Millisecond)
	}
}
