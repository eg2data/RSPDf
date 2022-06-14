package mq

import (
	"fmt"
	"github.com/eg2data/RSPDf/internal/pkg/transfer"
	"github.com/streadway/amqp"
	"log"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func Receiver() {
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

	msgs, err := ch.Consume(
		q.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
		)

	// goroutine이 쓰인다..!
	// = 비동기
	// = 그러면 변수에 담을 수 없는 것이 당연하네.
	forever := make(chan bool)

	go func() {
		for d := range msgs {
			data := transfer.Transfer(d.Body)
			//log.Printf("Received a message: %s", d.Body)
			fmt.Println(data)


			// d.Body는 []byte type인데.. 음.. => string으로 감싸면 다시 json으로 돌아온다.

			// d.Body를 변수에 잘 담아서 프로젝트에서 사용하고 싶다는 것. 결국.

			// 받아온 데이터를 다른 .go 파일로 던질 수 있나?
			// 이전에는 이게 안되서, 여기로 다 호출했었다. 여기서 다 처리했었다.
			// 왠~지 amqp pkg에 있을 것도 같은데 음.....
			// 1) queue에서 넘어온 json을 저장한다
			// 2) queue에서 받아온 데이터로 차트를 그린다
			// 3) 넘어온 json과 그려진 차트를 기반으로 pdf를 생성한다

		}
	}()

	log.Printf("[*] Waiting for messages. To exit, press CTRL+C")
	<- forever

}