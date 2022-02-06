package repo

import (
	pb "github.com/abdullohsattorov/mytaxi-service/genproto"
)

// Taxi Storage interface

type DriverStorageI interface {
	CreateDriver(pb.Driver) (pb.Driver, error)
	GetDriver(id string) (pb.Driver, error)
	UpdateDriver(pb.Driver) (pb.Driver, error)
	DeleteDriver(id string) error
}
