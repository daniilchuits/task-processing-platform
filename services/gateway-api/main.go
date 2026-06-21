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

	jwtSecret := os.Getenv("JWT")

	auth := "http://auth-service:8081"

	authURL, err := url.Parse(auth)
	if err != nil {
		log.Println("Err parsing auth url:", err)
		return
	}

	secret := jwtmiddlewear.NewSecret(jwtSecret)

	authProxy := httputil.NewSingleHostReverseProxy(authURL)

	secret.JwtMiddlewear()

	r := chi.NewRouter()

	srv := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}

	r.Handle(
		"/auth/*",
		authProxy,
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

	log.Println("Gateway-api ends")
	ctx, cancel := context.WithTimeout(
		context.Background(),
		5*time.Second,
	)
	defer cancel()

	srv.Shutdown(ctx)
}
