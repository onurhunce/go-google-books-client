package googleBooksClient

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
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
	url := getUrlWithParameters(title, author, isbn)
	fullUrl := url + "&" + APIFields
	formattedUrl := formatStringForUrl(fullUrl)

	var booksResponse = new(GoogleBooksApiResponse)
	response := getResponse(formattedUrl)
	jsonData := readJsonResponse(response)
	err := json.Unmarshal(jsonData, &booksResponse)
	if err != nil {
		fmt.Println("Error during parsing json data:", err)
	}
	return booksResponse
}

func getUrlWithParameters(title string, author string, isbn string) string {
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
	return url
}

func formatStringForUrl(stringToFormat string) string {
	return strings.ReplaceAll(stringToFormat, " ", "%20")
}

func getResponse(url string) *http.Response {
	response, err := myClient.Get(url)
	if err != nil {
		fmt.Println("Error occurred during Google Books API call: ", err)
	}
	return response
}

func readJsonResponse(response *http.Response) []byte {
	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Error occurred during reading json data: ", err)
		panic(err)
	}
	response.Body.Close()
	return data
}
