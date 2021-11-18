package consumer

import (
	"context"
	"fmt"

	config "github.com/matherique/microservice-go/internal/kafka"
	kafka "github.com/segmentio/kafka-go"
)

func MessageConsumer(ctx context.Context) {
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{config.BrokerAddr},
		Topic:   config.MessageTopic,
	})
	for {
		// the `ReadMessage` method blocks until we receive the next event
		msg, err := r.ReadMessage(ctx)
		if err != nil {
			fmt.Println("could not read message " + err.Error())
			return
		}
		// after receiving the message, log its value
		fmt.Println("received: ", string(msg.Value))
	}
}
