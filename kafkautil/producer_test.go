package kafkautil

import (
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
	//kafkaProducerMock.On("Close").Once()

	cases := []struct {
		in string
	}{
		{"http://abc"},
	}

	sendChan := make(chan string)

	for _, c := range cases {

		go func() {
			sendChan <- c.in
		}()

		p := NewProducer(sendChan, kafkaProducerMock, "testTopic")
		p.Produce()

		kafkaProducerMock.AssertExpectations(t)
	}
}

func TestProduce2(t *testing.T) {
	kafkaProducerMock := &mocks.KafkaProducer{}
	messageChan := make(chan *kafka.Message, 1)
	eventChan := make(chan kafka.Event, 1)
	go func() {
		message := <-messageChan
		eventChan <- message
	}()
	kafkaProducerMock.On("Events").Return(eventChan).Once()
	kafkaProducerMock.On("ProduceChannel").Return(messageChan).Once()
	//kafkaProducerMock.On("Close").Once()

	cases := []struct {
		in1 string
		in2 string
	}{
		{"url", "http://abc"},
	}

	sendKeyChan := make(chan string)
	sendValueChan := make(chan string)

	for _, c := range cases {

		go func() {
			sendKeyChan <- c.in1
			sendValueChan <- c.in2
		}()

		p := NewProducer2(sendKeyChan, sendValueChan, kafkaProducerMock, "testTopic")
		p.Produce()

		kafkaProducerMock.AssertExpectations(t)
	}
}
