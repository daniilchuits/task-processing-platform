package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"task-service/database"
	"task-service/internal/handlers"
	"task-service/internal/messages/rabbitmq"
	"task-service/internal/repo"
	"time"

	"github.com/go-chi/chi/v5"
	_ "github.com/lib/pq"
	amqp "github.com/rabbitmq/amqp091-go"
)

func fatalError(err error, msg string) {
	if err != nil {
		log.Panicf("Fail on %s with error: %v", msg, err)
	}
}

func main() {

	user := os.Getenv("USER")
	dbname := os.Getenv("DB")
	password := os.Getenv("PASSWORD")
	brockerURI := os.Getenv("BROKER_URI")
	queue := os.Getenv("QUEUE_NAME")
	conn, err := amqp.Dial(brockerURI) // заново подключить брокер, не забыть hostname: rabbitmq
	fatalError(err, "connecting to rabbitmq")
	defer conn.Close()

	cnnStr := fmt.Sprintf(
		"host=tasks-postgres user=%s password=%s dbname=%s sslmode=disable",
		user, password, dbname,
	)

	db, err := sql.Open("postgres", cnnStr)
	fatalError(err, "openning db")
	defer db.Close()

	err = db.Ping()
	fatalError(err, "pinging db")

	dbManager := database.NewDbManager(db)
	rabbitmqManager := rabbitmq.NewConnManager(conn)

	err = dbManager.CreateTable()
	fatalError(err, "creating table tasks")

	repoManager := repo.NewRepoManager(db)

	post := repoManager
	check := repoManager
	selectTasks := repoManager
	selectTask := repoManager

	ctx := context.Background()

	postTaskHandler := handlers.NewPostHandler(check, post, rabbitmqManager, queue)
	selectTasksHandler := handlers.NewSelectHandler(selectTasks)
	selectTaskHandler := handlers.NewSelectOneTaskHandler(selectTask)

	r := chi.NewMux()

	r.Post("/task", postTaskHandler.PostTask)
	r.Get("/task", selectTasksHandler.SelectAllTasks)
	r.Get("/task/{id}", selectTaskHandler.SelectTaskById)

	srv := &http.Server{
		Addr:    ":8082",
		Handler: r,
	}

	ctx, stop := signal.NotifyContext(
		context.Background(),
		syscall.SIGINT,
		syscall.SIGTERM,
	)
	defer stop()

	go srv.ListenAndServe()
	log.Println("Tasks-service started on:", srv.Addr)

	<-ctx.Done()

	ctx, cancel := context.WithTimeout(
		context.Background(),
		5*time.Second,
	)
	defer cancel()

	srv.Shutdown(ctx)
}
