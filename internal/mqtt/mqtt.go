// filepath: internal/mqtt/mqtt.go
package mqtt

import (
	"log"

	mqttSrv "github.com/mochi-mqtt/server/v2"
	"github.com/mochi-mqtt/server/v2/hooks/auth"
	"github.com/mochi-mqtt/server/v2/listeners"
	"github.com/mochi-mqtt/server/v2/packets"
	"github.com/yockii/Tianshu/pkg/config"
)

var Server *mqttSrv.Server

func Start() {
	Server = mqttSrv.New(&mqttSrv.Options{
		InlineClient: true, // Enable inline client for subscribing to topics
	})
	_ = Server.AddHook(new(auth.AllowHook), nil)

	tcp := listeners.NewTCP("t1", config.Cfg.MQTT.ListenAddr, nil)
	if err := Server.AddListener(tcp); err != nil {
		log.Fatalf("failed to add MQTT TCP listener: %v", err)
	}

	// Start broker
	go func() {
		log.Printf("Starting embedded MQTT broker at %s", config.Cfg.MQTT.ListenAddr)
		if err := Server.Serve(); err != nil {
			log.Fatalf("MQTT broker stopped: %v", err)
		}
	}()
}

func Close() {
	if Server == nil {
		return
	}
	if err := Server.Close(); err != nil {
		log.Printf("Error closing MQTT server: %v", err)
	} else {
		log.Println("MQTT server closed successfully")
	}
	Server = nil
}

func subscribeEmbed(topic string, subscriptionId int, callback func(cl *mqttSrv.Client, sub packets.Subscription, pk packets.Packet)) error {
	if Server == nil {
		return mqttSrv.ErrConnectionClosed
	}
	return Server.Subscribe(topic, subscriptionId, callback)
}

func unsubscribeEmbed(topic string, subscriptionId int) error {
	if Server == nil {
		return mqttSrv.ErrConnectionClosed
	}
	return Server.Unsubscribe(topic, subscriptionId)
}
