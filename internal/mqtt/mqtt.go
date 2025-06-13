// filepath: internal/mqtt/mqtt.go
package mqtt

import (
	"fmt"
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
	_ = Server.AddHook(new(auth.Hook), &auth.Options{
		Ledger: &auth.Ledger{
			Auth: auth.AuthRules{
				{Username: "dji", Password: "dji", Allow: true},     // Example user
				{Username: "admin", Password: "admin", Allow: true}, // Example admin user
			},
		},
	})

	tcp := listeners.NewTCP("t1", fmt.Sprintf(":%d", config.Cfg.MQTT.TcpPort), nil)
	if err := Server.AddListener(tcp); err != nil {
		log.Fatalf("failed to add MQTT TCP listener: %v", err)
	}

	ws := listeners.NewWebsocket("w1", fmt.Sprintf(":%d", config.Cfg.MQTT.WsPort), nil)
	if err := Server.AddListener(ws); err != nil {
		log.Fatalf("failed to add MQTT WebSocket listener: %v", err)
	}

	// Start broker
	go func() {
		log.Printf("Starting embedded MQTT broker at %d", config.Cfg.MQTT.TcpPort)
		if err := Server.Serve(); err != nil {
			log.Fatalf("MQTT broker stopped: %v", err)
		}
	}()

	// 每10秒打印一次客户端数量
	// go func() {
	// 	for {
	// 		if Server != nil {
	// 			clientCount := Server.Clients.Len()
	// 			log.Printf("Current MQTT client count: %d", clientCount)
	// 		}
	// 		// Sleep for 10 seconds
	// 		time.Sleep(10 * time.Second)
	// 	}
	// }()
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
