package repository

import "database/sql"

type IRepository interface {
	
}

type Repository struct {
	db	*sql.DB
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{db: db}
}