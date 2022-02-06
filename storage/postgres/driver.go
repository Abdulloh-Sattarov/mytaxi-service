package postgres

import (
	"github.com/jmoiron/sqlx"
)

type driverRepo struct {
	db *sqlx.DB
}

// NewTaxiRepo ...
func NewDriverRepo(db *sqlx.DB) *driverRepo {
	return &driverRepo{db: db}
}
