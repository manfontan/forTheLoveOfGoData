package bookstore

import (
	"fmt"

	"github.com/google/uuid"
)

type Book struct {
	Title           string
	Author          []string
	Copies          int
	Description     string
	DiscountPercent int
	Edition         int
	ID              string
	PriceCents      int64
	SeriesName      string
	SeriesNumber    int
	Featured        bool
}

var Books = map[string]Book{}
var Authors = map[string][]string{}

//GetAllBooks returns all the books in the bookstore
func GetAllBooks() map[string]Book {
	return Books
}

//GetBookDetails reads the books ID and returns its details in a non elegant
// nor scalable way
func GetBookDetails(id string) (string, error) {
	for _, book := range Books {
		if book.ID == id {
			details := fmt.Sprintf("Title: %s, ID: %s", book.Title, book.ID)
			return details, nil
		}
	}
	return "", fmt.Errorf("there are no books with id: %s", id)
}

//GetAllByAuthor returns all the books from a given Author
func GetAllByAuthor(author string) map[string]Book {
	books := map[string]Book{}
	bookIDs := Authors[author]
	for _, id := range bookIDs {
		books[id] = Books[id]
	}
	return books
}

//NewID returns a unique string ID
func NewID() string {
	return uuid.New().String()
}

//AddBook takes a Book input an appends it to the Books map
func AddBook(b Book) {
	Books[b.ID] = b
	for _, a := range b.Author {
		authorIDs := Authors[a]
		Authors[a] = append(authorIDs, b.ID)
	}
}
