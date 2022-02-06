package service

import (
	"context"

	pb "github.com/abdullohsattorov/mytaxi-service/genproto"
	l "github.com/abdullohsattorov/mytaxi-service/pkg/logger"
	"github.com/gofrs/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (t *MyTaxiService) CreateOrder(ctx context.Context, req *pb.OrderReq) (*pb.OrderRes, error) {
	id, err := uuid.NewV4()
	if err != nil {
		t.logger.Error("failed while generating uuid", l.Error(err))
		return nil, status.Error(codes.Internal, "failed generate uuid")
	}
	req.Id = id.String()

	fullorder, err := t.storage.Order().CreateOrder(*req)
	if err != nil {
		t.logger.Error("failed to create order", l.Error(err))
		return nil, status.Error(codes.Internal, "failed to create order")
	}

	return &fullorder, nil
}

func (t *MyTaxiService) GetOrder(ctx context.Context, req *pb.ByIdReq) (*pb.OrderRes, error) {
	order, err := t.storage.Order().GetOrder(req.GetId())
	if err != nil {
		t.logger.Error("failed to get order", l.Error(err))
		return nil, status.Error(codes.Internal, "failed to get order")
	}

	return &order, nil
}

func (t *MyTaxiService) UpdateOrder(ctx context.Context, req *pb.OrderReq) (*pb.OrderRes, error) {
	order, err := t.storage.Order().UpdateOrder(*req)
	if err != nil {
		t.logger.Error("failed to update order", l.Error(err))
		return nil, status.Error(codes.Internal, "failed to update order")
	}

	return &order, nil
}

func (t *MyTaxiService) DeleteOrder(ctx context.Context, req *pb.ByIdReq) (*pb.EmptyResp, error) {
	err := t.storage.Order().DeleteOrder(req.Id)
	if err != nil {
		t.logger.Error("failed to delete order", l.Error(err))
		return nil, status.Error(codes.Internal, "failed to delete order")
	}

	return &pb.EmptyResp{}, nil
}

func (t *MyTaxiService) ListOrders(ctx context.Context, req *pb.ListOrdersReq) (*pb.ListOrdersRes, error) {
	orders, count, err := t.storage.Order().ListOrders(req.ClientId, req.Page, req.Limit)
	if err != nil {
		t.logger.Error("failed to list orders", l.Error(err))
		return nil, status.Error(codes.Internal, "failed to list orders")
	}

	return &pb.ListOrdersRes{
		Orders: orders,
		Count:  count,
	}, nil
}
