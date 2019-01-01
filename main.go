package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/alexsasharegan/dotenv"
	"github.com/daffaputradamar/go-gorm-restapi/config"
	"github.com/daffaputradamar/go-gorm-restapi/controllers"
	"github.com/gorilla/mux"
)

func handleRequest() {
	r := mux.NewRouter()

	r.HandleFunc("/api/posts", controllers.Index).Methods("GET")
	r.HandleFunc("/api/posts/{id}", controllers.Show).Methods("GET")
	r.HandleFunc("/api/posts", controllers.Store).Methods("POST")
	r.HandleFunc("/api/posts/{id}", controllers.Update).Methods("PUT")
	r.HandleFunc("/api/posts/{id}", controllers.Destroy).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8081", r))
}

func main() {
	err := dotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
	config.ConnectDB()
	fmt.Println("App Listening...")
	handleRequest()
}
