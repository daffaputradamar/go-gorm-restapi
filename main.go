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

var postController controllers.PostController

func handleRequest() {
	r := mux.NewRouter()

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello World")
	})
	r.HandleFunc("/api/posts", postController.Index).Methods("GET")
	r.HandleFunc("/api/posts/{id}", postController.Show).Methods("GET")
	r.HandleFunc("/api/posts", postController.Store).Methods("POST")
	r.HandleFunc("/api/posts/{id}", postController.Update).Methods("PUT")
	r.HandleFunc("/api/posts/{id}", postController.Destroy).Methods("DELETE")

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
