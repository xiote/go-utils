package kafkautil

import (
	"github.com/stretchr/testify/mock"
	"github.com/xiote/go-utils/kafkautil/mocks"
	"gopkg.in/confluentinc/confluent-kafka-go.v1/kafka"
	"testing"
)

func TestConsume(t *testing.T) {
	message := `abc`
	kafkaConsumerMock := &mocks.KafkaConsumer{}
	topic := "testTopic"
	kafkaConsumerMock.On("SubscribeTopics", []string{topic}, mock.AnythingOfType("kafka.RebalanceCb")).Return(nil)
	kafkaConsumerMock.On("Poll", 100).Return(&kafka.Message{TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny}, Value: []byte(message)})
	kafkaConsumerMock.On("CommitMessage", mock.Anything).Return(nil, nil)
	//kafkaConsumerMock.On("Close").Return(nil)

	want := "abc"
	TopicsChan := make(chan string)
	DoCommitChan := make(chan bool)

	c := Consumer{kafkaConsumerMock, []string{"testTopic"}, TopicsChan, DoCommitChan}
	go c.Consume()

	got := <-TopicsChan
	DoCommitChan <- true

	if got != want {
		t.Errorf("Consume() == %q, want %q", got, want)
	}
	//kafkaConsumerMock.AssertExpectations(t)
}
