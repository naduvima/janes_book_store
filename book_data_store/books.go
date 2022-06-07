package bookdatastore

import (
	"fmt"
)

type Book struct {
	Title       string  `json:"title"`
	AuthorID    int     `json:"author_id"`
	BookID      int     `json:"book_id"`
	Description string  `json:"description"`
	ImageS3Url  string  `json:"image_s3_url"`
	Price       float64 `json:"price"`
}

type BooksWithAuthor struct {
	Book   Book   `json:"book"`
	Author Author `json:"author"`
}

var ColumnsInsert = "title,author_id,description,image_s3_url,price"
var ValuesInsert = "%s,%d,%s,%s,%d"

func PublishBook(ba BooksWithAuthor) (string, error) {
	//find if there is an author exists
	author, err := FindAuthor(ba.Author)
	if err == nil {
		//proceed to publish it with author id
		ba.Book.AuthorID = author.AuthorID
		ValuesInsert = fmt.Sprintf(ValuesInsert, ba.Book.Title, ba.Book.AuthorID, ba.Book.Description, ba.Book.ImageS3Url, ba.Book.Price)
		DbInsert("books", ColumnsInsert, ValuesInsert)

	}
	return "", err
}
