package bookdatastore

import (
	"database/sql"
	"fmt"
	"log"

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

var InsertSql = "INSERT INTO %s(%s) VALUES (%s) RETURNING %s"

func DbInsert(table string, columnsInsert string, valuesInsert string, primaryKey string) (int,error) {
	var id int
	db, err := DbConnection()
	if err == nil {
		InsertSql = fmt.Sprintf(InsertSql, table, columnsInsert,valuesInsert, primaryKey)
		log.Println("SQL: ", InsertSql)
		err := db.QueryRow(InsertSql).Scan(&id)
		if err != nil {
			return 0,err
		}
	}
	return id,nil
}
