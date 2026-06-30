package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	difftypes "worker/diff_types"
	"worker/rabbitmq"
	"worker/repo"

	_ "github.com/lib/pq"
)

func fatalError(err error, msg string) {
	if err != nil {
		log.Panicf("Fail on %s with err: %v", msg, err)
	}
}

func main() {

	brokerURI := os.Getenv("BROKER_URI")
	queueNameConsume := os.Getenv("QUEUE_NAME_CONSUME")
	queueNameProduce := os.Getenv("QUEUE_NAME_PRODUCE")

	user := os.Getenv("USER")
	dbname := os.Getenv("DB_NAME")
	password := os.Getenv("PASSWORD")

	cnnStr := fmt.Sprintf(
		"host=tasks-postgres user=%s password=%s dbname=%s sslmode=disable",
		user, password, dbname,
	)

	db, err := sql.Open("postgres", cnnStr)
	fatalError(err, "openning db")
	defer db.Close()

	err = db.Ping()
	fatalError(err, "pinging db")

	tasksRepo := repo.NewTasksRepo(db)

	conn, err := rabbitmq.Connect(brokerURI)
	fatalError(err, "connecting to rabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	fatalError(err, "creating channel")
	defer ch.Close()

	qCons, err := rabbitmq.CreateQueueCosume(ch, queueNameConsume)
	fatalError(err, "creating queue to consume")

	qProd, err := rabbitmq.CreateQueueProduce(ch, queueNameProduce)
	fatalError(err, "creating queue to produce")

	rabbimqProducer := rabbitmq.NewRabbitMqProducer(ch, qProd.Name)
	log.Println(rabbimqProducer)

	deliveryChan, err := rabbitmq.Consume(ch, qCons.Name)
	fatalError(err, "receiving delivery")

	notifyCtx, stop := signal.NotifyContext(
		context.Background(),
		syscall.SIGINT,
		syscall.SIGTERM,
	)
	defer stop()

	switcher := tasksRepo

	txtUpdate := tasksRepo
	jpgUpdate := tasksRepo
	mp3Update := tasksRepo
	csvUpdate := tasksRepo
	zipUpdate := tasksRepo

	newUltimateStruct := difftypes.NewUltimateStruct(
		txtUpdate,
		switcher,
		jpgUpdate,
		mp3Update,
		csvUpdate,
		zipUpdate,
		rabbimqProducer,
	)

	go func() {
		for msg := range deliveryChan {
			if err := newUltimateStruct.DistributeFiles(msg); err != nil {
				log.Println("Error decoding delivery:", err)
			}
			msg.Ack(false)
		}
	}()

	<-notifyCtx.Done()
	log.Println("Worker done")
}
