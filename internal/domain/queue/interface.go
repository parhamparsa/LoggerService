package queue

import "context"

type MessageHandler interface {
	HandleWithRetry([]byte, int) error
}

//go:generate mockgen -source=interface.go -destination=../../../mocks/queue/queue.go
type Interface interface {
	Produce(context.Context, []byte) error
	Consume(handler MessageHandler) error
	Close() error
}
