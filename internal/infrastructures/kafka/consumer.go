package kafka

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/segmentio/kafka-go"

	"5g-v2x-data-management-service/internal/config"
	"5g-v2x-data-management-service/internal/services"
)

// Consumer ...
type Consumer struct {
	config           *config.Config
	AccidentServices *services.AccidentService
}

// NewConsumer ...
func NewConsumer(config *config.Config, as *services.AccidentService) *Consumer {
	c := &Consumer{
		config:           config,
		AccidentServices: as,
	}
	return c
}

func (c *Consumer) consume(ctx context.Context) {
	// initialize a new reader with the brokers and topic
	// the groupID identifies the consumer and prevents
	// it from receiving duplicate messages
	l := log.New(os.Stdout, "kafka reader: ", 0)
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{c.config.Broker1Address, c.config.Broker2Address, c.config.Broker3Address},
		Topic:   c.config.Topic,
		Logger:  l,
	})
	for {
		// the `ReadMessage` method blocks until we receive the next event
		msg, err := r.ReadMessage(ctx)
		if err != nil {
			panic("could not read message " + err.Error())
		}
		// after receiving the message, log its value
		fmt.Println("received: ", string(msg.Value))
		username := fmt.Sprintf("%s", string(msg.Value))
		go c.AccidentServices.StoreData(username)
	}
}

// Start ...
func (c *Consumer) Start() error {
	// create a new context
	ctx := context.Background()
	c.consume(ctx)
	return nil
}
