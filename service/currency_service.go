package service

import (
	"context"
	"fmt"
	"grpc_currency_converter/dao"
	"grpc_currency_converter/proto"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

type CurrencyService struct {
	proto.UnimplementedCurrencyServiceServer
	DAO dao.CurrencyDAO
	DB  *dynamodb.Client
}

func NewCurrencyService(dao dao.CurrencyDAO, db *dynamodb.Client) *CurrencyService {
	return &CurrencyService{DAO: dao, DB: db}
}

func (s *CurrencyService) ConvertCurrency(ctx context.Context, req *proto.ConvertRequest) (*proto.ConvertResponse, error) {
	fmt.Println(req.GetMoney().Currency, req.ToCurrency)
	conversionRate, err := s.DAO.GetConversionRate(req.Money.Currency, req.ToCurrency)
	if err != nil {
		return nil, fmt.Errorf("error fetching conversion rate from %s to %s: %v", req.Money.Currency, req.ToCurrency, err)
	}

	convertedAmount := req.Money.Amount * conversionRate

	fmt.Printf("Converting %.2f %s to %s at rate %.6f\n", req.Money.Amount, req.Money.Currency, req.ToCurrency, conversionRate)

	return &proto.ConvertResponse{ConvertedMoney: convertedAmount}, nil
}

func (s *CurrencyService) GetAllRates(ctx context.Context, req *proto.Empty) (*proto.AllRatesResponse, error) {
	rates, err := s.DAO.GetAllRates()
	if err != nil {
		return nil, fmt.Errorf("error fetching all rates: %v", err)
	}
	return &proto.AllRatesResponse{Rates: rates}, nil
}
