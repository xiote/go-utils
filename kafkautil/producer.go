package kafkautil

import (
	"fmt"
	"gopkg.in/confluentinc/confluent-kafka-go.v1/kafka"
)

type KafkaProducer interface {
	Events() chan kafka.Event
	ProduceChannel() chan *kafka.Message
	Close()
}

type Producer struct {
	OutChan chan string
	KafkaProducer
	Topic      string
	MessageKey string
}

func (p *Producer) Produce() {

	fmt.Println("[Producer] Start producing")

	fmt.Println("[Producer] Getting page")
	outString := <-p.OutChan

	doneChan := make(chan bool)

	go func() {
		defer close(doneChan)
		fmt.Println("[Producer] Getting events")
		for e := range p.Events() {
			fmt.Println("[Producer] Fetching events")
			switch ev := e.(type) {
			case *kafka.Message:
				m := ev
				if m.TopicPartition.Error != nil {
					fmt.Errorf("[Producer] Delivery failed: %v\n", m.TopicPartition.Error)
				} else {
					fmt.Printf("[Producer] Delivered message to topic %s [%d] at offset %v\n",
						*m.TopicPartition.Topic, m.TopicPartition.Partition, m.TopicPartition.Offset)
				}
				return

			default:
				fmt.Printf("[Producer] Ignored event: %s\n", ev)
			}
		}
	}()

	fmt.Println("[Producer] Messaging")
	p.KafkaProducer.ProduceChannel() <- &kafka.Message{TopicPartition: kafka.TopicPartition{Topic: &p.Topic, Partition: kafka.PartitionAny}, Value: []byte(outString), Key: []byte(p.MessageKey)}

	fmt.Println("[Producer] Waiting")
	// wait for delivery report goroutine to finish
	_ = <-doneChan
	fmt.Println("[Producer] Closing")

	//p.Close()
	//fmt.Println("Closed")
}
