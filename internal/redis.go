package internal

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
)

func StartRedisSubscriber() {
	rdb := redis.NewClient(&redis.Options{Addr: "localhost:6379"})
	sub := rdb.Subscribe(context.Background(), "notifications")
	ch := sub.Channel()

	for msg := range ch {
		fmt.Println("📩 Received:", msg.Payload)
		HandleMessage(msg.Payload)
	}
}
