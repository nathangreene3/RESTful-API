package main

// PhonebookStatus indicates error types related to managing a Phonebook
type PhonebookStatus int

const (
	// IDExists indicates a Person.ID already exists
	IDExists PhonebookStatus = iota
	// IDNotFound indicates a Person.ID could not be found
	IDNotFound
	// UnknownError indicates an unknown error occurred related to a Phonebook operation
	UnknownError
)

// Status returns a PhonebookStatus similar to Error
func (pe PhonebookStatus) Status() (status string) {
	return status
}

// Error returns an error message related to managing a Phonebook
func (pe PhonebookStatus) Error() (err string) {
	switch pe {
	case IDExists:
		err = "Person ID already exists"
	case IDNotFound:
		err = "Person ID could not be found"
	case UnknownError:
		fallthrough
	default:
		err = "Unknown phonebook error occured"
	}
	return err
}
