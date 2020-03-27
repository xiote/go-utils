// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import (
	confluent_kafka_go_v1kafka "gopkg.in/confluentinc/confluent-kafka-go.v1/kafka"

	mock "github.com/stretchr/testify/mock"
)

// KafkaProducer is an autogenerated mock type for the KafkaProducer type
type KafkaProducer struct {
	mock.Mock
}

// Close provides a mock function with given fields:
func (_m *KafkaProducer) Close() {
	_m.Called()
}

// Events provides a mock function with given fields:
func (_m *KafkaProducer) Events() chan confluent_kafka_go_v1kafka.Event {
	ret := _m.Called()

	var r0 chan confluent_kafka_go_v1kafka.Event
	if rf, ok := ret.Get(0).(func() chan confluent_kafka_go_v1kafka.Event); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(chan confluent_kafka_go_v1kafka.Event)
		}
	}

	return r0
}

// ProduceChannel provides a mock function with given fields:
func (_m *KafkaProducer) ProduceChannel() chan *confluent_kafka_go_v1kafka.Message {
	ret := _m.Called()

	var r0 chan *confluent_kafka_go_v1kafka.Message
	if rf, ok := ret.Get(0).(func() chan *confluent_kafka_go_v1kafka.Message); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(chan *confluent_kafka_go_v1kafka.Message)
		}
	}

	return r0
}