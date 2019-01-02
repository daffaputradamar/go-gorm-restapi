package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/daffaputradamar/go-gorm-restapi/config"
	"github.com/daffaputradamar/go-gorm-restapi/models"
	"github.com/gorilla/mux"
)

type PostController struct{}

func (p PostController) Index(w http.ResponseWriter, r *http.Request) {
	db := config.ConnectDB()
	var posts []models.Post
	db.Find(&posts)
	defer db.Close()
	json.NewEncoder(w).Encode(posts)
}

func (p PostController) Show(w http.ResponseWriter, r *http.Request) {
	db := config.ConnectDB()
	params := mux.Vars(r)
	var post models.Post
	db.First(&post, params["id"])
	defer db.Close()
	json.NewEncoder(w).Encode(post)
}

func (p PostController) Store(w http.ResponseWriter, r *http.Request) {
	db := config.ConnectDB()
	var post models.Post
	_ = json.NewDecoder(r.Body).Decode(&post)
	db.Create(&post)
	defer db.Close()
	json.NewEncoder(w).Encode(post)
}

func (p PostController) Update(w http.ResponseWriter, r *http.Request) {
	db := config.ConnectDB()
	params := mux.Vars(r)
	var post models.Post
	var updatedPost models.Post
	_ = json.NewDecoder(r.Body).Decode(&updatedPost)
	db.First(&post, params["id"]).Update(&updatedPost)
	defer db.Close()
	json.NewEncoder(w).Encode(post)
}

func (p PostController) Destroy(w http.ResponseWriter, r *http.Request) {
	db := config.ConnectDB()
	params := mux.Vars(r)
	var post models.Post
	db.First(&post, params["id"]).Delete(&post)
	defer db.Close()
	w.WriteHeader(http.StatusNoContent)
}
