package bookdatastore

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

const (
	//host = "janes_books_db" //replace with local host to test on local
	host = "localhost"
	port = 5432
)

func databaseName() string {
	if os.Getenv("TEST_ENV") == "" {
		return "janes_books_db"
	}
	return "janes_books_test_db"
}

func DbConnection() (*sql.DB, error) {
	psqlconn := fmt.Sprintf("host=%s port=%d  dbname=%s sslmode=disable", host, port, databaseName())
	db, err := sql.Open("postgres", psqlconn)
	if err != nil {
		return nil, err
	}
	return db, nil
}

func DbInsert(table string, columnsInsert string, valuesInsert string, primaryKey string) (int, error) {
	var id int
	var insertSql = "INSERT INTO %s(%s) VALUES (%s) RETURNING %s"
	db, err := DbConnection()
	if err == nil {
		insertSql = fmt.Sprintf(insertSql+" "+primaryKey, table, columnsInsert, valuesInsert, primaryKey)
		log.Println("SQL: ", insertSql)
		err := db.QueryRow(insertSql).Scan(&id)
		if err != nil {
			return 0, err
		}
	}
	return id, nil
}
