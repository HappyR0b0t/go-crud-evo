package repository

import (
	"database/sql"
)

type NumberRepository interface {
	Save(number int) error
	GetAll() ([]int, error)
}

type PostgresNumberRepository struct {
	db *sql.DB
}

func NewPostgresNumberRepository(db *sql.DB) *PostgresNumberRepository {
	return &PostgresNumberRepository{db: db}
}

func (r *PostgresNumberRepository) Save(number int) error {
	_, err := r.db.Exec("INSERT INTO numbers (value) VALUES ($1)", number)
	return err
}

func (r *PostgresNumberRepository) GetAll() ([]int, error) {
	rows, err := r.db.Query("SELECT value FROM numbers ORDER BY value ASC")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var numbers []int
	for rows.Next() {
		var n int
		if err := rows.Scan(&n); err != nil {
			return nil, err
		}
		numbers = append(numbers, n)
	}
	return numbers, nil
}
