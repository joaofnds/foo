package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func GetConn(host, port, user, password, database string) (*sql.DB, error) {
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
