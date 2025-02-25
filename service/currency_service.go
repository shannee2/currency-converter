package service

import (
	"context"
	"database/sql"
	"fmt"
	"grpc_currency_converter/currency_converter/proto"
	"grpc_currency_converter/dao"
)

type CurrencyService struct {
	DB *sql.DB
}

// ConvertCurrency handles the business logic for currency conversion
func (s *CurrencyService) ConvertCurrency(ctx context.Context, req *proto.ConvertRequest) (*proto.ConvertResponse, error) {

	fromRate, err := dao.GetCurrencyRate(s.DB, req.FromCurrency)
	if err != nil {
		return nil, fmt.Errorf("error fetching from currency: %v", err)
	}

	toRate, err := dao.GetCurrencyRate(s.DB, req.ToCurrency)
	if err != nil {
		return nil, fmt.Errorf("error fetching to currency: %v", err)
	}

	convertedAmount := (req.Amount / fromRate) * toRate
	fmt.Println("Converted amount:", convertedAmount)
	return &proto.ConvertResponse{ConvertedAmount: convertedAmount}, nil
}
