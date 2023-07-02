package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	amqp "github.com/rabbitmq/amqp091-go"
)

func main() {
	godotenv.Load("../.env")
	conn, err := amqp.Dial(fmt.Sprintf("amqps://%s:%s@hummingbird-01.rmq.cloudamqp.com/ypxavwoj", os.Getenv("QUSER"), os.Getenv("PASSWORD")))
	if err != nil {
		log.Panic(err)
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		log.Panic(err)
	}
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"jobs",
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Panic(err)
	}

	msgs, err := ch.Consume(
		q.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Panic(err)
	}

	var forever chan struct{}

	go func() {
		msg := make(map[string]string)
		for d := range msgs {
			json.Unmarshal(d.Body, &msg)
			log.Printf("Received a message: %s", msg["number"])
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}
