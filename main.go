// main.go
// Nathan Greene
// August 2018
//
// This RESTful API follows code mentor at the following link:
// https://www.codementor.io/codehakase/building-a-restful-api-with-golang-a6yivzqdo
// Several changes were made in experimentation.

package main

import (
	"encoding/csv"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
)

// A Person consists of information regarding a person with a phone.
type Person struct {
	ID          int     `json:"id,omitempty"`
	FirstName   string  `json:"firstname,omitempty"`
	LastName    string  `json:"lastname,omitempty"`
	PhoneNumber int     `json:"phonenumber,omitempty"`
	Address     Address `json:"address,omitempty"`
}

// An Address consists of a simple location.
type Address struct {
	City  string `json:"city,omitempty"`
	State string `json:"state,omitempty"`
}

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

// GetPeople returns all people from the phonebook.
func GetPeople(w http.ResponseWriter, r *http.Request) {
	if people, err := importPhonebook("phonebook.csv"); err != nil {
		w.Write([]byte(err.Error()))
	} else {
		json.NewEncoder(w).Encode(people)
	}
}

// GetPerson retrieves a Person given a valid id.
func GetPerson(w http.ResponseWriter, r *http.Request) {
	people := importPhonebook("phonebook.csv")
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if i, exists := personExists(id, people); err == nil && exists {
		json.NewEncoder(w).Encode(people[i])
	} else {
		w.Write([]byte(err.Error()))
	}
}

func CreatePerson(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	person := newPerson()
	person.ID, _ = strconv.Atoi(params["id"])
	person.FirstName = params["firstname"]
	person.LastName = params["lastname"]
	person.PhoneNumber, _ = strconv.Atoi(params["phonenumber"])
	person.Address.City = params["city"]
	person.Address.State = params["state"]
	insertIntoPhonebook("phonebook.csv", person)
}

func newPerson() (person Person) {
	return Person{
		ID:          0,
		FirstName:   "",
		LastName:    "",
		PhoneNumber: 0,
		Address: Address{
			City:  "",
			State: "",
		},
	}
}

// DeletePerson removes a person from the Phonebook given a valid ID.
func DeletePerson(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	people, err := importPhonebook("phonebook.csv")
	if err != nil {
		log.Fatal(err)
	}
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		log.Fatal(err)
	} else {
		people, err = deleteFromPhonebook(id, people)
		err = updatePhonebook("phonebook.csv", people)
	}
}

func importPhonebook(filepath string) (people []Person, err error) {
	r := csv.NewReader(strings.NewReader(filepath))
	for i := 0; err != io.EOF; i++ {
		if data, err := r.Read(); err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		} else {
			// Consider checking for error on conversions
			id, _ := strconv.Atoi(data[0])
			phoneNumber, _ := strconv.Atoi(data[3])
			people = append(people[:i], Person{
				ID:          id,
				FirstName:   data[1],
				LastName:    data[2],
				PhoneNumber: phoneNumber,
				Address: Address{
					City:  data[4],
					State: data[5],
				},
			})
		}
	}
	return people, err
}

func updatePhonebook(filepath string, people []Person) (err error) {
	return err
}

//
func insertIntoPhonebook(filepath string, person Person) (err error) {
	people := importPhonebook(filepath)
	if personExists(person.ID, people) {
		err = IDExists
	} else {
		people = append(people, person)
	}
	return err
}

// deleteFromPhonebook removes a Person from a []Person.
func deleteFromPhonebook(id int, people []Person) (lesspeople []Person, err error) {
	if i, exists := personExists(id, people); exists {
		lesspeople = people[:i]
		if i+1 < len(people) {
			lesspeople = append(people[i+1:])
		}
	} else {
		lesspeople, err = people, IDNotFound
	}
	return lesspeople, err
}

// toCSVString converts a Person to a string with each field delimited with a delimeter del (,;\t).
func toCSVString(person Person, del string) (info string) {
	if !strings.ContainsAny(del, ",;\t") {
		del = ","
	}
	return string(person.ID) + del + person.FirstName + del + person.LastName + del + string(person.PhoneNumber) + del + person.Address.City + del + person.Address.State
}

// personExists determines if a Person is in a []Person. The index of the Person is returned -1 if not found.
func personExists(id int, people []Person) (index int, exists bool) {
	for index, person := range people {
		if id == person.ID {
			exists = true
			break
		}
	}
	if !exists {
		index = -1
	}
	return index, exists
}
