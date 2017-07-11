package main

import (
	"log"
	"net/http"
	"os"

	"github.com/dharnnie/alufunds/fundraiser/handlers"
	"github.com/gorilla/mux"
)

func main() {
	serve()
}

func serve() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "2020"
	}

	http.HandleFunc("/assets/", handlers.ServeResource)

	myMux := mux.NewRouter()
	myMux.HandleFunc("/", handlers.Home)
	myMux.HandleFunc("/drop-a-message", handlers.Message)

	http.Handle("/", myMux)
	err := (http.ListenAndServe(":"+port, nil))
	if err != nil {
		log.Fatal("Server error", err)
	}
}
