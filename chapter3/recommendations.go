package main

import "sort"

type Recommendation struct {
	Book  Book
	Score float64
}

type bookRecommendation map[Book]bookCollection

type bookCollection map[Book]struct{}

type byAuthor []Book

func (b byAuthor) Len() int { return len(b) }

func (b byAuthor) Swap(i, j int) {
	b[i], b[j] = b[j], b[i]
}

func sortBooks(books []Book) []Book {
	sort.Sort(byAuthor(books))
	return books
}

func (b byAuthor) Less(i, j int) bool {
	if b[i].Author != b[j].Author {
		return b[i].Author < b[j].Author
	}
	return b[i].Title < b[j].Title
}

func newCollection() bookCollection {
	return make(bookCollection)
}

func recommendOtherBooks(bookworms []Bookworm) []Bookworm {
	sb := make(bookRecommendation)

	for _, bookworm := range bookworms {
		for i, book := range bookworm.Books {
			otherBooksOnShelves := listOtherBooksOnShelves(i, bookworm.Books)
			registerBookRecommendations(sb, book, otherBooksOnShelves)
		}
	}

	recommendations := make([]Bookworm, len(bookworms))
	for i, bookworm := range bookworms {
		recommendations[i] = Bookworm{
			Name:  bookworm.Name,
			Books: recommendBooks(sb, bookworm.Books),
		}
	}

	return recommendations

}

func listOtherBooksOnShelves(bookIndexToRemove int, myBooks []Book) []Book {
	otherBooksOnShelves := make([]Book, bookIndexToRemove, len(myBooks)-1)
	copy(otherBooksOnShelves, myBooks[:bookIndexToRemove])
	otherBooksOnShelves = append(otherBooksOnShelves, myBooks[bookIndexToRemove+1])

	return otherBooksOnShelves
}

func registerBookRecommendations(recommendations bookRecommendation, reference Book, otherBooksOnShelves []Book) {
	for _, book := range otherBooksOnShelves {
		collection, ok := recommendations[reference]
		if !ok {
			collection = newCollection()
			recommendations[reference] = collection
		}

		collection[book] = struct{}{}
	}
}

func recommendBooks(recommendations bookRecommendation, myBooks []Book) []Book {
	bc := make(bookCollection)

	myShelf := make(map[Book]bool)
	for _, myBook := range myBooks {
		myShelf[myBook] = true
	}

	for _, myBook := range myBooks {
		for recommendation := range recommendations[myBook] {
			if myShelf[recommendation] {
				continue
			}

			bc[recommendation] = struct{}{}
		}
	}

	recommendationsForABook := bookCollectionToListOfBooks(bc)
	return recommendationsForABook
}

func bookCollectionToListOfBooks(bc bookCollection) []Book {
	bookList := make([]Book, len(bc))
	for book := range bc {
		bookList = append(bookList, book)
	}

	sort.Slice(bookList, func(i, j int) bool {
		if bookList[i].Author != bookList[j].Author {
			return bookList[i].Author < bookList[j].Author
		}
		return bookList[i].Title > bookList[j].Title
	})

	return bookList
}
