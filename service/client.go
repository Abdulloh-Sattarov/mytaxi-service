package service

import (
	"context"

	l "github.com/abdullohsattorov/mytaxi-service/pkg/logger"
	"github.com/abdullohsattorov/mytaxi-service/storage"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// TaxiService is an object that implements TaxiServiceServer interface in genproto.
type ClientService struct {
	storage storage.IStorage
	logger  l.Logger
}

// NewTaxiService ...
func NewClientService(storage storage.IStorage, log l.Logger) *ClientService {
	return &ClientService{
		storage: storage,
		logger:  log,
	}
}

func (t *ClientService) CreateClient(ctx context.Context, req *pb.Client) (*pb.Client, error) {
	id, err := uuid.NewV4()
	if err != nil {
		t.logger.Error("failed while generating uuid", l.Error(err))
		return nil, status.Error(codes.Internal, "failed generate uuid")
	}
	req.Id = id.String()

	client, err := t.storage.Taxi().CreateClient(*req)
	if err != nil {
		t.logger.Error("failed to create client", l.Error(err))
		return nil, status.Error(codes.Internal, "failed to create order")
	}

	return &client, nil
}

func (t *ClientService) GetClient(ctx context.Context, req *pb.ByIdReq) (*pb.Client, error) {
	client, err := t.storage.Taxi().GetClient(req.GetId())
	if err != nil {
		t.logger.Error("failed to get client", l.Error(err))
		return nil, status.Error(codes.Internal, "failed to get order")
	}

	return &client, nil
}

func (t *ClientService) UpdateClient(ctx context.Context, req *pb.Client) (*pb.Client, error) {
	client, err := t.storage.Taxi().UpdateClient(*req)
	if err != nil {
		t.logger.Error("failed to update client", l.Error(err))
		return nil, status.Error(codes.Internal, "failed to update order")
	}

	return &client, nil
}

func (t *ClientService) DeleteClient(ctx context.Context, req *pb.ByIdReq) (*pb.EmptyResp, error) {
	err := t.storage.Taxi().DeleteClient(req.Id)
	if err != nil {
		t.logger.Error("failed to delete client", l.Error(err))
		return nil, status.Error(codes.Internal, "failed to delete order")
	}

	return &pb.EmptyResp{}, nil
}
