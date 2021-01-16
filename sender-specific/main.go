package main

import (
	"fmt"
	"os"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

var messagePubHandler mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
	fmt.Printf("TOPIC: %s MESSAGE: %s \n", msg.Topic(), msg.Payload())
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
	hostname += "-sender-second"
	//fmt.Println("var1 = ", reflect.ValueOf(hostname).Kind())

	// DO SOME STUFF
	opts := mqtt.NewClientOptions()
	opts.AddBroker(fmt.Sprintf("tcp://%s:%d", broker, port))
	opts.SetClientID(hostname) // TEMPORARY CLIENTID
	//opts.SetUsername("dawidkruk-sender")
	//opts.SetPassword("zxczxc")
	opts.SetDefaultPublishHandler(messagePubHandler)
	opts.OnConnect = connectHandler
	opts.OnConnectionLost = connectLostHandler
	client := mqtt.NewClient(opts)

	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	sub(client)
	publish(client, hostname)

	// DIRTY PLAY :V
	for {
	}

}

// func publish(client mqtt.Client, hostname string) {
// 	i := 0
// 	for {
// 		text := fmt.Sprintf("Hello there #%d, USERNAME: %s", i, hostname)
// 		token := client.Publish("dawidkruk/messages", 0, false, text)
// 		token.Wait()
// 		time.Sleep(500 * time.Millisecond)
// 		i++
// 	}
// }

func publish(client mqtt.Client, hostname string) {
	i := 0
	for {
		text := fmt.Sprintf("SOMETHING")
		token := client.Publish("dawidkruk/messages", 0, false, text)
		token.Wait()
		time.Sleep(500 * time.Millisecond)
		i++
	}
}

func sub(client mqtt.Client) {
	topic := "dawidkruk/messages"
	token := client.Subscribe(topic, 1, nil)
	token.Wait()
	fmt.Printf("Subscribed to topic: %s \n", topic)
}

