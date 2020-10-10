package data

// Article data model
type Article struct {
	ID     string `json:"id"`
	UserID int64  `json:"user_id"`
	Title  string `json:"title"`
	Slug   string `json:"slug"`
}

// Query database for every row in articles table
func (store *DataStore) GetArticleList() ([]*Article, error) {
	rows, err := store.SelectAll()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	articles := make([]*Article, 0)
	for rows.Next() {
		a := new(Article)
		err := rows.Scan(&a.ID, &a.UserID, &a.Title, &a.Slug)
		if err != nil {
			return nil, err
		}
		articles = append(articles, a)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return articles, nil
}

// Query database for a specific row in articles table
func (store *DataStore) GetArticle(id string) (*Article, error) {
	row := store.SelectByID(id)
	a := new(Article)

	err := row.Scan(&a.ID, &a.UserID, &a.Title, &a.Slug)
	if err != nil {
		return nil, err
	}

	return a, nil
}
