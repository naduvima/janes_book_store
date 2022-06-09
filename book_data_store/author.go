package bookdatastore

import (
	"database/sql"
	"log"
)

type Author struct {
	AuthorID   int    `json:"author_id"`
	Password   string `json:"password"`
	AuthorName string `json:"author_name"`
}

func FindAuthor(author Author) (Author, error) {

	selectUser := `select author_id,password,author_name from authors where author_name = $1`
	db, err := DbConnection()

	if err == nil {
		log.Println(selectUser, "---", author.AuthorName)
		err = db.QueryRow(selectUser, author.AuthorName).Scan(&author.AuthorID, &author.Password, &author.AuthorName)
		log.Println("Found: author", author)
		if sql.ErrNoRows != nil {
			return author, nil
		}
	}
	return Author{}, err
}
