package service

import (
	"context"

	"github.com/gofrs/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	pb "github.com/abdullohsattorov/mytaxi-service/genproto"
	l "github.com/abdullohsattorov/mytaxi-service/pkg/logger"
)

func (t *MyTaxiService) CreateDriver(ctx context.Context, req *pb.Driver) (*pb.Driver, error) {
	id, err := uuid.NewV4()
	if err != nil {
		t.logger.Error("failed while generating uuid", l.Error(err))
		return nil, status.Error(codes.Internal, "failed generate uuid")
	}
	req.Id = id.String()

	driver, err := t.storage.Driver().CreateDriver(*req)
	if err != nil {
		t.logger.Error("failed to create order", l.Error(err))
		return nil, status.Error(codes.Internal, "failed to create order")
	}

	return &driver, nil
}

func (t *MyTaxiService) GetDriver(ctx context.Context, req *pb.ByIdReq) (*pb.Driver, error) {
	driver, err := t.storage.Driver().GetDriver(req.GetId())
	if err != nil {
		t.logger.Error("failed to get task", l.Error(err))
		return nil, status.Error(codes.Internal, "failed to get order")
	}

	return &driver, nil
}

func (t *MyTaxiService) UpdateDriver(ctx context.Context, req *pb.Driver) (*pb.Driver, error) {
	driver, err := t.storage.Driver().UpdateDriver(*req)
	if err != nil {
		t.logger.Error("failed to update driver", l.Error(err))
		return nil, status.Error(codes.Internal, "failed to update order")
	}

	return &driver, nil
}

func (t *MyTaxiService) DeleteDriver(ctx context.Context, req *pb.ByIdReq) (*pb.EmptyResp, error) {
	err := t.storage.Driver().DeleteDriver(req.Id)
	if err != nil {
		t.logger.Error("failed to delete driver", l.Error(err))
		return nil, status.Error(codes.Internal, "failed to delete order")
	}

	return &pb.EmptyResp{}, nil
}
