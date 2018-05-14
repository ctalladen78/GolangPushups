
package main

import(
	"encoding/json"
	"log"
	"net/http"
	"github.com/gorilla/mux
)

type Person struct {
	ID		string 		`json:"id, omitempty"`
	Firstname	string		`json:"firstname, omitempty"`
	Lastname	string		`json:"lastname, omitempty"`
	Address		*Address	`json:"address, omitempty"`
}

type Address struct {
	City	string	`json:"city, omitempty"`
	State	string	`json:"state, omitempty"`
}

var people []Person

func GetPeople(w http.ResponseWriter, r *http.Request){
	params := mux.Vars(r)

	json.NewEncoder(w).Encode(people)

}
func GetPerson(w http.ResponseWriter, r *http.Request){
	params := mux.Vars(r)
	for _, item := range people {
		if item.ID == params["id"]{
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	// return empty person object
	json.NewEncoder(w).Encode(&Person{})

}
func CreatePerson(w http.ResponseWriter, r *http.Request){
	params := mux.Vars(r)
	var person Person
	_ = json.NewDecoder(r.Body).Decode(&person)
	person.ID = params["id"]
	people = append(people, person)
	json.NewEncoder(w).Encode(people)

}
func UpdatePerson(w http.ResponseWriter, r *http.Request){
	params := mux.Vars(r)
	// TODO log.Println(r.Body)

	json.NewEncoder(w).Encode(people)

}
func DeletePerson(w http.ResponseWriter, r *http.Request){
	params := mux.Vars(r)
	for i, item := range people {
		if item.ID == params["id"] {
			people = append(people[:i], people[i+1:]...)
			break
		}
		json.NewEncoder(w).Encode(people)
	}
}

func main(){

	router := mux.Newrouter()
	people = append(people, Person{ID: "1", Firstname: "John", Lastname: "Doe", Address: &Address{City: "Austin", State: "Texas"}})
	people = append(people, Person{ID: "2", Firstname: "Alice", Lastname: "Doe", Address: &Address{City: "Dallas", State: "Texas"}})
	router.HandleFunc("/api", GetPeople).Methods("GET")
	router.HandleFunc("/api/{id}", GetPerson).Methods("GET")
	router.HandleFunc("/api/{id}", CreatePerson).Methods("POST")
	router.HandleFunc("/api/{id}", DeletePerson).Methods("DELETE")
	router.HandleFunc("/api/{id}", UpdatePerson).Methods("POST")
	log.Fatal(http.ListenAndServe(":8000, router))
	

	

}
