package service

import (
	l "github.com/abdullohsattorov/mytaxi-service/pkg/logger"
	"github.com/abdullohsattorov/mytaxi-service/storage"
)

// TaxiService is an object that implements TaxiServiceServer interface in genproto.
type OrderService struct {
	storage storage.IStorage
	logger  l.Logger
}

// NewTaxiService ...
func NewOrderService(storage storage.IStorage, log l.Logger) *OrderService {
	return &OrderService{
		storage: storage,
		logger:  log,
	}
}

