package service

//
//import (
//	"context"
//	"errors"
//	"grpc_currency_converter/currency_converter/proto"
//	"testing"
//
//	"github.com/stretchr/testify/assert"
//	"github.com/stretchr/testify/mock"
//)
//
//type MockCurrencyDAO struct {
//	mock.Mock
//}
//
//func (m *MockCurrencyDAO) GetCurrencyRate(currency string) (float64, error) {
//	args := m.Called(currency)
//	return args.Get(0).(float64), args.Error(1)
//}
//
//func TestConvertCurrency_Success(t *testing.T) {
//	mockDAO := new(MockCurrencyDAO)
//	service := &CurrencyService{DAO: mockDAO}
//
//	mockDAO.On("GetCurrencyRate", "USD").Return(1.0, nil)
//	mockDAO.On("GetCurrencyRate", "INR").Return(82.0, nil)
//
//	req := &proto.ConvertRequest{
//		Money: &proto.Money{
//			Amount:   100,
//			Currency: "USD",
//		},
//		ToCurrency: "INR",
//	}
//
//	resp, err := service.ConvertCurrency(context.Background(), req)
//
//	assert.NoError(t, err)
//	assert.NotNil(t, resp)
//	assert.Equal(t, 8200.0, resp.ConvertedMoney.Amount)
//	assert.Equal(t, "INR", resp.ConvertedMoney.Currency)
//
//	mockDAO.AssertExpectations(t)
//}
//
//func TestConvertCurrency_SourceCurrencyNotFound(t *testing.T) {
//	mockDAO := new(MockCurrencyDAO)
//	service := &CurrencyService{DAO: mockDAO}
//
//	mockDAO.On("GetCurrencyRate", "USD").Return(0.0, errors.New("currency not found"))
//
//	req := &proto.ConvertRequest{
//		Money: &proto.Money{
//			Amount:   100,
//			Currency: "USD",
//		},
//		ToCurrency: "INR",
//	}
//
//	resp, err := service.ConvertCurrency(context.Background(), req)
//
//	assert.Error(t, err)
//	assert.Nil(t, resp)
//	assert.Contains(t, err.Error(), "error fetching from currency")
//
//	mockDAO.AssertExpectations(t)
//}
//
//func TestConvertCurrency_TargetCurrencyNotFound(t *testing.T) {
//	mockDAO := new(MockCurrencyDAO)
//	service := &CurrencyService{DAO: mockDAO}
//
//	mockDAO.On("GetCurrencyRate", "USD").Return(1.0, nil)
//	mockDAO.On("GetCurrencyRate", "INR").Return(0.0, errors.New("currency not found"))
//
//	req := &proto.ConvertRequest{
//		Money: &proto.Money{
//			Amount:   100,
//			Currency: "USD",
//		},
//		ToCurrency: "INR",
//	}
//
//	resp, err := service.ConvertCurrency(context.Background(), req)
//
//	assert.Error(t, err)
//	assert.Nil(t, resp)
//	assert.Contains(t, err.Error(), "error fetching to currency")
//
//	mockDAO.AssertExpectations(t)
//}
