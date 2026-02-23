package mqtt

import (
	"fmt"
	"log"
	mqtt "github.com/eclipse/paho.mqtt.golang"
)

func NewMQTTClient(broker, port, user, password, clientID string) mqtt.Client {
	opts := mqtt.NewClientOptions()
	opts.AddBroker(fmt.Sprintf("tcp://%s:%s", broker, port))
	opts.SetClientID(clientID)
	opts.SetUsername(user)
	opts.SetPassword(password)
	opts.SetCleanSession(true)

	opts.OnConnect = func(c mqtt.Client) {
		log.Println("✅ Connected to MQTT Broker")
	}

	opts.OnConnectionLost = func(c mqtt.Client, err error) {
		log.Printf("❌ MQTT Connection lost: %v", err)
	}

	client := mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		log.Fatalf("Failed to connect MQTT: %v", token.Error())
	}

	return client
}