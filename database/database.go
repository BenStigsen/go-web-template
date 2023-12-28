package database

import (
	"database/sql"
	"log"
	"os"
	"path/filepath"

	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

func Init() {
	// create data folder if it doesn't exist
	os.MkdirAll("data", os.ModePerm)

	// open database
	var err error
	db, err = sql.Open("sqlite3", filepath.Join("data", "database.sqlite")+"?_journal_mode=WAL")
	if err != nil {
		log.Fatal(err)
	}

	// create tables
	_, err = db.Exec(`
    CREATE TABLE IF NOT EXISTS SomeTable (
		id INTEGER PRIMARY KEY
	);`)
	if err != nil {
		log.Panic(err)
	}
}

func Close() {
	db.Close()
}
