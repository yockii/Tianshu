package mqtt

import (
	mqttSrv "github.com/mochi-mqtt/server/v2"
	"github.com/mochi-mqtt/server/v2/packets"
	"github.com/yockii/Tianshu/pkg/config"
)

func Subscribe(topic string, callback func(msg []byte)) error {
	if config.Cfg.MQTT.UseEmbedded {
		return subscribeEmbed(topic, 0, func(cl *mqttSrv.Client, sub packets.Subscription, pk packets.Packet) {
			payload := pk.Payload
			callback(payload)
		})
	} else {
		return subscribeExternal(topic, callback)
	}
}

func subscribeExternal(_ string, _ func(msg []byte)) error {
	return nil
}

func Unsubscribe(topic string) error {
	if config.Cfg.MQTT.UseEmbedded {
		return unsubscribeEmbed(topic, 0)
	} else {
		return unsubscribeExternal(topic)
	}
}

func unsubscribeExternal(_ string) error {
	return nil
}
