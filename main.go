//package main
//
//import (
//	"google.golang.org/grpc"
//	"google.golang.org/grpc/credentials/insecure"
//	"grpc_currency_converter/consumer"
//	"grpc_currency_converter/dao"
//	pb "grpc_currency_converter/proto"
//	"grpc_currency_converter/server"
//	"log"
//	"net"
//	"net/http"
//)
//
//func main() {
//	db, err := dao.InitDB()
//	if err != nil {
//		log.Fatalf("Failed to connect to database: %v", err)
//	}
//
//	daoInstance := dao.NewCurrencyDAOImpl(db)
//	go consumer.StartKafkaConsumer("localhost:9092", "currency_updates", "currency-consumer-group", daoInstance)
//	go server.StartGrpcServer(db, daoInstance)
//	//go server.StartHTTPGatewayServer()
//
//	// Block main thread
//	select {}
//}

package main

import (
	"grpc_currency_converter/consumer"
	"grpc_currency_converter/dao"
	"grpc_currency_converter/server"
	"log"
)

func main() {
	db, err := dao.InitDB()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

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

// Working find independently:
//package main
//
//import (
//	"context"
//	"log"
//	"net"
//	"net/http"
//
//	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
//	"google.golang.org/grpc"
//	"google.golang.org/grpc/credentials/insecure"
//
//	pb "grpc_currency_converter/proto"
//)
//
//type server struct {
//	pb.UnimplementedCurrencyServiceServer
//}
//
//func (s *server) ConvertCurrency(ctx context.Context, req *pb.ConvertRequest) (*pb.ConvertResponse, error) {
//	// Dummy conversion logic
//	rate := 1.2 // Example conversion rate
//	convertedAmount := req.Money.Amount * rate
//	return &pb.ConvertResponse{ConvertedMoney: convertedAmount}, nil
//}
//
//func main() {
//	// Start gRPC Server
//	lis, err := net.Listen("tcp", ":5000")
//	if err != nil {
//		log.Fatalf("Failed to listen: %v", err)
//	}
//
//	s := grpc.NewServer()
//	pb.RegisterCurrencyServiceServer(s, &server{})
//
//	log.Println("Serving gRPC on 0.0.0.0:5000")
//	go func() {
//		log.Fatal(s.Serve(lis))
//	}()
//
//	// Start gRPC-Gateway
//	conn, err := grpc.DialContext(
//		context.Background(),
//		"localhost:5000", // Use localhost instead of 0.0.0.0
//		grpc.WithTransportCredentials(insecure.NewCredentials()),
//	)
//	if err != nil {
//		log.Fatalf("Failed to dial server: %v", err)
//	}
//
//	gwmux := runtime.NewServeMux()
//	err = pb.RegisterCurrencyServiceHandler(context.Background(), gwmux, conn)
//	if err != nil {
//		log.Fatalf("Failed to register gateway: %v", err)
//	}
//
//	gwServer := &http.Server{
//		Addr:    ":6000",
//		Handler: gwmux,
//	}
//
//	log.Println("Serving gRPC-Gateway on http://localhost:6000") // Fixed log message
//	log.Fatal(gwServer.ListenAndServe())
//}
