package main

import (
	"log"
	"net"

	pb "grpc_currency_converter/currency_converter/proto"
	"grpc_currency_converter/dao"
	"grpc_currency_converter/service"

	"google.golang.org/grpc"
)

func main() {
	db, err := dao.InitDB()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	daoInstance := dao.NewCurrencyDAOImpl(db)

	currencyService := service.NewCurrencyService(daoInstance, db)

	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen on port 50051: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterCurrencyConverterServer(grpcServer, currencyService)

	log.Println("gRPC server running on port 50051...")
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
