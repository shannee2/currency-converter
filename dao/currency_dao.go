package dao

import (
	"database/sql"
	"errors"
	"fmt"
	"grpc_currency_converter/model"
	"log"
)

type CurrencyDAO interface {
	GetCurrencyRate(currency string) (float64, error)
}

type CurrencyDAOImpl struct {
	DB *sql.DB
}

func NewCurrencyDAOImpl(db *sql.DB) *CurrencyDAOImpl {
	return &CurrencyDAOImpl{DB: db}
}

func (dao *CurrencyDAOImpl) GetCurrencyRate(currency string) (float64, error) {
	var rate float64
	query := "SELECT rate FROM currencies WHERE code = $1"

	err := dao.DB.QueryRow(query, currency).Scan(&rate)
	if err != nil {
		if err == sql.ErrNoRows {
			return 0, fmt.Errorf("currency rate not found for: %s", currency)
		}
		return 0, fmt.Errorf("failed to fetch currency rate: %v", err)
	}

	return rate, nil
}

func AddCurrency(currency model.Currency) error {
	db := GetDB()
	query := "INSERT INTO currencies (code, rate) VALUES ($1, $2)"
	_, err := db.Exec(query, currency.Code, currency.Rate)
	if err != nil {
		log.Println("Error inserting currency:", err)
		return err
	}
	return nil
}

// Retrieve currency details
func GetCurrency(code string) (*model.Currency, error) {
	db := GetDB()
	var currency model.Currency
	query := "SELECT code, rate FROM currencies WHERE code=$1"

	err := db.QueryRow(query, code).Scan(&currency.Code, &currency.Rate)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		log.Println("Error fetching currency:", err)
		return nil, err
	}

	return &currency, nil
}
