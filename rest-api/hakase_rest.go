
// github.com/hakaselabs/source-blog/blob/master/rest-api/main.go

package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type Person struct {
	ID		string `json:"id, omitempty"`
	Firstname	string `json:"firstname, omitempty"`
	Lastname	string `json:"lastname, omitempty"`
	Address		*Address `json:"address, omitempty"`
}

type Address struct {
	City		string `json:"city, omitempty"`
	State		stirng `json:"state, omitempty"`
}

var people []Person

func GetPeopleEndpoint(w http.ResponseWriter, r *http.Request){
	json.NewEncoder(w).Encode(people)
}

func GetPersonEndpoint(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for _, p := range people {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(p)
			return
		}
	}
	json.NewEncoder(w).Encode(&Person{})
}

func CreatePersonEndpoint(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var person Person	
	_ = json.NewDecoder(r.Body).Decode(&person)
	person.ID = params["id"]
	people = append(people, person)
	json.NewEncoder(w).Encode(people)
}


func DeletePersonEndpoint(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for index, i := range people {
		if i.ID == params["id"] {
			// delete transaction by doing a slice operation 
			people = append(people[:i], people[i+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(people)
}

func main() {
	router := mux.NewRouter()
	people = append(people, Person{ID: "1", Firstname: "Texas", Lastname: "Doe", Address: &Address{City: "Dallas", State:"Texas"}})
	people = append(people, Person{ID: "1", Firstname: "Texas", Lastname: "Doe", Address: &Address{City: "Dallas", State:"Texas"}})
	router.HandleFunc("/people", GetPeopleEndpoint).Methods("GET")
	router.HandleFunc("/people/{id}", GetPersonEndpoint).Methods("GET")
	router.HandleFunc("/people/{id}", CreatePersonEndpoint).Methods("POST")
	router.HandleFunc("/people/{id}", DeletePersonEndpoint).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8000", router))
}
