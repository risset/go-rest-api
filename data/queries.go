package data

import (
	"database/sql"
)

// Add a row to the article table
func (store *DataStore) AddArticle(article *Article) error {
	query := "INSERT INTO articles (id, user_id, title, slug) VALUES ($1, $2, $3, $4)"
	_, err := store.db.Exec(query, article.ID, article.UserID, article.Title, article.Slug)
	if err != nil {
		return err
	}

	return nil
}

// Update row in table with new article
func (store *DataStore) UpdateArticle(article *Article, id string) error {
	query := "UPDATE articles SET title = $1 WHERE id = $2"
	_, err := store.db.Exec(query, article.Title, id)
	if err != nil {
		return err
	}

	return nil
}

// Delete article with certain id from database
func (store *DataStore) DeleteArticle(id string) error {
	query := "DELETE FROM articles WHERE id = $1"
	_, err := store.db.Exec(query, id)
	if err != nil {
		return err
	}

	return nil
}

// Select all rows from a table
func (store *DataStore) SelectAll() (*sql.Rows, error) {
	query := "SELECT * FROM articles"
	rows, err := store.db.Query(query)
	if err != nil {
		return nil, err
	}

	return rows, nil
}

// Select single row from a table matching a given key
func (store *DataStore) SelectByID(id string) *sql.Row {
	query := "SELECT * FROM articles WHERE id = $1"
	return store.db.QueryRow(query, id)
}
