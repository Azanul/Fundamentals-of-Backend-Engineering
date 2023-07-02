package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	amqp "github.com/rabbitmq/amqp091-go"
)

func main() {
	godotenv.Load("../.env")
	conn, err := amqp.Dial(fmt.Sprintf("amqps://%s:%s@hummingbird.rmq.cloudamqp.com/ypxavwoj", os.Getenv("QUSER"), os.Getenv("PASSWORD")))
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

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	msg := map[string]string{"number": os.Args[1]}
	body, _ := json.Marshal(msg)
	err = ch.PublishWithContext(ctx,
		"",
		q.Name,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        body,
		})
	if err != nil {
		log.Panic(err)
	}
	log.Printf("Job sent successfully %s\n", msg["number"])
}
