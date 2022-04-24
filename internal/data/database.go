package data

import (
	"database/sql"
	"log"
	"sync"
)

var (
	database *Database
	once     sync.Once
)

// Manages the connection to the database
type Database struct {
	DB *sql.DB
}

func (d *Database) getConnection() sql.DB {
	return *d.DB
}

// Close resources used by database
func Close() error {
	if database == nil {
		return nil
	}

	return database.DB.Close()
}

// Returns a new instance of database with the connection ready
func New() *Database {
	once.Do(initDB)
	return database
}

// Initialize the database variable with the connection to the database
func initDB() {
	db, err := getConnection()
	if err != nil {
		log.Panic(err)
	}

	database = &Database{
		DB: db,
	}
}
