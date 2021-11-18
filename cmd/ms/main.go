package main

import (
	"context"

	"github.com/matherique/microservice-go/internal/kafka/consumer"
	"github.com/matherique/microservice-go/internal/kafka/producer"
)

func main() {
	ctx := context.Background()

	go producer.MessageProduce(ctx)
	consumer.MessageConsumer(ctx)
}
