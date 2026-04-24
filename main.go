package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Creating book struct
type Book struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
	Genre  string `json:"genre"`
	Read   bool   `json:"read"`
}

var booksHolder []Book

func bookHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	switch r.Method {
	// Get operation handler
	case http.MethodGet:
		// Returns all books in booksHolder
		if err := json.NewEncoder(w).Encode(booksHolder); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

	// Post operation handler
	case http.MethodPost:
		var newBook Book

		if err := json.NewDecoder(r.Body).Decode(&newBook); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		// Adds new book from user JSON input into booksHolder
		booksHolder = append(booksHolder, newBook)
		w.WriteHeader(http.StatusCreated)

		if err := json.NewEncoder(w).Encode(booksHolder); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

	// Put operation handler
	case http.MethodPut:
		var updatedBook Book
		if err := json.NewDecoder(r.Body).Decode(&updatedBook); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		// Updates book info
		bookFound := false
		for i := range booksHolder {
			// Checks valid bookID entered before updating
			if booksHolder[i].ID == updatedBook.ID {
				booksHolder[i].Title = updatedBook.Title
				booksHolder[i].Author = updatedBook.Author
				booksHolder[i].Genre = updatedBook.Genre
				booksHolder[i].Read = updatedBook.Read

				bookFound = true
				break
			}
		}
		if !bookFound {
			http.Error(w, "Book not found", http.StatusNotFound)
			return
		}

		_ = json.NewEncoder(w).Encode(booksHolder)
		return

	// Delete operation handler
	case http.MethodDelete:
		var bookToDelete Book

		if err := json.NewDecoder(r.Body).Decode(&bookToDelete); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		for i := range booksHolder {
			// Checks for valid bookID
			if booksHolder[i].ID == bookToDelete.ID {
				// Deletes book from booksHolder
				booksHolder = append(booksHolder[:i], booksHolder[i+1:]...)
				break
			} else {
				http.Error(w, "Book not found", http.StatusNotFound)
			}
		}
	}
}

func main() {
	http.HandleFunc("/books", bookHandler)
	fmt.Println("Listening on port 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}

}
