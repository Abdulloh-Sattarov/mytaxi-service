package service

import (
	"context"

	pb "github.com/abdullohsattorov/mytaxi-service/genproto"
	l "github.com/abdullohsattorov/mytaxi-service/pkg/logger"
	"github.com/gofrs/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (t *MyTaxiService) CreateClient(ctx context.Context, req *pb.Client) (*pb.Client, error) {
	id, err := uuid.NewV4()
	if err != nil {
		t.logger.Error("failed while generating uuid", l.Error(err))
		return nil, status.Error(codes.Internal, "failed generate uuid")
	}
	req.Id = id.String()

	client, err := t.storage.Client().CreateClient(*req)
	if err != nil {
		t.logger.Error("failed to create client", l.Error(err))
		return nil, status.Error(codes.Internal, "failed to create order")
	}

	return &client, nil
}

func (t *MyTaxiService) GetClient(ctx context.Context, req *pb.ByIdReq) (*pb.Client, error) {
	client, err := t.storage.Client().GetClient(req.GetId())
	if err != nil {
		t.logger.Error("failed to get client", l.Error(err))
		return nil, status.Error(codes.Internal, "failed to get order")
	}

	return &client, nil
}

func (t *MyTaxiService) UpdateClient(ctx context.Context, req *pb.Client) (*pb.Client, error) {
	client, err := t.storage.Client().UpdateClient(*req)
	if err != nil {
		t.logger.Error("failed to update client", l.Error(err))
		return nil, status.Error(codes.Internal, "failed to update order")
	}

	return &client, nil
}

func (t *MyTaxiService) DeleteClient(ctx context.Context, req *pb.ByIdReq) (*pb.EmptyResp, error) {
	err := t.storage.Client().DeleteClient(req.Id)
	if err != nil {
		t.logger.Error("failed to delete client", l.Error(err))
		return nil, status.Error(codes.Internal, "failed to delete order")
	}

	return &pb.EmptyResp{}, nil
}
