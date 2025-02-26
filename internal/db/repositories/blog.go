package repositories

import (
	"database/sql"

	m "github.com/gagoto-dev/aevoleyback/internal/api/models"
)

type BlogEntryRepository struct {
	db *sql.DB
}

func NewBlogRepository(db *sql.DB) *BlogEntryRepository {
	return &BlogEntryRepository{
		db: db,
	}
}

func (repo *BlogEntryRepository) CreateBlogEntry(b *m.BlogEntry) (int64, error) {
	result, err := repo.db.Exec("INSERT INTO blog_entry (author_id, title) VALUES (?, ?);", b.Title, b.AuthorId)

	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (repo *BlogEntryRepository) GetBlogEntries(id int64) ([]m.BlogEntry, error) {
	query := `
        SELECT b.*, a.name as author_name 
        FROM blog_entry b 
        LEFT JOIN users a ON a.author_id = b.author_id `
	var values []any

	if id != 0 {
		query += "WHERE id = ? "
		values = append(values, id)
	}
	rows, err := repo.db.Query(query, values...)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var entries []m.BlogEntry
	for rows.Next() {
		var r m.BlogEntry

		if err := rows.Scan(
			&r.Id,
			&r.Title,
			&r.AuthorId,
			&r.AuthorName,
		); err != nil {
			return nil, err
		}

		entries = append(entries, r)
	}

	return entries, nil
}

func (repo *BlogEntryRepository) DeleteBlogEntry(id int64) error {
	_, err := repo.db.Exec("DELETE FROM blog_entry WHERE id = ?;", id)
	return err
}
