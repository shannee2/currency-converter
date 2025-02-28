package service

import (
	"context"
	"database/sql"
	"fmt"
	"grpc_currency_converter/dao"
	"grpc_currency_converter/proto"
)

type CurrencyService struct {
	proto.UnimplementedCurrencyServiceServer
	DAO dao.CurrencyDAO
	DB  *sql.DB
}

func NewCurrencyService(dao dao.CurrencyDAO, db *sql.DB) *CurrencyService {
	return &CurrencyService{DAO: dao, DB: db}
}

type CurrencyResponse struct {
	Result float64 `json:"result"`
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

//
//func (s *CurrencyService) ConvertCurrency(ctx context.Context, req *proto.ConvertRequest) (*proto.ConvertResponse, error) {
//	convertedAmount := getConvertedAmount(req)
//
//	fmt.Printf("Converted %.2f %s to %s: %f\n", req.Money.Amount, req.Money.Currency, req.ToCurrency, convertedAmount)
//
//	return &proto.ConvertResponse{
//		ConvertedMoney: &proto.Money{
//			Amount:   convertedAmount,
//			Currency: req.ToCurrency,
//		},
//	}, nil
//}

//func getConvertedAmount(req *proto.ConvertRequest) (convertedAmount float64) {
//	const baseURL = "http://apilayer.net/api/convert?access_key=8bbb3ae162150cfab8f05beb07603221"
//	url := fmt.Sprintf("%s&from=%s&to=%s&amount=%f", baseURL, req.Money.Currency, req.ToCurrency, req.Money.Amount)
//
//	// Make the GET request
//	resp, err := http.Get(url)
//	if err != nil {
//		fmt.Println("Error making request:", err)
//		return
//	}
//	defer resp.Body.Close()
//
//	// Read the response body
//	body, err := io.ReadAll(resp.Body)
//	if err != nil {
//		fmt.Println("Error reading response:", err)
//		return
//	}
//
//	// Parse the JSON response
//	var currencyResp CurrencyResponse
//	err = json.Unmarshal(body, &currencyResp)
//	if err != nil {
//		fmt.Println("Error parsing JSON:", err)
//		return
//	}
//
//	return currencyResp.Result
//}
