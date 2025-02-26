package service

import (
	"context"
	"database/sql"
	"fmt"
	"grpc_currency_converter/currency_converter/proto"
	"grpc_currency_converter/dao"
)

type CurrencyService struct {
	proto.UnimplementedCurrencyConverterServer
	DAO dao.CurrencyDAO
	DB  *sql.DB
}

func NewCurrencyService(dao dao.CurrencyDAO, db *sql.DB) *CurrencyService {
	return &CurrencyService{DAO: dao, DB: db}
}

func (s *CurrencyService) ConvertCurrency(ctx context.Context, req *proto.ConvertRequest) (*proto.ConvertResponse, error) {
	fromRate, err := s.DAO.GetCurrencyRate(req.Money.Currency)
	if err != nil {
		return nil, fmt.Errorf("error fetching from currency: %v", err)
	}

	toRate, err := s.DAO.GetCurrencyRate(req.ToCurrency)
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
