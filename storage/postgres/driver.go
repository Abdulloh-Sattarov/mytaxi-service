package postgres

import (
	"database/sql"
	"time"

	"github.com/jmoiron/sqlx"

	pb "github.com/abdullohsattorov/mytaxi-service/genproto"
)

type driverRepo struct {
	db *sqlx.DB
}

// NewTaxiRepo ...
func NewDriverRepo(db *sqlx.DB) *driverRepo {
	return &driverRepo{db: db}
}

func (t *driverRepo) CreateDriver(driver pb.Driver) (pb.Driver, error) {
	var id string
	err := t.db.QueryRow(`
        INSERT INTO drivers(id, first_name, last_name, phone, car_model, created_at, updated_at)
        VALUES ($1, $2, $3, $4, $5, $6, $7) returning id`, driver.Id, driver.FirstName, driver.LastName, driver.Phone, driver.CarModel, time.Now().UTC(), time.Now().UTC()).Scan(&id)
	if err != nil {
		return pb.Driver{}, err
	}
	driver, err = t.GetDriver(id)

	if err != nil {
		return pb.Driver{}, err
	}

	return driver, nil
}

func (t *driverRepo) GetDriver(id string) (pb.Driver, error) {
	var driver pb.Driver

	err := t.db.QueryRow(`
        SELECT id, first_name, last_name, phone, car_model, created_at, updated_at FROM drivers
        WHERE id=$1 and deleted_at is null`, id).Scan(
		&driver.Id,
		&driver.FirstName,
		&driver.LastName,
		&driver.Phone,
		&driver.CarModel,
		&driver.CreatedAt,
		&driver.UpdatedAt,
	)
	if err != nil {
		return pb.Driver{}, err
	}
	return driver, nil
}

func (t *driverRepo) UpdateDriver(driver pb.Driver) (pb.Driver, error) {
	result, err := t.db.Exec(`UPDATE drivers SET first_name=$2, last_name=$3, phone=$4, car_model=$5, updated_at=$6 WHERE id=$1`,
		driver.Id, driver.FirstName, driver.LastName, driver.Phone, driver.CarModel, time.Now().UTC())
	if err != nil {
		return pb.Driver{}, err
	}

	if i, _ := result.RowsAffected(); i == 0 {
		return pb.Driver{}, sql.ErrNoRows
	}

	driver, err = t.GetDriver(driver.Id)
	if err != nil {
		return pb.Driver{}, err
	}

	return driver, nil
}

func (t *driverRepo) DeleteDriver(id string) error {
	result, err := t.db.Exec(`UPDATE drivers SET deleted_at = $2 WHERE id=$1`, id, time.Now().UTC())
	if err != nil {
		return err
	}

	if i, _ := result.RowsAffected(); i == 0 {
		return sql.ErrNoRows
	}

	return nil
}
