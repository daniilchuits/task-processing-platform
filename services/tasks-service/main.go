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
	"task-service/internal/repo"
	"time"

	"github.com/go-chi/chi/v5"
	_ "github.com/lib/pq"
)

func main() {

	user := os.Getenv("USER")
	dbname := os.Getenv("DB")
	password := os.Getenv("PASSWORD")

	cnnStr := fmt.Sprintf(
		"host=tasks-postgres user=%s password=%s dbname=%s sslmode=disable",
		user, password, dbname,
	)
	log.Println(cnnStr)

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
	// доделать insert в insert_task_usecase
	// сделать эндпоинт с POST task

	r := chi.NewMux()

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
