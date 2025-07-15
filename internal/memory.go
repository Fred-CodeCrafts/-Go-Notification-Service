package internal

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func StartInMemory() {
	// Phase 1: Simulated hardcoded messages
	messages := []string{
		`{"to":"alice","message":"Hello!","channel":"email","priority":1}`,
		`{"to":"bob","message":"OTP: 1234","channel":"sms","priority":0}`,
	}

	for _, msg := range messages {
		fmt.Println("🧪 Simulating:", msg)
		HandleMessage(msg)
		time.Sleep(1 * time.Second)
	}

	// Phase 2: Live input loop
	fmt.Println("💬 Waiting for live messages (type `exit` to quit)...")

	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("📨 Enter JSON: ")
		input, _ := reader.ReadString('\n')

		if input == "exit\n" || input == "exit\r\n" { // Windows + Unix
			fmt.Println("👋 Exiting.")
			break
		}

		HandleMessage(input)
	}
}
