package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
	"worker/rabbitmq"

	amqp "github.com/rabbitmq/amqp091-go"
)

func fatalError(err error, msg string) {
	if err != nil {
		log.Panicf("Fail on %s with err: %v", msg, err)
	}
}

func main() {

	ctx, stop := signal.NotifyContext(
		context.Background(),
		syscall.SIGINT,
		syscall.SIGTERM,
	)
	defer stop()

	brokerURI := os.Getenv("BROKER_URI")
	queueName := os.Getenv("QUEUE_NAME")

	conn, err := amqp.Dial(brokerURI)
	fatalError(err, "connecting to rabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	fatalError(err, "creating channel")
	defer ch.Close()

	q, err := rabbitmq.CreateQueue(ch, queueName)
	fatalError(err, "creating queue")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	deliveryChan, err := rabbitmq.Consume(ctx, ch, q.Name)
	fatalError(err, "receiving delivery")

	<-ctx.Done()
	log.Println("Worker done")
}
