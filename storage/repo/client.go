package repo

import (
	pb "github.com/abdullohsattorov/mytaxi-service/genproto"
)

// Taxi Storage interface

type ClientStorageI interface {
	CreateClient(pb.Client) (pb.Client, error)
	GetClient(id string) (pb.Client, error)
	UpdateClient(pb.Client) (pb.Client, error)
	DeleteClient(id string) error
}
