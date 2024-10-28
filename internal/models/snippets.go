package models

import "database/sql"

// Define a Snippet type to hold the data for an individual snippet. Notice how
// the fields of the struct correspond to the fields in our PSQL snippets
// table?
type Snippet struct {
	ID        int
	Title     string
	Content   string
	CreatedAt string
	Expires   string
}

type SnippetModel struct {
	DB *sql.DB
}

func (m *SnippetModel) Get(id int) (*Snippet, error) {
	return nil, nil
}

func (m *SnippetModel) Latest() ([]*Snippet, error) {
	return nil, nil
}

func (m *SnippetModel) Insert(title, content string, expires int) (int, error) {

	stmt := `INSERT INTO snippets (title, content, created_at, expires) 
	VALUES ($1, $2,  NOW(), NOW() + ($3 * INTERVAL '1 day')) RETURNING id`

	latestID := 0
	err := m.DB.QueryRow(stmt, title, content, expires).Scan(&latestID)

	if err != nil {
		return 0, err
	}

	return int(latestID), nil
}
