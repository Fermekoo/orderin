package mq

import "sync"

type MQAdapter interface {
	Connect() error
	Disconnect() error
	Publish(topic string, key string, message []byte) error
	Subscribe(topic string, wg *sync.WaitGroup) error
}
