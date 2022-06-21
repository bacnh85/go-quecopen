package main

import (
	"flag"
	"fmt"

	MQTT "github.com/eclipse/paho.mqtt.golang"
)

func main() {
	// Broker info
	broker := flag.String("broker", "tcp://broker.hivemq.com:1883", "MQTT broker URL")
	id := flag.String("id", "b09d6290f14311ec932b0537b80d5e5f", "Client ID")
	sub_topic := flag.String("topic", "/QuecOpen/action", "Topic to subscribe to")
	pub_topic := flag.String("pub_topic", "/QuecOpen/register", "Topic to publish to")

	qos := flag.Int("qos", 0, "The Quality of Service 0,1,2 (default 0)")

	// MQTT opts
	opts := MQTT.NewClientOptions()
	opts.AddBroker(*broker)
	opts.SetClientID(*id)
	opts.SetDefaultPublishHandler(mqttReceivedCallback)

	// Create and start a client
	client := MQTT.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	// Subscribe to a topic
	if token := client.Subscribe(*sub_topic, byte(*qos), nil); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	// Publish a message
	token := client.Publish(*pub_topic, byte(*qos), false, []byte(*id))
	token.Wait()

	// Wait forever or doing anything else
	select {}
}

// mqttReceivedCallback is called when a message is received on the subscribed topic
func mqttReceivedCallback(client MQTT.Client, msg MQTT.Message) {
	fmt.Println("Received message on topic: ", msg.Topic(), " with payload: ", string(msg.Payload()))
}
