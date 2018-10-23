package main

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// A Person consists of information regarding a person with a phone.
type Person struct {
	ID          int      `json:"id,omitempty"`
	FirstName   string   `json:"firstname,omitempty"`
	LastName    string   `json:"lastname,omitempty"`
	PhoneNumber int      `json:"phonenumber,omitempty"`
	Address     *Address `json:"address,omitempty"`
}

// GetPeople returns all people from the phonebook.
func GetPeople(w http.ResponseWriter, r *http.Request) {
	phonebook, err := importPhonebook("phonebook.csv")
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}
	err = json.NewEncoder(w).Encode(phonebook.People)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}
}

// GetPerson retrieves a Person given a valid id.
func GetPerson(w http.ResponseWriter, r *http.Request) {
	phonebook, err := importPhonebook("phonebook.csv")
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}
	index := phonebook.indexOf(id)
	if index == len(phonebook.People) {
		w.Write([]byte(err.Error()))
		return
	}
	json.NewEncoder(w).Encode(phonebook.People[index])
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}
}

// CreatePerson inserts a new person into the phonebook.
func CreatePerson(w http.ResponseWriter, r *http.Request) {
	// params := mux.Vars(r)
	// person := newPerson()
	// person.ID, _ = strconv.Atoi(params["id"])
	// person.FirstName = params["firstname"]
	// person.LastName = params["lastname"]
	// person.PhoneNumber, _ = strconv.Atoi(params["phonenumber"])
	// person.Address.City = params["city"]
	// person.Address.State = params["state"]
	// insertIntoPhonebook("phonebook.csv", person)
}

func newPerson() (person Person) {
	return Person{
		ID:          0,
		FirstName:   "",
		LastName:    "",
		PhoneNumber: 0,
		Address: &Address{
			City:  "",
			State: "",
		},
	}
}

// DeletePerson removes a person from the Phonebook given a valid ID.
func DeletePerson(w http.ResponseWriter, r *http.Request) {
	// params := mux.Vars(r)
	// people, err := importPhonebook("phonebook.csv")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// id, err := strconv.Atoi(params["id"])
	// if err != nil {
	// 	log.Fatal(err)
	// } else {
	// 	people, err = deleteFromPhonebook(id, people)
	// 	err = updatePhonebook("phonebook.csv", people)
	// }
}
