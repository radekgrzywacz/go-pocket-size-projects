package main

import "testing"

var (
	handmaidsTale = Book{Author: "Margaret Atwood", Title: "The Handmaid's Tale"}
	oryxAndCrake  = Book{Author: "Margaret Atwood", Title: "Oryx and Crake"}
	theBellJar    = Book{Author: "Sylvia Plath", Title: "The Bell Jar"}
	janeEyre      = Book{Author: "Charlotte Brontë", Title: "Jane Eyre"}
	hobbit        = Book{Author: "Tolkien", Title: "The hobbit"}
)

func TestLoadBookworms(t *testing.T) {
	type testCase struct {
		bookwormsFile string
		want          []Bookworm
		wantErr       bool
	}
	tests := map[string]testCase{
		"File exists": {
			bookwormsFile: "testdata/bookworms.json",
			want: []Bookworm{
				{Name: "Fadi", Books: []Book{handmaidsTale, theBellJar}},
				{Name: "Peggy", Books: []Book{oryxAndCrake, handmaidsTale, janeEyre}},
			},
			wantErr: false,
		},
		"File doesnt exists": {
			bookwormsFile: "testdata/no_file_here.json",
			want:          nil,
			wantErr:       true,
		},
		"Invalid json file": {
			bookwormsFile: "testdata/invalid.json",
			want:          nil,
			wantErr:       true,
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got, err := loadBookworms(tc.bookwormsFile)
			if tc.wantErr {
				if err == nil {
					t.Fatal("expected err, got nothing")
				}
				return
			}

			if err != nil {
				t.Fatalf("expected no error, got %v", err)
			}
			if !equalBookworms(t, got, tc.want) {
				t.Fatalf("different result: got %v, expected %v", got, tc.want)
			}
		})
	}
}

func TestBookCount(t *testing.T) {
	testCases := map[string]struct {
		input []Bookworm
		want  map[Book]uint
	}{
		"Nominal use case": {
			input: []Bookworm{
				{Name: "Fadi", Books: []Book{handmaidsTale, theBellJar}},
				{Name: "Peggy", Books: []Book{oryxAndCrake, handmaidsTale, janeEyre}},
			},
			want: map[Book]uint{
				handmaidsTale: 2,
				theBellJar:    1,
				oryxAndCrake:  1,
				janeEyre:      1,
			},
		},
		"No bookworms": {
			input: []Bookworm{},
			want:  map[Book]uint{},
		},
		"Bookworm without books": {
			input: []Bookworm{
				{Name: "Fadi"},
				{Name: "Peggy"},
			},
			want: map[Book]uint{},
		},
		"Bookworm with twice the same book": {
			input: []Bookworm{
				{Name: "Fadi", Books: []Book{handmaidsTale, handmaidsTale, theBellJar}},
				{Name: "Peggy", Books: []Book{oryxAndCrake, handmaidsTale, janeEyre}},
			},
			want: map[Book]uint{
				handmaidsTale: 3,
				theBellJar:    1,
				oryxAndCrake:  1,
				janeEyre:      1,
			},
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			got := booksCount(tc.input)
			if len(got) != len(tc.want) {
				t.Fatalf("Wrong books count")
			}

			isEqual := equalBooksCount(t, got, tc.want)
			if !isEqual {
				t.Fatalf("Books are not the same")
			}
		})
	}

}

func TestFindCommonBooks(t *testing.T) {
	testCases := map[string]struct {
		input []Bookworm
		want  []Book
	}{
		"Same book": {
			input: []Bookworm{
				{Name: "Fadi", Books: []Book{handmaidsTale, theBellJar}},
				{Name: "Radi", Books: []Book{handmaidsTale, theBellJar}},
			},
			want: []Book{handmaidsTale, theBellJar},
		},
		"Completely different": {
			input: []Bookworm{
				{Name: "Fadi", Books: []Book{handmaidsTale, theBellJar}},
				{Name: "Peggy", Books: []Book{oryxAndCrake, janeEyre}},
			},
			want: []Book{},
		},
		"More than 2 bookworms": {
			input: []Bookworm{
				{Name: "Fadi", Books: []Book{handmaidsTale, theBellJar}},
				{Name: "Peggy", Books: []Book{oryxAndCrake, janeEyre, handmaidsTale}},
				{Name: "Maggy", Books: []Book{handmaidsTale, hobbit}},
			},
			want: []Book{handmaidsTale},
		},
		"One with no books": {
			input: []Bookworm{
				{Name: "Peggy", Books: []Book{oryxAndCrake, janeEyre, handmaidsTale}},
				{Name: "Maggy", Books: []Book{}},
			},
			want: []Book{},
		},
		"Nobody has any book": {
			input: []Bookworm{
				{Name: "Peggy", Books: []Book{}},
				{Name: "Maggy", Books: []Book{}},
			},
			want: []Book{},
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			commonBooks := findCommonBooks(tc.input)

			if !equalBooks(t, commonBooks, tc.want) {
				t.Errorf("Common books different than expected")
			}

		})
	}
}

func equalBookworms(t *testing.T, bookworms, target []Bookworm) bool {
	t.Helper()
	if len(bookworms) != len(target) {
		return false
	}

	for i := range bookworms {
		if bookworms[i].Name != target[i].Name {
			return false
		}

		if !equalBooks(t, bookworms[i].Books, target[i].Books) {
			return false
		}
	}

	return true
}

func equalBooks(t *testing.T, books, target []Book) bool {
	t.Helper()
	if len(books) != len(target) {
		return false
	}

	for i := range books {
		if books[i] != target[i] {
			return false
		}
	}

	return true
}

func equalBooksCount(t *testing.T, got, want map[Book]uint) bool {
	t.Helper()
	if len(got) != len(want) {
		return false
	}

	for book, targetCount := range want {
		count, ok := got[book]
		if !ok || targetCount != count {
			return false
		}
	}

	return true
}
