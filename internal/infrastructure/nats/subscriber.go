package nats

import (
	"github.com/nats-io/nats.go"
	"github.com/talon-one/talon-backend-assingment/internal/domain/queue"
	"time"
)

func (c *Connection) Consume(handler queue.MessageHandler) error {
	sub, err := c.js.PullSubscribe(c.subject, c.consumerName)
	if err != nil {
		return err
	}
	c.sub = sub

	for {
		//just for presentation
		messages, err := sub.Fetch(10, nats.MaxWait(time.Hour))
		if err != nil {
			return err
		}
		for _, message := range messages {
			if err = handler.HandleWithRetry(message.Data, 3); err != nil {
				return err
			}
			if err = message.Ack(); err != nil {
				return err
			}
		}
	}
}
