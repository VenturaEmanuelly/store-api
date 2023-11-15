package repository

import "database/sql"

type Repository struct {
	db *sql.DB
}

func (r Repository) QueryRow(query string, args []any, dest ...any) error {
	return r.db.QueryRow(query, args...).Scan(dest...)

}


func NewRepository(db *sql.DB) Repository {
	return Repository{db: db}
}

