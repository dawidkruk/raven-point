package main

import (
	"fmt"
	"os"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

// BytesToString dupa
func BytesToString(data []byte) string {
	// dupa
	return string(data[:])
}

// var messagePubHandler mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
// 	message := msg.Payload()
// 	convertedMessage := BytesToString(message)
// 	fmt.Println(convertedMessage)

// 	if convertedMessage == "SOMETHING" {
// 		fmt.Println("DAMN IT WORKS!")
// 	}

// }

// var messagePubHandler mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {

// 	if BytesToString(msg.Payload()) == "SOMETHING" {
// 		fmt.Println("DAMN IT WORKS!")
// 	}

// }

var messagePubHandler mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {

	if BytesToString(msg.Payload()) == "SOMETHING" {
		fmt.Println("DAMN IT WORKS!")
	}

}

var connectHandler mqtt.OnConnectHandler = func(client mqtt.Client) {
	fmt.Println("Connected")
}

var connectLostHandler mqtt.ConnectionLostHandler = func(client mqtt.Client, err error) {
	fmt.Printf("Connect lost: %v", err)
}

func main() {
	// WHERE TO CONNECT TO
	var broker = "localhost"
	var port = 1883

	// GET THE HOSTNAME

	hostname, err := os.Hostname()
	if err == nil {
		fmt.Println("hostname:", hostname)
	}
	hostname += "-receiver-second"
	//fmt.Println("var1 = ", reflect.ValueOf(hostname).Kind())

	// DO SOME STUFF
	opts := mqtt.NewClientOptions()
	opts.AddBroker(fmt.Sprintf("tcp://%s:%d", broker, port))
	opts.SetClientID(hostname) // TEMPORARY CLIENTID
	//opts.SetUsername("dawidkruk-receiver")
	//opts.SetPassword("zxczxc")
	opts.SetDefaultPublishHandler(messagePubHandler)
	opts.OnConnect = connectHandler
	opts.OnConnectionLost = connectLostHandler
	client := mqtt.NewClient(opts)

	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	sub(client)

	// DIRTY PLAY :V
	for {
	}

}

func sub(client mqtt.Client) {
	topic := "dawidkruk/messages"
	token := client.Subscribe(topic, 1, nil)
	token.Wait()
	fmt.Printf("Subscribed to topic: %s \n", topic)
}
