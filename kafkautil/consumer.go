package kafkautil

import (
	"fmt"
	"gopkg.in/confluentinc/confluent-kafka-go.v1/kafka"
	"log"
	"os"
	"os/signal"
	"syscall"
)

type KafkaConsumer interface {
	SubscribeTopics(topics []string, rebalanceCb kafka.RebalanceCb) (err error)
	Poll(timeoutMs int) (event kafka.Event)
	Close() (err error)
	CommitMessage(m *kafka.Message) ([]kafka.TopicPartition, error)
}

type Consumer struct {
	KafkaConsumer
	Topics           []string
	ReceiveKeyChan   chan string
	ReceiveValueChan chan string
	DoCommitChan     chan bool
}

func (c *Consumer) Consume() {
	topics := c.Topics
	sigchan := make(chan os.Signal, 1)
	signal.Notify(sigchan, syscall.SIGINT, syscall.SIGTERM)

	err := c.KafkaConsumer.SubscribeTopics(topics, nil)
	fmt.Printf("[Consumer] SubscribeTopics(%v)\n", topics)
	if err != nil {
		log.Fatal(err)
	}

	run := true

	for run == true {
		select {
		case sig := <-sigchan:
			fmt.Printf("[Consumer] Caught signal %v: terminating\n", sig)
			run = false
		default:
			ev := c.KafkaConsumer.Poll(100)
			if ev == nil {
				continue
			}

			switch e := ev.(type) {
			case *kafka.Message:
				fmt.Printf("[Consumer] %% Message on %s:\n",
					e.TopicPartition)
				if e.Headers != nil {
					fmt.Printf("[Consumer] %% Headers: %v\n", e.Headers)
				}
				c.ReceiveKeyChan <- string(e.Key)
				c.ReceiveValueChan <- string(e.Value)
				<-c.DoCommitChan // wait until done
				fmt.Printf("[Consumer] Committing...\n")
				c.CommitMessage(e) // commit

			case kafka.Error:
				// Errors should generally be considered
				// informational, the client will try to
				// automatically recover.
				// But in this example we choose to terminate
				// the application if all brokers are down.
				//fmt.Fprintf(os.Stderr, "[Consumer] %% Error: %v: %v\n", e.Code(), e)
				//if e.Code() == kafka.ErrAllBrokersDown {
				//	run = false
				//}
				fmt.Errorf("[Consumer] %% Error: %v: %v\n", e.Code(), e)
			default:
				fmt.Printf("[Consumer] Ignored %v\n", e)
			}
		}
	}

	fmt.Printf("[Consumer] Closing consumer\n")
	c.KafkaConsumer.Close()
}
