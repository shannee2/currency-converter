package server

import (
	"log"
	"net"

	"grpc_currency_converter/dao"
	pb "grpc_currency_converter/proto"
	"grpc_currency_converter/service"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"google.golang.org/grpc"
)

// StartGrpcServer initializes the gRPC server with DynamoDB
func StartGrpcServer(db *dynamodb.Client, daoInstance dao.CurrencyDAO) {
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
