package main

import "sync"

const (
	// Topic specifies which topic will be used to send the message
	Topic = "dawidkruk/messages"
)

func main() {
	wg := &sync.WaitGroup{}
	wg.Add(1)

	go func() {
		mqttClient := CreateMqttClient()
		if token := mqttClient.Connect(); token.Wait() && token.Error() != nil {
			panic(token.Error())
		}
		Publish(mqttClient, Topic)
	}()

	wg.Wait()
}
