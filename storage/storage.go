package storage

import (
	"github.com/jmoiron/sqlx"

	"github.com/abdullohsattorov/mytaxi-service/storage/postgres"
	"github.com/abdullohsattorov/mytaxi-service/storage/repo"
)

// IStorage ...
type IStorage interface {
	Client() repo.ClientStorageI
	Driver() repo.DriverStorageI
	Order() repo.OrderStorageI
}

type storagePg struct {
	db         *sqlx.DB
	clientRepo repo.DriverStorageI
	driverRepo repo.ClientStorageI
	orderRepo  repo.OrderStorageI
}

// NewStoragePg ...
func NewStoragePg(db *sqlx.DB) *storagePg {
	return &storagePg{
		db:         db,
		clientRepo: postgres.NewClientRepo(db),
		driverRepo: postgres.NewDriverRepo(db),
		orderRepo:  postgres.NewOrderRepo(db),
	}
}

func (s storagePg) Client() repo.ClientStorageI {
	return s.clientRepo
}

func (s storagePg) Driver() repo.DriverStorageI {
	return s.driverRepo
}

func (s storagePg) Order() repo.OrderStorageI {
	return s.orderRepo
}
