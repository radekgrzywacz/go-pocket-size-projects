package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Bookworm struct {
	Name  string `json:name`
	Books []Book `json:books`
}

type Book struct {
	Author string `json:author`
	Title  string `json:title`
}

func loadBookworms(filePath string) ([]Bookworm, error) {
	f, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	var bookworms []Bookworm
	err = json.NewDecoder(f).Decode(&bookworms)
	if err != nil {
		return nil, err
	}
	return bookworms, nil
}

func findCommonBooks(bookworms []Bookworm) []Book {
	booksOnTheShelves := booksCount(bookworms)

	var commonBooks []Book
	for book, count := range booksOnTheShelves {
		if count > 1 {
			commonBooks = append(commonBooks, book)
		}
	}

	return sortBooks(commonBooks)
}

func booksCount(bookworms []Bookworm) map[Book]uint {
	count := make(map[Book]uint)

	for _, bookworm := range bookworms {
		for _, book := range bookworm.Books {
			count[book]++
		}
	}

	return count
}

func displayBooks(books []Book) {
	for _, book := range books {
		fmt.Println("-", book.Title, "by", book.Author)
	}
}
