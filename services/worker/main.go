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
	queueName := os.Getenv("QUEUE_NAME")

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

	q, err := rabbitmq.CreateQueue(ch, queueName)
	fatalError(err, "creating queue")

	deliveryChan, err := rabbitmq.Consume(ch, q.Name)
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

	newUltimateStruct := difftypes.NewUltimateStruct(
		txtUpdate,
		switcher,
		jpgUpdate,
		mp3Update,
	)

	go func() {
		for i := 0; i < 2; i++ {
			for msg := range deliveryChan {
				if err = newUltimateStruct.DistributeFiles(msg); err != nil {
					log.Println("Error decoding delivery:", err)
				}
				msg.Ack(false)
			}
		}
	}()

	<-notifyCtx.Done()
	log.Println("Worker done")
}
