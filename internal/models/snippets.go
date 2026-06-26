package models

import (
	"database/sql"
	"errors"
	"time"
)


type Snippet struct {
	ID      int
	Title   string
	Content string
	Created time.Time
	Expires time.Time
}

type SnippetModel struct {
	DB *sql.DB
}

func (m *SnippetModel) Insert(title, content string, expires int) (int, error) {
	stmt := `INSERT INTO snippets (title, content, created, expires) VALUES(?, ?, UTC_TIMESTAMP(), DATE_ADD(UTC_TIMESTAMP(), INTERVAL ? DAY))`

	result, err := m.DB.Exec(stmt, title, content, expires)
	if err != nil {
		return 0, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	return int(id), nil
}

func (m *SnippetModel) Get(id int) (*Snippet, error) {
	stmt := `SELECT id, title, content, created,expires FROM snippets WHERE expires > UTC_TIMESTAMP() AND id = ?`

	row := m.DB.QueryRow(stmt, id)

	var s Snippet

	err := row.Scan(&s.ID, &s.Title, &s.Content, &s.Created, &s.Expires)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {

			return &Snippet{}, ErrNoRecord
		} else {
			return &Snippet{}, err
		}
	}

	return &s, nil
}

func (m *SnippetModel) Latest() ([]*Snippet, error) {
	stmt := `SELECT id, title, content, created, expires FROM snippets WHERE expires > UTC_TIMESTAMP() ORDER BY created DESC LIMIT 10`

	// Use the Query() method to execute the SQL statement. This returns a sql.rows resultset containing the result of the query.

	rows, err := m.DB.Query(stmt)
	if err != nil {
		return nil, err
	}

	// Use the defer keyword to schedule a call to the rows.Close() method to ensure that the connection is always returned to the connection pool.
	defer rows.Close()

	// Initialize an empty slice to hold the snippet data.

	var snippets []*Snippet

	// Use rows.Next() to iterate through the rows in the resultset. This prepares the next row for reading with the rows.Scan() method. It returns true if there is another row to be processed, or false if there are no more rows or an error occurs.

	for rows.Next() {
		var s Snippet

		//Use rows.scan() to copy values from each row into a new Snippet struct. The arguments to rows.Scan() should be pointers to the destination values.
		err = rows.Scan(&s.ID, &s.Title, &s.Content, &s.Created, &s.Expires)
		if err != nil {
			return nil, err
		}
		// Append it to the slice of snippe
		snippets = append(snippets, &s)

	}

	// Use rows.Err() to check for any error that may have occurred during the iteration. If an error has occurred, return it from the function.

	if err = rows.Err(); err != nil {
		return nil, err
	}
	
	return snippets, nil
}
