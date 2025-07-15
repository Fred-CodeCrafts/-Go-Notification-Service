package main

import (
	"fmt"
	"os"
	"notification-service/internal"
)

func main() {
	if os.Getenv("MOCK_REDIS") == "true" {
		fmt.Println("🧪 Running without Redis")
		internal.StartInMemory()
	} else {
		fmt.Println("🔌 Subscribing to Redis")
		internal.StartRedisSubscriber()
	}
}
