package main

import (
	"log"
	"net"

	pb "github.com/vegitobluefan/OrdersManagementSystem/genproto/stock"
	"github.com/vegitobluefan/OrdersManagementSystem/stock"

	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", ":50054")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterStockServiceServer(s, &stock.StockServer{})

	log.Println("Stock gRPC server listening on :50054")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
