package models

import (
	"database/sql"
	"errors"
)

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
	stmt := `SELECT * FROM snippets WHERE expires > NOW() AND id = $1`

	row := m.DB.QueryRow(stmt, id)

	s := &Snippet{}

	err := row.Scan(&s.ID, &s.Title, &s.Content, &s.CreatedAt, &s.Expires)

	if err != nil {

		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrNoRecord
		} else {
			return nil, err
		}
	}

	return s, nil
}

func (m *SnippetModel) Latest() ([]*Snippet, error) {
	stmt := `SELECT * FROM snippets WHERE expires > NOW() ORDER BY id DESC LIMIT 10`

	rows, err := m.DB.Query(stmt)

	if err != nil {
		return nil, err
	}

	// We defer rows.Close() to ensure the sql.Rows resultset is
	// always properly closed before the Latest() method returns. This defer
	// statement should come *after* you check for an error from the Query()
	// method. Otherwise, if Query() returns an error, you'll get a panic
	// trying to close a nil resultset.
	defer rows.Close()

	snippets := []*Snippet{}

	for rows.Next() {
		s := &Snippet{}

		err = rows.Scan(&s.ID, &s.Title, &s.Content, &s.CreatedAt, &s.Expires)

		if err != nil {
			return nil, err
		}

		snippets = append(snippets, s)

		if err = rows.Err(); err != nil {
			return nil, err
		}
	}

	return snippets, nil
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
