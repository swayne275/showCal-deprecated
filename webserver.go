// Webserver backend for showCal
package main

// Package externed methods have capital first letter
import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Person struct {
	ID           string `json:"id,omitempty"`
	Firstname    string `json:"firstname,omitempty"`
	Lastname     string `json:"lastname,omitempty"`
	Favoriteshow string `json:"favoriteshow,omitempty"`
}

var persons []Person

func GetAuthors(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(persons)
}

func GetAuthor(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for _, item := range persons {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
}

func CreateAuthor(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var person Person
	_ = json.NewDecoder(r.Body).Decode(&person)
	person.ID = params["id"]
	persons = append(persons, person)
	json.NewEncoder(w).Encode(persons)
}

func DeleteAuthor(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for index, item := range persons {
		if item.ID == params["id"] {
			persons = append(persons[:index], persons[index+1:]...)
			break
		}
		json.NewEncoder(w).Encode(persons)
	}
}

func main() {
	persons = append(persons, Person{ID: "1", Firstname: "Stephen", Lastname: "Wayne", Favoriteshow: "Always sunny"})
	persons = append(persons, Person{ID: "2", Firstname: "Elon", Lastname: "Musk", Favoriteshow: "Bill Cosby's Jailtime Show"})
	router := mux.NewRouter()
	router.HandleFunc("/authors", GetAuthors).Methods("GET")
	router.HandleFunc("/author/{id}", GetAuthor).Methods("GET")
	router.HandleFunc("/author/{id}", CreateAuthor).Methods("POST")
	router.HandleFunc("/author/{id}", DeleteAuthor).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8000", router))
}
