package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func SetupDB() *sql.DB {
	db, err := sql.Open("sqlite3", "./database.db")
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	return db
}

func CheckDB(db *sql.DB) bool {

	// Get the database schema
	schema, err := db.Query("SELECT name FROM sqlite_master WHERE type='table'")
	if err != nil {
		log.Fatalf("Failed to retrieve schema: %v", err)
		return false
	}
	defer schema.Close()

	// Print the table names
	for schema.Next() {
		var tableName string
		err := schema.Scan(&tableName)
		if err != nil {
			log.Fatalf("Failed to scan table name: %v", err)
			return false
		}
		fmt.Println(tableName)
	}

	if err := schema.Err(); err != nil {
		log.Fatalf("Error occurred during schema retrieval: %v", err)
		return false
	}

	return true
}
