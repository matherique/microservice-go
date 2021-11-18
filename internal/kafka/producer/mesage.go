package producer

import (
	"context"
	"fmt"
	"strconv"

	config "github.com/matherique/microservice-go/internal/kafka"
	kafka "github.com/segmentio/kafka-go"
)

func MessageProduce(ctx context.Context) {
	i := 0

	w := kafka.NewWriter(kafka.WriterConfig{
		Brokers: []string{config.BrokerAddr},
		Topic:   config.MessageTopic,
	})
	err := make(chan error)

	go func() {
		for {
			msg := kafka.Message{
				Key:   []byte(strconv.Itoa(i)),
				Value: []byte("ola mundo"),
			}

			e := w.WriteMessages(ctx, msg)

			if err != nil {
				err <- e
			}

			i++
		}
	}()

	select {
	case <-ctx.Done():
		fmt.Println("Canceled")
		close(err)
	case err := <-err:
		fmt.Println(err)
	}
}
