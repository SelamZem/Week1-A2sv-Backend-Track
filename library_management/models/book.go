package models

type Book struct {
	ID     int
	Title  string
	Author string
	Status string
}

const (
	StatusAvailable = "Available"
	StatusBorrowed  = "Borrowed"
)
