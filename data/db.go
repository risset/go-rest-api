package data

import (
	"database/sql"
	"fmt"
)

// Data source for application
type DataStore struct {
	db *sql.DB
}

// Database access parameters
type DbParam struct {
	User    string
	Pass    string
	Name    string
	Address string
}

// Open a connection to a postgresql database and return the DB object
func NewPostgresDB(params DbParam) (*sql.DB, error) {
	dataSourceName := fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable",
		params.User, params.Pass, params.Address, params.Name)

	db, err := sql.Open("postgres", dataSourceName)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}

// Create a new DataStore object
func NewDataStore(params DbParam) (*DataStore, error) {
	var err error
	store := new(DataStore)

	store.db, err = NewPostgresDB(params)
	if err != nil {
		return nil, err
	}

	return store, nil
}
