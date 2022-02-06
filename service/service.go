package service

import (
	l "github.com/abdullohsattorov/mytaxi-service/pkg/logger"
	"github.com/abdullohsattorov/mytaxi-service/storage"
)

// TaxiService is an object that implements TaxiServiceServer interface in genproto.
type MyTaxiService struct {
	storage storage.IStorage
	logger  l.Logger
}

// NewTaxiService ...
func NewMyTaxiService(storage storage.IStorage, log l.Logger) *MyTaxiService {
	return &MyTaxiService{
		storage: storage,
		logger:  log,
	}
}
