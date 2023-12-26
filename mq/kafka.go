package mq

import (
	"fmt"
	"sync"
	"time"

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
	return err
}

func (k *KafkaMQ) Subscribe(topic string, wg *sync.WaitGroup) error {
	defer wg.Done()
	err := k.consumer.Subscribe(topic, nil)
	if err != nil {
		return err
	}

	run := true

	for run {
		msg, err := k.consumer.ReadMessage(1 * time.Second)
		if err == nil {
			fmt.Printf("message on %s: %s\n", msg.TopicPartition, string(msg.Value))
		} else if !err.(kafka.Error).IsTimeout() {
			fmt.Printf("consumer error: %v (%v)\n", err, msg)
		}
	}
	return nil
}
