package mq

import (
	"log"

	"github.com/nats-io/nats"
)

// Handler Mq Handler
type Handler interface{}

// MQ MQ
type MQ struct {
	conn *nats.EncodedConn
}

// Publish Publish
func (mq *MQ) Publish(key string, v Handler) error {
	return mq.conn.Publish(key, v)
}

// Subscribe Subscribe
func (mq *MQ) Subscribe(key string, v Handler) (*nats.Subscription, error) {
	return mq.conn.Subscribe(key, v)
}

// Unsubscribe Unsubscribe
func (mq *MQ) Unsubscribe(handle interface{}) error {
	c := handle.(*nats.Subscription)
	return c.Unsubscribe()
}

// MQD MQD
var MQD MQ

// NewMQ NewMQ
func NewMQ() *MQ {
	return &MQD
}

func init() {
	nc, _ := nats.Connect(nats.DefaultURL)
	c, err := nats.NewEncodedConn(nc, nats.DEFAULT_ENCODER)
	if err != nil {
		log.Print("mq init error")
	} else {
		MQD.conn = c
	}
}
