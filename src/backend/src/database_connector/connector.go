package connector

import (
	"database/sql"
	"fmt"
	"log"
)

// Psqlconn is the connection string for the database
const (
	host     = "localhost"
	port     = 5433
	user     = "postgres"
	password = "root"
	dbname   = "postgres"
)

var psqlconn = fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

// ConnectDB connects to the database
func ConnectDB() (*sql.DB, error) {
	db, err := sql.Open("postgres", psqlconn)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	return db, nil
}

// CloseDB closes the database connection
func CloseDB(db *sql.DB) {
	db.Close()
}

// // Get an element from a table in the database
// func Get(selection string, table string, column string, value string) (*sql.Rows, error) {
// 	db, err := ConnectDB()
// 	if err != nil {
// 		log.Fatal(err)
// 		return nil, err
// 	}
// 	defer CloseDB(db)
// 	statment := fmt.Sprintf("SELECT %s FROM %s WHERE %s = '%s'", selection, table, column, value)
// }
