package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/letiercamila/api-go-rest/database"
	"github.com/letiercamila/api-go-rest/models"
)

func Home(wr http.ResponseWriter, r *http.Request) {
	fmt.Fprint(wr, "Home")
}

func GetPersonalities(wr http.ResponseWriter, r *http.Request) {
	var p []models.Personality
	database.DB.Find(&p)
	json.NewEncoder(wr).Encode(p)
}

func GetOnePersonality(wr http.ResponseWriter, r *http.Request) {
	param := mux.Vars(r)
	id := param["id"]
	var p models.Personality
	database.DB.First(&p, id)
	json.NewEncoder(wr).Encode(p)
}

func CreateNewPersonality(wr http.ResponseWriter, r *http.Request) {
	var newp models.Personality
	json.NewDecoder(r.Body).Decode(&newp)
	database.DB.Create(&newp)
	json.NewEncoder(wr).Encode(newp)
}

func DeletePersonality(wr http.ResponseWriter, r *http.Request) {
	param := mux.Vars(r)
	id := param["id"]
	var p models.Personality
	database.DB.Delete(&p, id)
	json.NewEncoder(wr).Encode(p)
}

func UpdatePersonality(wr http.ResponseWriter, r *http.Request) {
	param := mux.Vars(r)
	id := param["id"]
	var p models.Personality
	database.DB.First(&p, id)
	json.NewDecoder(r.Body).Decode(&p)
	database.DB.Save(&p)
	json.NewEncoder(wr).Encode(p)
}
