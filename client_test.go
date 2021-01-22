package googleBooksClient

import (
	"strings"
	"testing"
)

const BookTitle = "Harry Potter and the Chamber of Secrets"

func TestStringIsFormatted(t *testing.T) {
	formattedTitle := formatStringForUrl(BookTitle)
	expected := "Harry%20Potter%20and%20the%20Chamber%20of%20Secrets"
	if strings.Contains(string(formattedTitle), expected) == false {
		t.Errorf("handler returned unexpected body: got %v want %v",
			string(formattedTitle), expected)
	}
}

func TestSUrlWithParameters(t *testing.T) {
	formattedTitle := formatStringForUrl(BookTitle)
	author := "Rowling"
	isbn := ""
	urlWithParameters := getUrlWithParameters(formattedTitle, author, isbn)
	expected := "https://www.googleapis.com/books/v1/volumes?q=intitle:Harry%20Potter%20and%20the%20Chamber%20of%20Secrets+inauthor:Rowling"
	if strings.Contains(string(urlWithParameters), expected) == false {
		t.Errorf("handler returned unexpected body: got %v want %v",
			string(urlWithParameters), expected)
	}
}
