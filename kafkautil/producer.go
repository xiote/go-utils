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
	SendKeyChan   chan string
	SendValueChan chan string
	KafkaProducer
	Topic string
}

func NewProducer(sendChan chan string, kafkaProducer KafkaProducer, topic string) Producer {
	return Producer{nil, sendChan, kafkaProducer, topic}
}

func NewProducer2(sendKeyChan chan string, sendValueChan chan string, kafkaProducer KafkaProducer, topic string) Producer {
	return Producer{sendKeyChan, sendValueChan, kafkaProducer, topic}
}

func (p *Producer) Produce() {

	fmt.Println("[Producer] Start producing")

	fmt.Println("[Producer] Getting data to send")
	var outKey string
	if p.SendKeyChan != nil {
		outKey = <-p.SendKeyChan
	}
	outValue := <-p.SendValueChan

	doneChan := make(chan bool)

	go func() {
		defer close(doneChan)
		fmt.Printf("[Producer] [%s] Getting events\n", p.Topic)
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
	if outKey == "" {
		p.KafkaProducer.ProduceChannel() <- &kafka.Message{TopicPartition: kafka.TopicPartition{Topic: &p.Topic, Partition: kafka.PartitionAny}, Value: []byte(outValue)}
	} else {
		p.KafkaProducer.ProduceChannel() <- &kafka.Message{TopicPartition: kafka.TopicPartition{Topic: &p.Topic, Partition: kafka.PartitionAny}, Key: []byte(outKey), Value: []byte(outValue)}
	}

	fmt.Println("[Producer] Waiting")
	// wait for delivery report goroutine to finish
	_ = <-doneChan
	fmt.Println("[Producer] Closing")

	//p.Close()
	//fmt.Println("Closed")
}
