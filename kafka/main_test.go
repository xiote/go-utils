package kafka

import (
	"log"
	"testing"

	"gopkg.in/confluentinc/confluent-kafka-go.v1/kafka"
)

func TestProduce(t *testing.T) {
	cases := []struct {
		in1 string
		in2 string
		in3 func() kafka.Producer
	}{
		{"testMessage", "test", NewProducer},
	}

	for _, c := range cases {
		Produce(c.in1, c.in2, c.in3)
	}
}

func NewProducer() kafka.Producer {
	p, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": "localhost:9092"})
	if err != nil {
		panic(err)
	}
	return *p
}

func TestConsume(t *testing.T) {
	cases := []struct {
		in1 func(string)
		in2 string
		in3 func() kafka.Consumer
	}{
		{Process, "test", NewConsumer},
	}

	for _, c := range cases {
		Consume(c.in1, c.in2, c.in3)
	}
}

func NewConsumer() kafka.Consumer {
	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": "localhost:9092",
		"group.id":          "group3",
	})

	if err != nil {
		panic(err)
	}
	return *c
}

func Process(jsonString string) {
	log.Println(jsonString)
}
