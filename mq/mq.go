package mq

type MQAdapter interface {
	Connect() error
	Disconnect() error
	Publish(topic string, key string, message []byte) error
	Subscribe(topic string, handler func(message []byte) error) error
}
