package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

func handle(w http.ResponseWriter, r *http.Request) {
	log.Println("task")
	idStr := r.Header.Get("user_id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "User_id not found", 404)
		return
	}
	fmt.Fprintln(w, "User_id:", id)
}

func main() {

	r := chi.NewMux()

	srv := &http.Server{
		Addr:    ":8082",
		Handler: r,
	}

	r.Get("/task/", handle) // при GET на /task/ показываает код 200, но не дает ответ #чзхэ

	log.Fatal(srv.ListenAndServe())
}
