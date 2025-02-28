package server

import (
	"database/sql"
	"log"
	"net"

	"grpc_currency_converter/dao"
	pb "grpc_currency_converter/proto"
	"grpc_currency_converter/service"

	"google.golang.org/grpc"
)

func StartGrpcServer(db *sql.DB, daoInstance dao.CurrencyDAO) {
	currencyService := service.NewCurrencyService(daoInstance, db)
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen on port 50051: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterCurrencyServiceServer(grpcServer, currencyService)

	log.Println("gRPC server running on port 50051...")
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
