package main

import (
	"github.com/gorilla/mux"
	"github.com/sougat818/golang-bank-rest-api--ubiquitous-system/internal/app"
	"log"
	"net/http"
	"time"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/healthcheck", app.GetHealth)

	srv := &http.Server{
		Handler: router,
		Addr:    "127.0.0.1:8080",

		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())

}
