# RabbitMQ Backoff

So as can see if you run this locally the messages are consumed ascendingly (from the smaller delay to the biggest delay). In other words the first message with the biggest delay is not blocking the rest of the messages to be consumed. On the contrary, they are all consumed from the shortest delay to the longest:

```bash
$ docker compose up -d
$ go run main.go
broker: RabbitMQ 4.4.0
queue: exp.1794258410735097988

publishing in this order (id = position in plan):
  published #0  delay  8000ms -> due at t=+ 8000ms
  published #1  delay   200ms -> due at t=+  200ms
  published #2  delay   500ms -> due at t=+  500ms
  published #3  delay  4000ms -> due at t=+ 4000ms
  published #4  delay   100ms -> due at t=+  100ms
published 5 messages in 0 ms

arrivals as they happen (t = ms since first publish):
  t=  104ms  arrived #4  (delay   100ms) lateness    +4ms
  t=  203ms  arrived #1  (delay   200ms) lateness    +4ms
  t=  506ms  arrived #2  (delay   500ms) lateness    +6ms
  t= 4007ms  arrived #3  (delay  4000ms) lateness    +7ms
  t= 8009ms  arrived #0  (delay  8000ms) lateness    +9ms

What to look for:
  - If every message arrives near its due time -> NO head-of-line blocking.
  - If a short-delay message arrives only after an earlier-published
    long-delay message is due -> HEAD-OF-LINE BLOCKING.
```

Do **not** forget to run `go mod init` and `go mod tidy` before try to start the app.

<details markdown="1">
<summary><code>main.go</code></summary>

```go
package main

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

var plan = []struct {
	ID      int
	DelayMs int64
}{
	{ID: 0, DelayMs: 8000}, // published first, long delay  -> long due time
	{ID: 1, DelayMs: 200},  // published after #0, short delay -> due very soon
	{ID: 2, DelayMs: 500},
	{ID: 3, DelayMs: 4000},
	{ID: 4, DelayMs: 100},
}

const amqpURL = "amqp://test:test@localhost:35674/"

type payload struct {
	ID            int   `json:"id"`
	DelayMs       int64 `json:"delayMs"`
	PublishedAtMs int64 `json:"publishedAtMs"`
}

func main() {
	conn, err := amqp.Dial(amqpURL)
	must(err)
	defer conn.Close()

	ch, err := conn.Channel()
	must(err)

	qname := fmt.Sprintf("exp.%d", time.Now().UnixNano())
	_, err = ch.QueueDeclare(qname, true, false, false, false, amqp.Table{"x-queue-type": "quorum"})
	must(err)
	fmt.Printf("broker: RabbitMQ %v\nqueue: %s\n\n", conn.Properties["version"], qname)

	// Start consuming BEFORE publishing so nothing is missed.
	deliveries, err := ch.Consume(qname, "", true, false, false, false, nil)
	must(err)

	start := time.Now()
	fmt.Println("publishing in this order (id = position in plan):")
	for _, p := range plan {
		now := time.Now()
		body, _ := json.Marshal(payload{ID: p.ID, DelayMs: p.DelayMs, PublishedAtMs: now.UnixMilli()})
		err = ch.PublishWithContext(context.Background(), "", qname, false, false, amqp.Publishing{
			Body: body,
			Headers: amqp.Table{
				"x-opt-delivery-time": now.Add(time.Duration(p.DelayMs) * time.Millisecond).UnixMilli(),
			},
		})
		must(err)
		due := now.Add(time.Duration(p.DelayMs) * time.Millisecond)
		fmt.Printf("  published #%-2d delay %5dms -> due at t=+%5dms\n",
			p.ID, p.DelayMs, due.Sub(start).Milliseconds())
	}
	fmt.Printf("published %d messages in %d ms\n\n", len(plan), time.Since(start).Milliseconds())

	fmt.Println("arrivals as they happen (t = ms since first publish):")
	deadline := time.After(30 * time.Second)
	for received := 0; received < len(plan); {
		select {
		case d := <-deliveries:
			var b payload
			must(json.Unmarshal(d.Body, &b))
			delay := time.Duration(b.DelayMs) * time.Millisecond
			due := time.UnixMilli(b.PublishedAtMs).Add(delay)
			arrived := time.Now()
			lateness := arrived.Sub(due).Milliseconds()
			fmt.Printf("  t=%5dms  arrived #%-2d (delay %5dms) lateness %+5dms\n",
				arrived.Sub(start).Milliseconds(), b.ID, b.DelayMs, lateness)
			received++
		case <-deadline:
			fmt.Printf("\nTIMEOUT: only %d of %d messages arrived\n", received, len(plan))
			return
		}
	}

	fmt.Println("\nWhat to look for:")
	fmt.Println("  - If every message arrives near its due time -> NO head-of-line blocking.")
	fmt.Println("  - If a short-delay message arrives only after an earlier-published")
	fmt.Println("    long-delay message is due -> HEAD-OF-LINE BLOCKING.")
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}
```

</details>

<details markdown="1">
<summary><code>compose.yml</code></summary>

```yaml
services:
  rabbitmq:
    image: rabbitmq:4.4-management-alpine
    hostname: rabbitmq
    ports:
      - 35674:5672 # AMQP — matches amqpURL in main.go
      - 35675:15672 # management UI (optional, open in browser)
    environment:
      RABBITMQ_DEFAULT_USER: test
      RABBITMQ_DEFAULT_PASS: test
    healthcheck:
      test: ["CMD", "rabbitmq-diagnostics", "ping"]
      interval: 5s
      timeout: 3s
      retries: 10
```

</details>

## Shared in Social Media

- https://www.instagram.com/p/DfiS4CMt5IK/?dlrf=MzRlODBiNWFlZA==&srtk=MzRlODBiNWFlZA==
