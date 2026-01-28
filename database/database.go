package database

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

func ConnectDB(connection string) (*sql.DB, error) {
	db, err := sql.Open("postgres", connection)
	if err != nil {
		return nil, err
	}

	//test connection
	if err := db.Ping(); err != nil {
		return nil, err
	}

	//set max open connections
	db.SetMaxOpenConns(25)
	//set max idle connections
	db.SetMaxIdleConns(5)

	log.Println("Database connected successfully")

	return db, nil
}
