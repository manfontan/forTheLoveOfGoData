package bookstore_test

import (
	"bookstore"
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestBook(t *testing.T) {
	_ = bookstore.Book{
		Title:           "Spark Joy",
		Author:          []string{"Marie Kondō"},
		Copies:          10,
		Description:     "A tiny, cheerful Japanese woman explains tidying.",
		DiscountPercent: 10,
		Edition:         2,
		ID:              "Book1",
		PriceCents:      1199,
		SeriesName:      "Tidying the world",
		SeriesNumber:    1,
		Featured:        true,
	}
}

func TestGetAllBooks(t *testing.T) {

	book1 := bookstore.Book{Title: "Nicholas Chuckleby", Author: []string{"Charles Dickens"}, ID: bookstore.NewID()}
	book2 := bookstore.Book{Title: "Easy Times", Author: []string{"Charles Dickens"}, ID: bookstore.NewID()}
	book3 := bookstore.Book{Title: "Christmas Pudding", Author: []string{"Charles Dickens"}, ID: bookstore.NewID()}

	bookstore.Books = map[string]bookstore.Book{}
	bookstore.Authors = map[string][]string{}
	bookstore.AddBook(book1)
	bookstore.AddBook(book2)
	bookstore.AddBook(book3)

	want := map[string]bookstore.Book{book1.ID: book1, book2.ID: book2, book3.ID: book3}
	got := bookstore.GetAllBooks()

	if !cmp.Equal(got, want) {
		t.Error(cmp.Diff(got, want))
	}
}

func TestGetBookDetails(t *testing.T) {

	book1 := bookstore.Book{Title: "Database Reliability Engineering", ID: bookstore.NewID()}
	book2 := bookstore.Book{Title: "Database Reliability Engineering", ID: bookstore.NewID()}

	bookstore.Books = map[string]bookstore.Book{}
	bookstore.Authors = map[string][]string{}
	bookstore.AddBook(book1)
	bookstore.AddBook(book2)

	want := fmt.Sprintf("Title: %s, ID: %s", book2.Title, book2.ID)
	got, error := bookstore.GetBookDetails(book2.ID)

	if error != nil {
		t.Fatal(error)
	}

	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}

func TestGetAllByAuthor(t *testing.T) {

	book1 := bookstore.Book{Title: "Nicholas Chuckleby", Author: []string{"Charles Dickens"}, ID: bookstore.NewID()}
	book2 := bookstore.Book{Title: "Easy Times", Author: []string{"Charles Dickens"}, ID: bookstore.NewID()}
	book3 := bookstore.Book{Title: "Christmas Pudding", Author: []string{"Charles Dickens", "Manuel Fontan"}, ID: bookstore.NewID()}
	book4 := bookstore.Book{Title: "Metamorphosis", Author: []string{"Kafka"}, ID: bookstore.NewID()}

	bookstore.Books = map[string]bookstore.Book{}
	bookstore.Authors = map[string][]string{}
	bookstore.AddBook(book1)
	bookstore.AddBook(book2)
	bookstore.AddBook(book3)
	bookstore.AddBook(book4)

	want := map[string]bookstore.Book{book1.ID: book1, book2.ID: book2, book3.ID: book3}
	got := bookstore.GetAllByAuthor("Charles Dickens")
	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}

}

func TestNewID(t *testing.T) {

	uniqueIds := map[string]string{}

	for i := 0; i < 10; i++ {
		id := bookstore.NewID()
		_, exist := uniqueIds[id]
		uniqueIds[id] = ""
		if exist {
			t.Errorf("Duplicated ID %s", id)
		}
	}
}

func TestAddBook(t *testing.T) {

	book1 := bookstore.Book{Title: "Nicholas Chuckleby", Author: []string{"Charles Dickens"}, ID: bookstore.NewID()}

	bookstore.Books = map[string]bookstore.Book{}
	bookstore.Authors = map[string][]string{}
	bookstore.AddBook(book1)

	want := fmt.Sprintf("Title: %s, ID: %s", book1.Title, book1.ID)
	got, error := bookstore.GetBookDetails(book1.ID)

	if error != nil {
		t.Fatal(error)
	}

	if want != got {
		t.Error(cmp.Diff(got, want))
	}

	want2 := map[string]bookstore.Book{book1.ID: book1}

	got2 := bookstore.GetAllByAuthor("Charles Dickens")

	if !cmp.Equal(want2, got2) {
		t.Error(cmp.Diff(got2, want2))
	}
}
