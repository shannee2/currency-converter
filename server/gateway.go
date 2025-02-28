package server

import (
	"context"
	"log"
	"net/http"

	"grpc_currency_converter/proto"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func StartHTTPGatewayServer() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalln("Failed to dial server:", err)
	}

	mux := runtime.NewServeMux()
	err = proto.RegisterCurrencyServiceHandler(context.Background(), mux, conn)
	if err != nil {
		log.Fatalln("Failed to register HTTP gateway:", err)
	}

	log.Println("HTTP Gateway running on port 8090 (ConvertCurrency only)...")
	if err := http.ListenAndServe(":8090", mux); err != nil {
		log.Fatalf("Failed to start HTTP Gateway: %v", err)
	}
}
