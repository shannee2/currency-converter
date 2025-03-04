package main

import (
	"grpc_currency_converter/consumer"
	"grpc_currency_converter/dao"
	"grpc_currency_converter/server"
	//"log"
)

func main() {
	db := dao.InitDB() // Initialize DynamoDB

	daoInstance := dao.NewCurrencyDAOImpl(db)

	// Start Kafka Consumer
	go consumer.StartKafkaConsumer("localhost:9092", "currency_updates", "currency-consumer-group", daoInstance)

	// Start gRPC Server
	go server.StartGrpcServer(db, daoInstance)

	// Start HTTP Gateway Server
	go server.StartHTTPGatewayServer()

	// Block main thread
	select {}
}
