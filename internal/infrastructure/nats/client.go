package nats

import (
	"fmt"
	"github.com/nats-io/nats.go"
	"github.com/talon-one/talon-backend-assingment/internal/domain/queue"
	"go.uber.org/zap"
	"log"
)

type Connection struct {
	subject      string
	consumerName string
	js           nats.JetStreamContext
	conn         *nats.Conn
	sub          *nats.Subscription
	stream       *nats.StreamInfo
}

var _ queue.Interface = &Connection{}

const streamName = "generic-stream"

func New(url, defaultSubject string) (*Connection, error) {
	nc, err := nats.Connect(url)
	if err != nil {
		//to replace with zap
		log.Fatalf("Error connecting to NATS: %v", err)
	}
	js, err := nc.JetStream(nats.PublishAsyncMaxPending(256))
	if err != nil {
		log.Fatalf("Error creating JetStream context: %v", err)
	}
	stream, err := createStreamIfDoesntExist(js, defaultSubject)
	if err != nil {
		return nil, err
	}
	return &Connection{conn: nc, subject: defaultSubject, js: js, stream: stream}, nil
}

func createStreamIfDoesntExist(jetStream nats.JetStreamContext, subject string) (*nats.StreamInfo, error) {
	stream, err := jetStream.StreamInfo(streamName)

	// stream not found, create it
	if stream == nil {
		log.Printf("Creating stream: %s\n", streamName)

		stream, err = jetStream.AddStream(&nats.StreamConfig{
			Name:     streamName,
			Subjects: []string{subject},
		})
		if err != nil {
			return nil, err
		}
	}
	return stream, nil
}

func NewConsumer(url, subject string, consumerName string) (*Connection, error) {
	zap.L().Info("Connecting to NATS", zap.String("url", url), zap.String("subject", subject))
	conn, err := New(url, subject)
	if err != nil {
		return nil, err
	}
	conn.consumerName = consumerName
	return conn, err
}
func (c *Connection) Close() error {
	fmt.Println("Closing nats connection")
	if c.sub != nil {
		if subErr := c.sub.Unsubscribe(); subErr != nil {
			return subErr
		}
	}
	c.conn.Close()
	return nil
}
