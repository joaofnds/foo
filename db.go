package main

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

var (
	host     = os.Getenv("POSTGRESQL_HOST")
	port     = os.Getenv("POSTGRESQL_PORT")
	user     = os.Getenv("POSTGRESQL_USER")
	password = os.Getenv("POSTGRESQL_PASSWORD")
	database = os.Getenv("POSTGRESQL_DATABASE")
)

func GetConn() (*sql.DB, error) {
	conn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, database)
	fmt.Println("conn: " + conn)

	db, err := sql.Open("postgres", conn)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return db, nil
}
