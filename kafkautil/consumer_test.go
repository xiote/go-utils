package kafkautil

import (
	"github.com/stretchr/testify/mock"
	"github.com/xiote/go-utils/kafkautil/mocks"
	"gopkg.in/confluentinc/confluent-kafka-go.v1/kafka"
	"testing"
)

func TestConsume(t *testing.T) {
	key := `message`
	value := `abc`
	kafkaConsumerMock := &mocks.KafkaConsumer{}
	topic := "testTopic"
	kafkaConsumerMock.On("SubscribeTopics", []string{topic}, mock.AnythingOfType("kafka.RebalanceCb")).Return(nil)
	kafkaConsumerMock.On("Poll", 100).Return(&kafka.Message{TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny}, Key: []byte(key), Value: []byte(value)})
	kafkaConsumerMock.On("CommitMessage", mock.Anything).Return(nil, nil)
	//kafkaConsumerMock.On("Close").Return(nil)

	wantKey := "message"
	wantValue := "abc"
	ReceiveKeyChan := make(chan string)
	ReceiveValueChan := make(chan string)
	DoCommitChan := make(chan bool)

	c := Consumer{kafkaConsumerMock, []string{"testTopic"}, ReceiveKeyChan, ReceiveValueChan, DoCommitChan}
	go c.Consume()

	gotKey := <-ReceiveKeyChan
	gotValue := <-ReceiveValueChan
	DoCommitChan <- true

	if gotKey != wantKey {
		t.Errorf("ConsumeKey() == %q, want %q", gotKey, wantKey)
	}
	if gotValue != wantValue {
		t.Errorf("ConsumeValue() == %q, want %q", gotValue, wantValue)
	}
	//kafkaConsumerMock.AssertExpectations(t)
}
