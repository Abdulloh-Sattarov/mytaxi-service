package postgres

import (
	"github.com/jmoiron/sqlx"
)

type clientRepo struct {
	db *sqlx.DB
}

// NewTaxiRepo ...
func NewClientRepo(db *sqlx.DB) *clientRepo {
	return &clientRepo{db: db}
}
