package repo

import (
	pb "github.com/abdullohsattorov/mytaxi-service/genproto"
)

// Taxi Storage interface

type OrderStorageI interface {
	CreateOrder(pb.OrderReq) (pb.OrderRes, error)
	GetOrder(id string) (pb.OrderRes, error)
	ListOrders(clientId string, page, limit int64) ([]*pb.OrderRes, int64, error)
	UpdateOrder(pb.OrderReq) (pb.OrderRes, error)
	DeleteOrder(id string) error
}
