package producer

import (
	"fmt"
	"github.com/streadway/amqp"
)

type IUserProducer interface {
	ProducerTransfer() error
}

type UseProducer struct {
}

func (u *UseProducer) ProducerTransfer(tokenString string) error {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	ch, err := conn.Channel()
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"hello", // name
		false,   // durable
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	fmt.Println(q)
	body := tokenString
	err = ch.Publish(
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		})
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	fmt.Println("Successfully Published Message to Queue")
	return nil
}

func NewUseProducer() *UseProducer {
	return &UseProducer{}
}
