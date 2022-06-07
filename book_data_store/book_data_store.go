package bookdatastore

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	//host = "janes_books_db" //replace with local host to test on local
	host     = "localhost"
	port     = 5432
	database = "janes_books_db"
)

func DbConnection() (*sql.DB, error) {
	psqlconn := fmt.Sprintf("host=%s port=%d  dbname=%s sslmode=disable", host, port, database)
	db, err := sql.Open("postgres", psqlconn)
	if err != nil {
		return nil, err
	}
	return db, nil
}

var INSERTSQL = "INSERT INTO %s(%s) VALUES ($1)"

func DbInsert(table string, columnsInsert string, valuesInsert string) error {
	db, err := DbConnection()
	if err == nil {
		INSERTSQL = fmt.Sprintf(INSERTSQL, table, columnsInsert)
		_, err := db.Exec(INSERTSQL, valuesInsert)
		if err != nil {
			return err
		}
	}
	return nil
}
