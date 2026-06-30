package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"notifycation-service/internal/rabbitmq"
	"notifycation-service/internal/repo"
	"notifycation-service/internal/usecase"
	"os"
	"os/signal"
	"syscall"

	_ "github.com/lib/pq"
)

func fatalError(err error, msg string) {
	if err != nil {
		log.Panicf("Fail to %s returned %v", msg, err)
	}
}

func main() {

	brokerURI := os.Getenv("BROKER_URI")
	queueName := os.Getenv("QUEUE_NAME_PRODUCE")

	user := os.Getenv("USER")
	dbname := os.Getenv("DB_NAME")
	password := os.Getenv("PASSWORD")

	cnnStr := fmt.Sprintf(
		"host=tasks-postgres dbname=%s user=%s password=%s sslmode=disable",
		dbname, user, password,
	)
	db, err := sql.Open("postgres", cnnStr)
	fatalError(err, "open db")
	defer db.Close()

	err = db.Ping()
	fatalError(err, "ping db")

	conn, err := rabbitmq.Conn(brokerURI)
	fatalError(err, "connect to rabbitMQ")
	defer conn.Close()

	connManager := rabbitmq.NewConnManager(conn, queueName)
	dbManager := repo.NewNotifyRepo(db)

	q, err := connManager.CreateQueue()
	fatalError(err, "create queue")
	log.Println("queue for noitfication:", q.Name)

	delivery, err := connManager.Consume()
	fatalError(err, "consume messages")

	ctx, stop := signal.NotifyContext(
		context.Background(),
		syscall.SIGINT,
		syscall.SIGTERM,
	)
	defer stop()

	update := usecase.NewUpdate(dbManager)

	go func() {
		for msg := range delivery {
			update.UpdateWithDelivery(msg)
			msg.Ack(false)
		}
	}()

	<-ctx.Done()

}
