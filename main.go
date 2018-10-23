// main.go
// Nathan Greene
// August 2018
//
// This RESTful API follows code mentor at the following link:
// https://www.codementor.io/codehakase/building-a-restful-api-with-golang-a6yivzqdo
// Several changes were made in experimentation.

package main

import (
	"log"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/people", GetPeople).Methods("GET")
	router.HandleFunc("/people/{id}", GetPerson).Methods("GET")
	router.HandleFunc("/people/{id}", CreatePerson).Methods("POST")
	router.HandleFunc("/people/{id}", DeletePerson).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8000", router))
}

func run() (err error) {
	return err
}

// toCSVString converts a Person to a string with each field delimited with a delimeter del (,;\t).
func (person *Person) toCSVString(del string) (info string) {
	if !strings.ContainsAny(*del, ",;\t") {
		del = ","
	}
	return string(person.ID) + del + person.FirstName + del + person.LastName + del + string(person.PhoneNumber) + del + person.Address.City + del + person.Address.State
}
