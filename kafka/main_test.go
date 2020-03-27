package kafka

import (
	"gopkg.in/confluentinc/confluent-kafka-go.v1/kafka"
	//"github.com/stretchr/testify/mock"
	"github.com/xiote/go-utils/kafka/mocks"
	"testing"
)

func TestProduce(t *testing.T) {
	kafkaProducerMock := &mocks.KafkaProducer{}
	messageChan := make(chan *kafka.Message, 1)
	eventChan := make(chan kafka.Event, 1)
	go func() {
		message := <-messageChan
		eventChan <- message
	}()
	kafkaProducerMock.On("Events").Return(eventChan).Once()
	kafkaProducerMock.On("ProduceChannel").Return(messageChan).Once()
	kafkaProducerMock.On("Close").Once()

	cases := []struct {
		in string
	}{
		{"TestMessage"},
	}

	for _, c := range cases {
		p := Producer{kafkaProducerMock, "test"}
		p.Produce(c.in)

		kafkaProducerMock.AssertExpectations(t)
	}
}

//type TopicPartition struct {
//	Topic     *string
//	Partition int32
//	Offset    Offset
//	Metadata  *string
//	Error     error
//}

//func (p *Producer) Produce(message string) {
//
//	doneChan := make(chan bool)
//
//	go func() {
//		defer close(doneChan)
//		for e := range p.Events() {
//			switch ev := e.(type) {
//			case *kafka.Message:
//				m := ev
//				if m.TopicPartition.Error != nil {
//					fmt.Printf("Delivery failed: %v\n", m.TopicPartition.Error)
//				} else {
//					fmt.Printf("Delivered message to topic %s [%d] at offset %v\n",
//						*m.TopicPartition.Topic, m.TopicPartition.Partition, m.TopicPartition.Offset)
//				}
//				return
//
//			default:
//				fmt.Printf("Ignored event: %s\n", ev)
//			}
//		}
//	}()
//
//	p.KafkaProducer.ProduceChannel() <- &kafka.Message{TopicPartition: kafka.TopicPartition{Topic: &p.Topic, Partition: kafka.PartitionAny}, Value: []byte(message)}
//
//	// wait for delivery report goroutine to finish
//	_ = <-doneChan
//
//	p.Close()
//}
