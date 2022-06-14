// main은 프로젝트에서 단 하나여야하지 않나.. 하는 생각..
// test를 위해 이것만 돌릴 때.. run으로 걍 하는 수밖에 없는 것인가.
// 그렇다면(=run만 필요하다면) main pkg + main func 조합이 여러 개여도 상관 없다는건가?

// 소장님 코드에서 보이는 mq.go를 publisher.go로서 그대로 가져왔을 때
// subscriber.go가 정상적으로 받아오면 되는거겠지..?
package main

import (
	"github.com/streadway/amqp"
	"log"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func main() {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"hello",
		true,
		false,
		false,
		false,
		nil,
		)
	failOnError(err, "Failed to declare a queue")

	body := "Hello World!"
	err = ch.Publish(
		"",
		q.Name,
		false,
		false,
		amqp.Publishing {
			ContentType: "text/plain", // 여기에 response.json을 보낼건데.. how?
			Body: []byte(body),
		})
	failOnError(err, "Failed to publish a message")
}

