# 📣 Go Notification Service (Mock-First Microservice)

A modular notification service written in Go, designed for teaching and real-world microservice patterns. It can send email, SMS, and WhatsApp-style messages using a simple pluggable adapter pattern.

✅ Currently runs **without Redis** (mock mode).  
🛠️ Redis/Docker support is **under implementation**.


---

## ▶️ How to Run

### 💡 Run in mock mode (no Redis needed)

On Windows:

    set MOCK_REDIS=true
    go run cmd\main.go

On Unix/macOS:

    MOCK_REDIS=true go run cmd/main.go

This will simulate 2 messages and then allow you to enter live JSON input:

    {"to":"alice","message":"Hello","channel":"email","priority":1}

---

## 💬 Features

- ✅ Modular structure (easy to extend)
- ✅ No Redis/Docker required to test
- ✅ Simulates Email, SMS, and WhatsApp adapters
- 🛠️ Redis pub/sub mode (coming soon)
- 🛠️ Docker support (planned)

---

## 💡 Design Highlights

| Layer         | Purpose                                    |
|---------------|---------------------------------------------|
| adapter.go    | Common interface for all channels           |
| handler.go    | Parses and routes messages                  |
| notifier.go   | Dispatches to correct adapter by `channel`  |
| memory.go     | Simulates messages in mock mode             |

---

## 🌱 Coming Soon

- Dockerfile + docker-compose for Redis
- REST API mode (`POST /notify`)
- Priority queue handling
- Retry and dead-letter queue

---

## 📦 Sample Message JSON

    {
      "to": "bob",
      "message": "Your code is 123456",
      "channel": "sms",
      "priority": 0
    }

---

## 📜 License

MIT — use, remix, teach, or deploy.

---

## 🤝 Contributing

This project is intentionally built for education — PRs welcome once Docker + Redis support is added.
