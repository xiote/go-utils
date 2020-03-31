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

	pageChan := make(chan string)

	for _, c := range cases {

		go func() {
			pageChan <- c.in
		}()

		p := Producer{pageChan, kafkaProducerMock, "testTopic", "MessageKey"}
		p.Produce()

		kafkaProducerMock.AssertExpectations(t)
	}
}
