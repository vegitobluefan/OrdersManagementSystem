package main

import (
	"context"
	"log"

	pb "github.com/vegitobluefan/OrdersManagementSystem/genproto/stock"
)

type StockServer struct {
	pb.UnimplementedStockServiceServer
}

func (s *StockServer) AddItem(ctx context.Context, req *pb.StockRequest) (*pb.StockResponse, error) {
	item := AddOrUpdateItem(req.Name, req.Quantity)
	log.Printf("Item added/updated: %s (%d)", item.Name, item.Quantity)

	return &pb.StockResponse{
		Name:     item.Name,
		Quantity: item.Quantity,
	}, nil
}

func (s *StockServer) ListStock(ctx context.Context, _ *pb.Empty) (*pb.StockList, error) {
	items := ListItems()
	var list []*pb.StockResponse

	for _, item := range items {
		list = append(list, &pb.StockResponse{
			Name:     item.Name,
			Quantity: item.Quantity,
		})
	}
	return &pb.StockList{Items: list}, nil
}

func (s *StockServer) UseItem(ctx context.Context, req *pb.StockRequest) (*pb.UseItemResponse, error) {
	ok := DecreaseItem(req.Name, req.Quantity)
	status := "used"
	if !ok {
		status = "not_enough"
	}
	log.Printf("Use item %s (%d): %s", req.Name, req.Quantity, status)

	return &pb.UseItemResponse{
		Status: status,
	}, nil
}
