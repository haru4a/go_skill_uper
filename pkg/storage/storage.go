package storage

import (
	"log"
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)
var DB *sql.DB

//methods for working with the database.

func getData(){
	rows, err := DB.Query("")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
}

func setData(){
	rows, err := DB.Query("")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
}