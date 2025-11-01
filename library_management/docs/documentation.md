# Library Management System (Go)

A simple console-based library system in Go to manage books and members.  
**Folders:** main.go, controllers/library_controller.go, models/book.go, models/member.go, services/library_service.go, go.mod  

**Structs:**  
- Book: ID, Title, Author, Status ("Available"/"Borrowed")  
- Member: ID, Name, BorrowedBooks ([]Book)  

**Interface (LibraryManager):** AddBook, RemoveBook, BorrowBook, ReturnBook, ListAvailableBooks, ListBorrowedBooks  

**Menu:** 1. Add Book 2. Remove Book 3. Borrow Book 4. Return Book 5. List Available Books 6. List Borrowed Books by Member 7. Exit  

**Run:** `go run main.go` and follow prompts.  

