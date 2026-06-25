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
	"task-service/internal/messaging/rabbitmq"
	"task-service/internal/repo"
	"time"

	"github.com/go-chi/chi/v5"
	_ "github.com/lib/pq"
)

func main() {

	user := os.Getenv("USER")
	dbname := os.Getenv("DB")
	password := os.Getenv("PASSWORD")
	brockerURI := os.Getenv("BROKER_URI")
	log.Println("Broker URI:", brockerURI)
	queue := os.Getenv("QUEUE_NAME")

	cnnStr := fmt.Sprintf(
		"host=tasks-postgres user=%s password=%s dbname=%s sslmode=disable",
		user, password, dbname,
	)

	db, err := sql.Open("postgres", cnnStr)
	if err != nil {
		log.Fatal("Error openning db:", err)
	}
	defer db.Close()

	if err = db.Ping(); err != nil {
		log.Fatal("Error pinging db:", err)
	}

	dbManager := database.NewDbManager(db)
	if err = dbManager.CreateTable(); err != nil {
		log.Fatal("Error creating table tasks:", err)
	}

	repoManager := repo.NewRepoManager(db)

	post := repoManager
	check := repoManager
	selectTasks := repoManager
	selectTask := repoManager

	ctx := context.Background()
	connection, err := rabbitmq.NewConn(brockerURI, ctx)
	if err != nil {
		log.Fatal("Err making connection:", err)
	}
	if err = connection.CreateQueue(queue); err != nil {
		log.Fatal("Creating queue error:", err)
	}
	myPublisher, err := connection.NewPublisher(queue)
	if err != nil {
		log.Fatal("Err creating publisher:", err)
	}

	postTaskHandler := handlers.NewPostHandler(check, post, *myPublisher)
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
