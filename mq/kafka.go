package mq

import (
	"sync"

	"github.com/Fermekoo/orderin-api/utils"
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

type KafkaMQ struct {
	config   *utils.Config
	producer *kafka.Producer
	consumer *kafka.Consumer
	wgPub    *sync.WaitGroup
}

func NewKafkaMQ(config *utils.Config) MQAdapter {
	return &KafkaMQ{config: config,
		wgPub: &sync.WaitGroup{},
	}
}

func (k *KafkaMQ) Connect() error {

	// producerConfig := &kafka.ConfigMap{
	// 	"bootstrap.servers": "pkc-ew3qg.asia-southeast2.gcp.confluent.cloud:9092",
	// 	"security.protocol": "SASL_SSL",
	// 	"sasl.mechanisms":   "PLAIN",
	// 	"sasl.username":     "NQHPLMAZTNH7UWVP",
	// 	"sasl.password":     "Sgg90Q3AziGWbIJZs/KtQRoX7WY1eUyN4BMk2GbaUPDN9yn2T/PpdU0c9TBGpRla",
	// }
	producerConfig := &kafka.ConfigMap{
		"bootstrap.servers": k.config.KafkaBootsrapServers,
	}

	producer, err := kafka.NewProducer(producerConfig)
	if err != nil {
		return err
	}

	consumerConfig := &kafka.ConfigMap{
		"bootstrap.servers": k.config.KafkaBootsrapServers,
		"group.id":          k.config.KafkaGroupID,
		"auto.offset.reset": k.config.KafkaAutoOffsetReset,
	}

	consumer, err := kafka.NewConsumer(consumerConfig)
	if err != nil {
		producer.Close()
		return err
	}

	k.producer = producer
	k.consumer = consumer

	return nil
}

func (k *KafkaMQ) Disconnect() error {
	k.wgPub.Wait()
	if k.producer != nil {
		k.producer.Close()
	}

	if k.consumer != nil {
		k.consumer.Close()
	}

	return nil
}

func (k *KafkaMQ) Publish(topic string, key string, message []byte) error {

	deliveryChan := make(chan kafka.Event, 1)
	msg := &kafka.Message{
		TopicPartition: kafka.TopicPartition{
			Topic:     &topic,
			Partition: kafka.PartitionAny,
		},
		Key:   []byte(key),
		Value: message,
	}
	err := k.producer.Produce(msg, deliveryChan)
	k.producer.Flush(10 * 1000)
	// k.producer.Close()

	return err
}

// func (k *KafkaMQ) Publish(topic string, key string, message []byte) error {

// 	if k.producer == nil {
// 		return errors.New("producer is nil, cannot publish message")
// 	}

// 	errChan := make(chan error)
// 	k.wgPub.Add(1)
// 	go func() {
// 		msg := &kafka.Message{
// 			TopicPartition: kafka.TopicPartition{
// 				Topic:     &topic,
// 				Partition: kafka.PartitionAny,
// 			},
// 			Key:   []byte(key),
// 			Value: message,
// 		}
// 		err := k.producer.Produce(msg, nil)
// 		if err != nil {
// 			errChan <- err
// 		}

// 		k.producer.Flush(4 * 1000)
// 		k.wgPub.Done()
// 	}()

// 	err := <-errChan
// 	return err
// }

func (k *KafkaMQ) Subscribe(topic string, handler func(message []byte) error) error {
	return nil
}
