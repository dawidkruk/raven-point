package main

import "sync"

const (
	topic = "dawidkruk/messages"
)

func main() {
	wg := &sync.WaitGroup{}
	wg.Add(1)

	go func() {
		mqttClient := CreateMqttClient()
		if token := mqttClient.Connect(); token.Wait() && token.Error() != nil {
			panic(token.Error())
		}
		Subscribe(mqttClient, topic)
		Publish(mqttClient, topic)
	}()

	wg.Wait()
}
