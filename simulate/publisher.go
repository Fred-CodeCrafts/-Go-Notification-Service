package main

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
)

func main() {
	rdb := redis.NewClient(&redis.Options{Addr: "localhost:6379"})

	payload := `{"to":"carol","message":"Sale this weekend!","channel":"whatsapp","priority":2}`
	err := rdb.Publish(context.Background(), "notifications", payload).Err()
	if err != nil {
		panic(err)
	}
	fmt.Println("✅ Message published")
}
