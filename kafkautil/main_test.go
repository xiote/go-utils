package kafkautil

import (
	"github.com/stretchr/testify/mock"
	"github.com/xiote/go-utils/kafkautil/mocks"
	"gopkg.in/confluentinc/confluent-kafka-go.v1/kafka"
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
func TestConsumer(t *testing.T) {
	wantMessage := "TestMessage"
	kafkaConsumerMock := &mocks.KafkaConsumer{}
	topic := "test"
	kafkaConsumerMock.On("SubscribeTopics", []string{"test"}, mock.AnythingOfType("kafka.RebalanceCb")).Return(nil).Once()
	kafkaConsumerMock.On("Poll", 100).Return(&kafka.Message{TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny}, Value: []byte(wantMessage)}).Once()
	kafkaConsumerMock.On("Close").Return(nil).Once()

	c := Consumer{
		kafkaConsumerMock, func(message string) bool {
			if message == wantMessage {
				// Test End
				return false
			} else {
				t.Errorf("ConsumedMessage() == %q, want %q", message, wantMessage)
				return true
			}
		}, []string{"test"}}
	c.Consume()

	kafkaConsumerMock.AssertExpectations(t)
}
