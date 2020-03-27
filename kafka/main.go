package kafka

/**
 * Copyright 2016 Confluent Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

import (
	"fmt"
	"gopkg.in/confluentinc/confluent-kafka-go.v1/kafka"
	"log"
	"os"
	"os/signal"
	"syscall"
)

type KafkaProducer interface {
	Events() chan kafka.Event
	ProduceChannel() chan *kafka.Message
	Close()
}

type Producer struct {
	KafkaProducer
	Topic string
}

func (p *Producer) Produce(message string) {
	//fmt.Println("Start producing")
	doneChan := make(chan bool)

	go func() {
		defer close(doneChan)
		//fmt.Println("Getting events")
		for e := range p.Events() {
			//fmt.Println("Fetching events")
			switch ev := e.(type) {
			case *kafka.Message:
				m := ev
				if m.TopicPartition.Error != nil {
					fmt.Printf("Delivery failed: %v\n", m.TopicPartition.Error)
				} else {
					fmt.Printf("Delivered message to topic %s [%d] at offset %v\n",
						*m.TopicPartition.Topic, m.TopicPartition.Partition, m.TopicPartition.Offset)
				}
				return

			default:
				fmt.Printf("Ignored event: %s\n", ev)
			}
		}
	}()

	//fmt.Println("Messaging")
	p.KafkaProducer.ProduceChannel() <- &kafka.Message{TopicPartition: kafka.TopicPartition{Topic: &p.Topic, Partition: kafka.PartitionAny}, Value: []byte(message)}

	//fmt.Println("Waiting")
	// wait for delivery report goroutine to finish
	_ = <-doneChan
	//fmt.Println("Closing")

	p.Close()
	//fmt.Println("Closed")
}

type KafkaConsumer interface {
	SubscribeTopics(topics []string, rebalanceCb kafka.RebalanceCb) (err error)
	Poll(timeoutMs int) (event kafka.Event)
	Close() (err error)
}
type Consumer struct {
	KafkaConsumer
	MessageHandler func(string) bool
	Topics         []string
}

func (c *Consumer) Consume() {
	topics := c.Topics
	sigchan := make(chan os.Signal, 1)
	signal.Notify(sigchan, syscall.SIGINT, syscall.SIGTERM)

	err := c.KafkaConsumer.SubscribeTopics(topics, nil)
	if err != nil {
		log.Fatal(err)
	}

	run := true

	for run == true {
		select {
		case sig := <-sigchan:
			fmt.Printf("Caught signal %v: terminating\n", sig)
			run = false
		default:
			ev := c.KafkaConsumer.Poll(100)
			if ev == nil {
				continue
			}

			switch e := ev.(type) {
			case *kafka.Message:
				fmt.Printf("%% Message on %s:\n%s\n",
					e.TopicPartition, string(e.Value))
				if e.Headers != nil {
					fmt.Printf("%% Headers: %v\n", e.Headers)
				}
				run = c.MessageHandler(string(e.Value))
			case kafka.Error:
				// Errors should generally be considered
				// informational, the client will try to
				// automatically recover.
				// But in this example we choose to terminate
				// the application if all brokers are down.
				fmt.Fprintf(os.Stderr, "%% Error: %v: %v\n", e.Code(), e)
				if e.Code() == kafka.ErrAllBrokersDown {
					run = false
				}
			default:
				fmt.Printf("Ignored %v\n", e)
			}
		}
	}

	fmt.Printf("Closing consumer\n")
	c.KafkaConsumer.Close()
}
