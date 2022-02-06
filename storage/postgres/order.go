package postgres

import (
	"github.com/jmoiron/sqlx"
)

type orderRepo struct {
	db *sqlx.DB
}

// NewTaxiRepo ...
func NewOrderRepo(db *sqlx.DB) *orderRepo {
	return &orderRepo{db: db}
}
