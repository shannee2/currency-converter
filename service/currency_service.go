package service

import (
	"context"
	"database/sql"
	"fmt"
	"grpc_currency_converter/currency_converter/proto"
	"grpc_currency_converter/dao"
	//"grpc_currency_converter/model"
)

type CurrencyService struct {
	DB *sql.DB
}

func (s *CurrencyService) ConvertCurrency(ctx context.Context, req *proto.ConvertRequest) (*proto.ConvertResponse, error) {
	fromRate, err := dao.GetCurrencyRate(s.DB, req.Money.Currency)
	if err != nil {
		return nil, fmt.Errorf("error fetching from currency: %v", err)
	}

	toRate, err := dao.GetCurrencyRate(s.DB, req.ToCurrency)
	if err != nil {
		return nil, fmt.Errorf("error fetching to currency: %v", err)
	}

	convertedAmount := (req.Money.Amount / fromRate) * toRate

	fmt.Printf("Converting %.2f %s to %s\n", req.Money.Amount, req.Money.Currency, req.ToCurrency)

	return &proto.ConvertResponse{
		ConvertedMoney: &proto.Money{
			Amount:   convertedAmount,
			Currency: req.ToCurrency,
		},
	}, nil
}
