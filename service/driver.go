package service

import (
	l "github.com/abdullohsattorov/mytaxi-service/pkg/logger"
	"github.com/abdullohsattorov/mytaxi-service/storage"
)

// TaxiService is an object that implements TaxiServiceServer interface in genproto.
type DriverService struct {
	storage storage.IStorage
	logger  l.Logger
}

// NewTaxiService ...
func NewDriverService(storage storage.IStorage, log l.Logger) *DriverService {
	return &DriverService{
		storage: storage,
		logger:  log,
	}
}
