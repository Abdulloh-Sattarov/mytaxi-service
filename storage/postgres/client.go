package postgres

import (
	"database/sql"
	"time"

	"github.com/jmoiron/sqlx"

	pb "github.com/abdullohsattorov/mytaxi-service/genproto"
)

type clientRepo struct {
	db *sqlx.DB
}

// NewTaxiRepo ...
func NewClientRepo(db *sqlx.DB) *clientRepo {
	return &clientRepo{db: db}
}

func (t *clientRepo) CreateClient(client pb.Client) (pb.Client, error) {
	var id string
	err := t.db.QueryRow(`
        INSERT INTO clients(id, first_name, last_name, phone, created_at, updated_at)
        VALUES ($1, $2, $3, $4, $5, $6) returning id`, client.Id, client.FirstName, client.LastName, client.Phone, time.Now().UTC(), time.Now().UTC()).Scan(&id)
	if err != nil {
		return pb.Client{}, err
	}
	client, err = t.GetClient(id)

	if err != nil {
		return pb.Client{}, err
	}

	return client, nil
}

func (t *clientRepo) GetClient(id string) (pb.Client, error) {
	var client pb.Client

	err := t.db.QueryRow(`
        SELECT id, first_name, last_name, phone, created_at, updated_at FROM clients
        WHERE id=$1 and deleted_at is null`, id).Scan(
		&client.Id,
		&client.FirstName,
		&client.LastName,
		&client.Phone,
		&client.CreatedAt,
		&client.UpdatedAt,
	)
	if err != nil {
		return pb.Client{}, err
	}
	return client, nil
}

func (t *clientRepo) UpdateClient(client pb.Client) (pb.Client, error) {
	result, err := t.db.Exec(`UPDATE clients SET first_name=$2, last_name=$3, phone=$4, updated_at = $5 WHERE id=$1`,
		client.Id, client.FirstName, client.LastName, client.Phone, time.Now().UTC())
	if err != nil {
		return pb.Client{}, err
	}

	if i, _ := result.RowsAffected(); i == 0 {
		return pb.Client{}, sql.ErrNoRows
	}

	res, err := t.GetClient(client.Id)
	if err != nil {
		return pb.Client{}, err
	}

	return res, nil
}

func (t *clientRepo) DeleteClient(id string) error {
	result, err := t.db.Exec(`UPDATE clients SET deleted_at = $2 WHERE id=$1`, id, time.Now().UTC())
	if err != nil {
		return err
	}

	if i, _ := result.RowsAffected(); i == 0 {
		return sql.ErrNoRows
	}

	return nil
}
