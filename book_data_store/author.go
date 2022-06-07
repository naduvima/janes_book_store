package bookdatastore

import "database/sql"

type Author struct {
	AuthorID   int    `json:"author_id"`
	Password   string `json:"password"`
	AuthorName string `json:"author_name"`
}

func FindAuthor(author Author) (Author, error) {

	selectUser := `select author_id,password,author_name from authors where author_name = $1`
	db, err := DbConnection()

	if err == nil {
		err := db.QueryRow(selectUser, author.AuthorName).Scan(&author.AuthorID,&author.Password,&author.AuthorName)
		if err != sql.ErrNoRows {
			return author, nil
		}
	}
	return Author{}, err
}
