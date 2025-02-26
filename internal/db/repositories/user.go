package repositories

import (
	"database/sql"

	m "github.com/gagoto-dev/aevoleyback/internal/api/models"
	"github.com/google/uuid"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (repo *UserRepository) CreateUserEntry(u *m.User) (int64, error) {
	result, err := repo.db.Exec(
		"INSERT INTO user (id, name, email, password, CreatedAt) VALUES (?, ?, ?, ?, ?);",
		u.Id,
		u.Name,
		u.Email,
		u.Password,
		u.CreatedAt,
	)

	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (repo *UserRepository) GetUsers() ([]m.User, error) {
	query := ` SELECT u.* FROM user u`
	rows, err := repo.db.Query(query)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var entries []m.User
	for rows.Next() {
		var user m.User

		if err := rows.Scan(
			&user.Id,
			&user.Name,
			&user.Password,
			&user.Email,
			&user.CreatedAt,
		); err != nil {
			return nil, err
		}

		entries = append(entries, user)
	}

	return entries, nil
}

func (repo *UserRepository) GetUser(id uuid.UUID) (m.User, error) {
	var user m.User

	query := ` SELECT u.* FROM user u WHERE id = ? `
	rows, err := repo.db.Query(query, id)

	if err != nil {
		return user, err
	}
	defer rows.Close()

	if err := rows.Scan(
		&user.Id,
		&user.Name,
		&user.Password,
		&user.Email,
		&user.CreatedAt,
	); err != nil {
		return user, err
	}

	return user, nil
}

func (repo *UserRepository) DeleteUser(id uuid.UUID) error {
	_, err := repo.db.Exec("DELETE FROM user WHERE id = ?;", id)
	return err
}
