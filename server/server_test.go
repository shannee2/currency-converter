package main

//
//import (
//	"context"
//	"log"
//	"net"
//	"testing"
//
//	"github.com/stretchr/testify/assert"
//	"github.com/stretchr/testify/mock"
//	"google.golang.org/grpc"
//	pb "grpc_currency_converter/currency_converter/proto"
//	"grpc_currency_converter/service"
//)
//
//// MockCurrencyDAO mocks the DAO layer
//type MockCurrencyDAO struct {
//	mock.Mock
//}
//
//func (m *MockCurrencyDAO) GetCurrencyRate(currency string) (float64, error) {
//	args := m.Called(currency)
//	return args.Get(0).(float64), args.Error(1)
//}
//
//func startTestServer(t *testing.T, mockDAO *MockCurrencyDAO) (pb.CurrencyServiceClient, func()) {
//	listener, err := net.Listen("tcp", ":0")
//	assert.NoError(t, err)
//
//	grpcServer := grpc.NewServer()
//	testService := service.NewCurrencyService(mockDAO, nil)
//	pb.RegisterCurrencyServiceServer(grpcServer, testService)
//
//	go func() {
//		if err := grpcServer.Serve(listener); err != nil {
//			log.Fatalf("Failed to serve test gRPC server: %v", err)
//		}
//	}()
//
//	conn, err := grpc.Dial(listener.Addr().String(), grpc.WithInsecure())
//	assert.NoError(t, err)
//
//	client := pb.NewCurrencyConverterClient(conn)
//
//	// Cleanup function
//	cleanup := func() {
//		grpcServer.Stop()
//		conn.Close()
//	}
//
//	return client, cleanup
//}
//
//func TestConvertCurrency(t *testing.T) {
//	mockDAO := new(MockCurrencyDAO)
//	client, cleanup := startTestServer(t, mockDAO)
//	defer cleanup()
//
//	mockDAO.On("GetCurrencyRate", "USD").Return(1.0, nil)
//	mockDAO.On("GetCurrencyRate", "INR").Return(83.0, nil)
//
//	req := &pb.ConvertRequest{
//		Money: &pb.Money{
//			Amount:   10.0,
//			Currency: "USD",
//		},
//		ToCurrency: "INR",
//	}
//
//	resp, err := client.ConvertCurrency(context.Background(), req)
//	assert.NoError(t, err)
//	assert.NotNil(t, resp)
//	assert.Equal(t, 830.0, resp.ConvertedMoney.Amount)
//	assert.Equal(t, "INR", resp.ConvertedMoney.Currency)
//
//	mockDAO.AssertExpectations(t)
//}
