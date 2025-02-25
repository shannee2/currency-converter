package main

import (
	"context"
	"grpc_currency_converter/currency_converter/proto"
	"grpc_currency_converter/dao"
	"grpc_currency_converter/service"
	"log"
	"net"

	_ "github.com/lib/pq"
	"google.golang.org/grpc"
)

type server struct {
	proto.UnimplementedCurrencyConverterServer
	Service *service.CurrencyService
}

func (s *server) ConvertCurrency(ctx context.Context, req *proto.ConvertRequest) (*proto.ConvertResponse, error) {
	return s.Service.ConvertCurrency(ctx, req)
}

func main() {
	db, err := dao.InitDB()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Inject dependencies
	service := &service.CurrencyService{DB: db}

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	proto.RegisterCurrencyConverterServer(s, &server{Service: service})

	log.Println("gRPC server running on port 50051...")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
