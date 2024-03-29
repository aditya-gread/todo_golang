package config

import "database/sql"

var (
	db *sql.DB
)

// function that connects to db
func Connect() {
	dsn := "root:12345678@tcp(127.0.0.1:3306)/testing"
	d, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}
	db = d
}

// returns instance of db
func GetDB() *sql.DB {
	return db
}
