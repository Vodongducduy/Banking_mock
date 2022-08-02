package qServer

import (
	"banking/internal/tranfer-service/usecases"
	"banking/packages/customResponse"
	"fmt"
	"github.com/streadway/amqp"
	"log"
)

type ITransferMQ interface {
	Consumer()
}

type TransferMQ struct {
	TranferUsecase usecases.ITranferUsecase
}

func NewTransferMQ(tranferUsecase usecases.ITranferUsecase) *TransferMQ {
	return &TransferMQ{TranferUsecase: tranferUsecase}
}

func failOnErrorReceive(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}

func (t *TransferMQ) Consumer() {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	failOnErrorReceive(err, "Failed to connect to RabbitMQ")

	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		fmt.Println(err)
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
	}

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	if err != nil {
		fmt.Println(err)
	}
	var forever chan struct{}
	go func() {
		for d := range msgs {
			err := t.TranferUsecase.CreateTranfer(string(d.Body))
			customResponse.FailErr("Fail to receive message by queue", err)
			log.Printf("[*] Received a message: %s", d.Body)
		}
	}()
	fmt.Println("[*] Waiting for messages. To exit press CTRL+C")
	<-forever
}
