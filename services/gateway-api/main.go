package main

import (
	"context"
	jwtmiddlewear "gateway/internal/jwt_middlewear"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-chi/chi/v5"
)

func main() {

	// slog.Log() // добавить продвинутый логер в каждом сервисе
	// узнать как вообще сервисы получают свой logger

	jwtSecret := os.Getenv("JWT")
	secret := jwtmiddlewear.NewSecret(jwtSecret)

	auth := "http://auth-service:8081"
	tasks := "http://tasks-service:8082"

	authURL, err := url.Parse(auth)
	if err != nil {
		log.Println("Err parsing auth url:", err)
		return
	}
	tasksURL, err := url.Parse(tasks)
	if err != nil {
		log.Println("Err parsing tasks url:", err)
		return
	}

	authProxy := httputil.NewSingleHostReverseProxy(authURL)
	tasksProxy := httputil.NewSingleHostReverseProxy(tasksURL)

	tasksProtected := secret.JwtMiddlewear(tasksProxy)

	r := chi.NewRouter()

	srv := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}

	r.Handle(
		"/auth/*",
		authProxy,
	)
	r.Handle(
		"/task/*",
		tasksProtected,
	)

	ctx, stop := signal.NotifyContext(
		context.Background(),
		syscall.SIGINT,
		syscall.SIGTERM,
	)
	defer stop()

	go srv.ListenAndServe()
	log.Println("Gateway started at:", srv.Addr)

	<-ctx.Done()

	ctx, cancel := context.WithTimeout(
		context.Background(),
		5*time.Second,
	)
	defer cancel()

	srv.Shutdown(ctx)
}
