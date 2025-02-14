package nats

import (
	"context"
)

func (c *Connection) Produce(ctx context.Context, bytes []byte) error {
	_, err := c.js.Publish(c.subject, bytes)
	return err
}
