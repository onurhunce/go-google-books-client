package main

import (
	"fmt"
	googleBooks "github.com/onurhunce/go-google-books-client"
)

func main() {
	// Example usages
	title := "tutunamayanlar"
	author := "atay"
	isbn := "9786059503174"

	// Search only via title
	booksResponse := googleBooks.FindBook(title, "atay", "")
	fmt.Println("the response is: ", booksResponse)
	// Returned response:
	//  &{[{{Tutunamayanlar [Oğuz Atay] 1993 724 [Authors, Turkish]
	// {http://books.google.com/books/content?id=fv4HAQAAMAAJ&printsec=frontcover&img=1&zoom=1&source=gbs_api} tr}}]}

	// Search with multiple parameters
	emptyResponse := googleBooks.FindBook(title, author, isbn)
	fmt.Println("the response is: ", emptyResponse)
	// Returned response : &{[]}
}
