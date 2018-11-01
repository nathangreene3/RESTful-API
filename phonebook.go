package main

import (
	"encoding/csv"
	"io"
	"sort"
	"strconv"
	"strings"
)

// PhoneBook is a slice of people stored in a csv file.
type PhoneBook struct {
	// FileName references the storage location of the phonebook as a csv file
	FileName string
	// People are listed with their phonenumbers and addresses
	People []Person
}

// inserts a person into the phonebook, then sorts the phonebook. Duplicate insertions are possible.
func (phonebook *PhoneBook) insert(person *Person) {
	if phonebook.indexOf(person.ID) == len(phonebook.People) {
		phonebook.People = append(phonebook.People, *person)
		sort.Sort(phonebook)
	}
}

// removes a person from the phonebook.
func (phonebook *PhoneBook) remove(id int) {
	index := phonebook.indexOf(id)
	if index < len(phonebook.People) {
		phonebook.People = append(phonebook.People[:index], phonebook.People[index+1:]...)
	}
}

// importPhonebook returns a phonebook read from a csv file.
func importPhonebook(filepath string) (*PhoneBook, error) {
	phonebook := &PhoneBook{FileName: filepath}
	reader := csv.NewReader(strings.NewReader(filepath))
	for i := 0; ; i++ {
		data, err := reader.Read()
		if err != nil {
			if err == io.EOF {
				break
			}
			return nil, err
		}
		id, err := strconv.Atoi(data[0])
		if err != nil {
			return nil, err
		}
		phoneNumber, err := strconv.Atoi(data[3])
		if err != nil {
			return nil, err
		}
		phonebook.insert(
			&Person{
				ID:          id,
				FirstName:   data[1],
				LastName:    data[2],
				PhoneNumber: phoneNumber,
				Address: &Address{
					City:  data[4],
					State: data[5],
				},
			},
		)
	}
	return phonebook, nil
}

// updatePhonebook TODO
func updatePhonebook(filepath *string, phonebook *PhoneBook) error {
	// file,err:=os.OpenFile(filepath,os.O_WRONLY,os.ModePerm)
	// if err!=nil{
	// 	return err
	// }
	// defer file.Close()

	// writer:=csv.NewWriter()
	return nil
}

// personExists determines if a person is in a phonebook.
func (phonebook *PhoneBook) personExists(id int) bool {
	if phonebook.indexOf(id) < len(phonebook.People) {
		return true
	}
	return false
}

// indexOf returns the index of a person in a phonebook. If the person is not found by ID, then n is returned.
func (phonebook *PhoneBook) indexOf(id int) int {
	return sort.Search(id, func(i int) bool { return phonebook.People[i].ID == id })
}

// Len returns the number of people in the phonebook. It is used to sort the phonebook.
func (phonebook *PhoneBook) Len() int {
	return len(phonebook.People)
}

// Less returns the less-than comparison of two people in the phonebook. It is used to sort the phonebook.
func (phonebook *PhoneBook) Less(i, j int) bool {
	return phonebook.People[i].ID < phonebook.People[j].ID
}

// Swap swaps two people in the phonebook. It is used to sort the phonebook.
func (phonebook *PhoneBook) Swap(i, j int) {
	phonebook.People[i], phonebook.People[j] = phonebook.People[j], phonebook.People[i]
}
