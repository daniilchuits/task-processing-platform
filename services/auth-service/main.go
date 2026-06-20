package main

import (
	"auth/database"
	"auth/internal/repo"
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-chi/chi/v5"
	_ "github.com/lib/pq"
)

func main() {

	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	cnnStr := fmt.Sprintf(
		"host=auth-postgres user=%s password=%s dbname=%s sslmode=disable",
		user, password, dbName,
	)
	log.Println("CnnStr:", cnnStr)

	db, err := sql.Open("postgres", cnnStr)
	if err != nil {
		log.Fatal("Error connecting to db:", err)
	}
	defer db.Close()

	if err = db.Ping(); err != nil {
		log.Fatal("Error pinging db:", err)
	}

	dbManager := database.NewDbManager(db)
	if err = dbManager.Create(); err != nil {
		log.Fatal("Error creating table:", err)
	}

	// repoManager228
	repoManager := repo.NewRepoManager(db)
	// repoManager228

	r := chi.NewRouter()

	srv := &http.Server{
		Addr:    ":8081",
		Handler: r,
	}

	ctx, stop := signal.NotifyContext(
		context.Background(),
		syscall.SIGINT,
		syscall.SIGTERM,
	)
	defer stop()

	go srv.ListenAndServe()
	log.Println("Auth-service started on:", srv.Addr)

	<-ctx.Done()

	ctx, cancel := context.WithTimeout(
		context.Background(),
		5*time.Second,
	)
	defer cancel()

	srv.Shutdown(ctx)
}
