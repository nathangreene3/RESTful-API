package main

// An Address consists of a simple location.
type Address struct {
	City  string `json:"city,omitempty"`
	State string `json:"state,omitempty"`
}