package controllers

import (
	"fmt"
	"library_management/models"
	"library_management/services"
)

func MyLibrary() {
	library := services.NewLibrary()

	for {
		fmt.Println("\n------- Library Management System ------------")
		fmt.Println("1. Add Book")
		fmt.Println("2. Remove Book")
		fmt.Println("3. Borrow Book")
		fmt.Println("4. Return Book")
		fmt.Println("5. List Available Books")
		fmt.Println("6. List Borrowed Books by Member")
		fmt.Println("7. Exit")

		var choice int
		fmt.Print("Enter your choice: ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			addBook(library)
		case 2:
			removeBook(library)
		case 3:
			borrowBook(library)
		case 4:
			returnBook(library)
		case 5:
			listAvailableBooks(library)
		case 6:
			listBorrowedBooks(library)
		case 7:
			fmt.Println("Exiting...")
			return
		default:
			fmt.Println("Invalid choice! Try again.")
		}
	}
}

func addBook(l *services.Library) {
	var id int
	var title, author string

	fmt.Print("Enter Book ID: ")
	fmt.Scan(&id)
	fmt.Print("Enter Title: ")
	fmt.Scan(&title)
	fmt.Print("Enter Author: ")
	fmt.Scan(&author)

	book := models.Book{
		ID:     id,
		Title:  title,
		Author: author,
		Status: models.StatusAvailable,
	}

	l.AddBook(book)
	fmt.Println("Book added successfully!")
}

func removeBook(l *services.Library) {
	var id int
	fmt.Print("Enter Book ID to remove: ")
	fmt.Scan(&id)

	l.RemoveBook(id)
	fmt.Println("Book removed successfully!")
}

func borrowBook(l *services.Library) {
	var memberID, bookID int
	var name string

	fmt.Print("Enter Member ID: ")
	fmt.Scan(&memberID)
	fmt.Print("Enter Member Name: ")
	fmt.Scan(&name)

	// Add member if not exists
	if _, ok := l.Members[memberID]; !ok {
		l.Members[memberID] = models.Member{ID: memberID, Name: name}
	}

	fmt.Print("Enter Book ID to borrow: ")
	fmt.Scan(&bookID)

	err := l.BorrowBook(bookID, memberID)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Book borrowed successfully!")
	}
}

func returnBook(l *services.Library) {
	var memberID, bookID int

	fmt.Print("Enter Member ID: ")
	fmt.Scan(&memberID)
	fmt.Print("Enter Book ID to return: ")
	fmt.Scan(&bookID)

	err := l.ReturnBook(bookID, memberID)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Book returned successfully!")
	}
}

func listAvailableBooks(l *services.Library) {
	books := l.ListAvailableBooks()
	if len(books) == 0 {
		fmt.Println("No available books.")
		return
	}

	fmt.Println("\nAvailable Books:")
	for _, book := range books {
		fmt.Printf("ID: %d | Title: %s | Author: %s | Status: %s\n",
			book.ID, book.Title, book.Author, book.Status)
	}
}

func listBorrowedBooks(l *services.Library) {
	var memberID int
	fmt.Print("Enter Member ID: ")
	fmt.Scan(&memberID)

	books := l.ListBorrowedBooks(memberID)
	if len(books) == 0 {
		fmt.Println("No borrowed books for this member.")
		return
	}

	fmt.Println("\nBorrowed Books:")
	for _, book := range books {
		fmt.Printf("ID: %d | Title: %s | Author: %s\n", book.ID, book.Title, book.Author)
	}
}
