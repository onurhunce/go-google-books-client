package googleBooksClient

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type GoogleBookItems struct {
	GoogleBookItem struct {
		Title         string   `json:"title"`
		Authors       []string `json:"authors"`
		PublishedDate string   `json:"publishedDate"`
		PageCount     int      `json:"pageCount"`
		Categories    []string `json:"categories"`
		ImageLinks    struct {
			Thumbnail string `json:"thumbnail"`
		}
		Language string `json:"language"`
	} `json:"volumeInfo"`
}

type GoogleBooksApiResponse struct {
	GoogleBookItemsList []GoogleBookItems `json:"items"`
}

var myClient = &http.Client{Timeout: 10 * time.Second}

const GoogleBooksApiUrl = "https://www.googleapis.com/books/v1/volumes?q="
const APIFields = "fields=totalItems,items(volumeInfo/title,volumeInfo/authors,volumeInfo/publishedDate,volumeInfo/pageCount,volumeInfo/categories,volumeInfo/language,volumeInfo/imageLinks)"

func FindBook(title string, author string, isbn string) *GoogleBooksApiResponse {
	url := GoogleBooksApiUrl
	if len(title) != 0 {
		url = url + "intitle:" + title
	}
	if len(author) != 0 {
		url = url + "+inauthor:" + author
	}

	if len(isbn) != 0 {
		url = url + "+isbn:" + isbn
	}
	fullUrl := url + "&" + APIFields
	var booksResponse = new(GoogleBooksApiResponse)
	jsonData := getJsonResponse(fullUrl)
	err := json.Unmarshal(jsonData, &booksResponse)
	if err != nil {
		fmt.Println("Error during parsing json data:", err)
	}
	return booksResponse
}

func getJsonResponse(url string) []byte {
	r, err := myClient.Get(url)
	if err != nil {
		fmt.Println("Error occurred during Google Books API call: ", err)
	}
	data, _ := ioutil.ReadAll(r.Body)
	r.Body.Close()
	return data
}
