package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	var path string
	flag.StringVar(&path, "path", "testdata/bookworms.json", "Please provide file with bookworms")
	flag.Parse()
	bookworms, err := loadBookworms(path)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "failed to load bookworms: %v\n", err)
		os.Exit(1)
	}

	commonBooks := findCommonBooks(bookworms)

	fmt.Println("Here are some common books")
	displayBooks(commonBooks)

	recommendations := recommendOtherBooks(bookworms)
	displayRecommendations(recommendations)
}

func displayRecommendations(recommendations []Bookworm) {
	for _, bookworm := range recommendations {
		fmt.Printf("\nHere are the recommendations for %s:\n", bookworm.Name)
		displayBooks(bookworm.Books)
		fmt.Println()
	}
}
