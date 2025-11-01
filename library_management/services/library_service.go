package services

import (
	"errors"
	"library_management/models"
)

// Interface
type LibraryManager interface {
	AddBook(book models.Book)
	RemoveBook(bookID int)
	BorrowBook(bookID int, memberID int) error
	ReturnBook(bookID int, memberID int) error
	ListAvailableBooks() []models.Book
	ListBorrowedBooks(memberID int) []models.Book
}

// Implementation
type Library struct {
	Books   map[int]models.Book
	Members map[int]models.Member
}

// Create new library
func NewLibrary() *Library {
	return &Library{
		Books:   make(map[int]models.Book),
		Members: make(map[int]models.Member),
	}
}

// Add a new book
func (l *Library) AddBook(book models.Book) {
	l.Books[book.ID] = book
}

// Remove a book by ID
func (l *Library) RemoveBook(bookID int) {
	delete(l.Books, bookID)
}

// Borrow a book
func (l *Library) BorrowBook(bookID int, memberID int) error {
	book, exists := l.Books[bookID]
	if !exists {
		return errors.New("book not found")
	}
	if book.Status == models.StatusBorrowed {
		return errors.New("book already borrowed")
	}

	member, exists := l.Members[memberID]
	if !exists {
		return errors.New("member not found")
	}

	book.Status = models.StatusBorrowed
	member.BorrowedBooks = append(member.BorrowedBooks, book)
	l.Books[bookID] = book
	l.Members[memberID] = member

	return nil
}

// Return a book
func (l *Library) ReturnBook(bookID int, memberID int) error {
	book, exists := l.Books[bookID]
	if !exists {
		return errors.New("book not found")
	}

	member, exists := l.Members[memberID]
	if !exists {
		return errors.New("member not found")
	}

	found := false
	for i, b := range member.BorrowedBooks {
		if b.ID == bookID {
			member.BorrowedBooks = append(member.BorrowedBooks[:i], member.BorrowedBooks[i+1:]...)
			found = true
			break
		}
	}
	if !found {
		return errors.New("member did not borrow this book")
	}

	book.Status = models.StatusAvailable
	l.Books[bookID] = book
	l.Members[memberID] = member

	return nil
}

// List all available books
func (l *Library) ListAvailableBooks() []models.Book {
	var available []models.Book
	for _, book := range l.Books {
		if book.Status == models.StatusAvailable {
			available = append(available, book)
		}
	}
	return available
}

// List all borrowed books by a member
func (l *Library) ListBorrowedBooks(memberID int) []models.Book {
	member, exists := l.Members[memberID]
	if !exists {
		return []models.Book{}
	}
	return member.BorrowedBooks
}
