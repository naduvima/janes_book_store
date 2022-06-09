package bookdatastore

/*
   books_id SERIAL PRIMARY KEY,
   author_id SERIAL REFERENCES authors,
   title character varying(255),
   created_at timestamp without time zone NOT NULL default current_timestamp,
   description character varying(255),
   image_s3_url character varying(255), --should store in s3
   price numeric(15,4)
*/

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
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

func PublishBook(ba BooksWithAuthor) (BooksWithAuthor, error) {
	var valuesInsert = "'%s',%d,'%s','%s',%f"
	//find if there is an author exists
	author, err := FindAuthor(ba.Author)
	if err == nil {
		//proceed to publish it with author id
		ba.Book.AuthorID = author.AuthorID
		valuesInsert = fmt.Sprintf(valuesInsert, ba.Book.Title, ba.Book.AuthorID, ba.Book.Description, ba.Book.ImageS3Url, ba.Book.Price)
		log.Println("ValuesInsert", valuesInsert)
		ba.Book.BookID, err = DbInsert("books", ColumnsInsert, valuesInsert, "books_id")
	}
	return ba, err
}

func UnPublishBook(ba BooksWithAuthor, ctx context.Context) error {
	//find if there is an author exists
	ba, err := FindBook(ba, "BooksWithAuthor")
	if err == nil {
		log.Println("current_user routine ", ba.Author.AuthorName, ctx.Value("current_user"))
		if ba.Author.AuthorName != ctx.Value("current_user") {
			return fmt.Errorf("%s is not the publisher of %s", ctx.Value("current_user"), ba.Book.Title)
		}
		//proceed to un publish it with author id
		db, err := DbConnection()
		if err == nil {
			db.QueryRow("delete from books where title = $1 and author_id IN (select author_id from authors where author_name = $2)", ba.Book.Title, ba.Author.AuthorName)
		}
	}
	log.Println(err)
	return err
}

//wants to move SQL operations to datastore
//keeping it as it is time being
//Refactor required to move this to an interface
//Make it clean by moving the SQL statements to datastore
func FindBook(queryParam BooksWithAuthor, argType string) (BooksWithAuthor, error) {
	var ba BooksWithAuthor
	db, err := DbConnection()
	if err == nil {
		switch argType {
		case "BooksWithAuthor":
			selectBookSQL := "select title,author_name,description,image_s3_url,price from books "
			selectBookSQL += ` join authors on authors.author_id = books.author_id where authors.author_name = $1 and books.title = $2`
			log.Println("BooksWithAuthor:", queryParam.Author.AuthorName, queryParam.Book.Title)
			log.Println("BooksWithAuthor:", selectBookSQL)
			err = db.QueryRow(selectBookSQL, queryParam.Author.AuthorName, queryParam.Book.Title).Scan(&ba.Book.Title, &ba.Author.AuthorName, &ba.Book.Description, &ba.Book.ImageS3Url, &ba.Book.Price)
		case "Book":
			selectBookSQL := "select title,author_id,description,image_s3_url,price from books "
			selectBookSQL += ` where books.title = $1`
			err = db.QueryRow(selectBookSQL, queryParam.Book.Title).Scan(&ba.Book.Title, &ba.Author.AuthorID, &ba.Book.Description, &ba.Book.ImageS3Url, &ba.Book.Price)
		case "Author":
			selectBookSQL := "select title,author_name,description,image_s3_url,price from books "
			selectBookSQL += ` join authors on authors.author_id = books.author_id where authors.author_name = $1 `
			err = db.QueryRow(selectBookSQL, queryParam.Author.AuthorName).Scan(&ba.Book.Title, &ba.Author.AuthorName, &ba.Book.Description, &ba.Book.ImageS3Url, &ba.Book.Price)
		}
		if err == nil {
			return ba, nil
		}
	}
	return BooksWithAuthor{}, err
}

func FindBooks(queryParam BooksWithAuthor) ([]BooksWithAuthor, error) {
	db, err := DbConnection()
	if err == nil {
		selectBookSQL := "select title,author_name,description,image_s3_url,price from books "
		selectBookSQL += " join authors on authors.author_id = books.author_id where  1 = 1"
		if queryParam.Author.AuthorName != "" {
			selectBookSQL += " and authors.author_name =  " + queryParam.Author.AuthorName
		}
		if queryParam.Book.Title != "" {
			selectBookSQL += " and books.ttile = " + queryParam.Book.Title
		}

		log.Println("BooksWithAuthor:", queryParam.Author.AuthorName, queryParam.Book.Title)
		log.Println("BooksWithAuthor:", selectBookSQL)
		rows, err := db.Query(selectBookSQL)
		defer rows.Close()

		if err == nil {
			var books []BooksWithAuthor
			for rows.Next() {
				var bk BooksWithAuthor
				if err := rows.Scan(&bk.Book.Title, &bk.Author.AuthorName, &bk.Book.Description,
					&bk.Book.ImageS3Url, &bk.Book.Price); err != nil {
					return books, err
				}
				books = append(books, bk)
			}
			if err = rows.Err(); err == nil {
				return books, err
			}
		}
	}
	return []BooksWithAuthor{}, err
}

func (ba BooksWithAuthor) FillRequest(r *http.Request) BooksWithAuthor {
	body, err := getRawBody(r)
	log.Println("Debug: ", string(body))
	if err == nil {
		err = json.Unmarshal(body, &ba)
	}
	if err != nil {
		log.Println("Error: ", err.Error())
		return BooksWithAuthor{}
	}
	log.Println("Debug: ", ba)
	return ba
}

func (ba BooksWithAuthor) FillRequestFromUrl(r *http.Request) BooksWithAuthor {
	ba = BooksWithAuthor{}
	ba.Book.Title = r.URL.Query().Get("title")
	ba.Author.AuthorName = r.URL.Query().Get("author")
	return ba
}

func getRawBody(r *http.Request) ([]byte, error) {
	if r.ContentLength == 0 {
		return []byte{}, fmt.Errorf("error cannot be empty")
	}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return []byte{}, fmt.Errorf("cannot read request payload")
	}
	return body, nil
}
