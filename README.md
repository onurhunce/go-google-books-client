# go-google-books-client
This library calls Google Books API to search books via specified parameters and it returns JSON response for the desired fields.
Documentation for the API: https://developers.google.com/books/docs/v1/using#PerformingSearch

### Usage

Three parameters can be used to search via Google Books API: `title`, `author`, `isbn`.
Empty string `""` needs to be passed for the not used parameters.

    import googleBooks "github.com/onurhunce/go-google-books-client"
    func main() {
        title = "tutunamayanlar"
        author = "atay"
        isbn = "" 
        booksResponse := googleBooks.FindBook(title, author, isbn)
    }

The returned response will be: 

    &{[{{Tutunamayanlar [Oğuz Atay] 1993 724 [Authors, Turkish] {http://books.google.com/books/content?id=fv4HAQAAMAAJ&printsec=frontcover&img=1&zoom=1&source=gbs_api} tr}}]}
    
    
It is possible to return more fields via extending `APIFields` constant, but you should also
update the `GoogleBookItem` struct via newly added fields as well.

