package dao

import (
	"database/sql"
	"fmt"
	"grpc_currency_converter/model"
	"log"
)

func GetCurrencyRate(db *sql.DB, currency string) (float64, error) {
	var rate float64
	err := db.QueryRow("SELECT rate FROM currencies WHERE code = $1", currency).Scan(&rate)
	if err != nil {
		return 0, fmt.Errorf("failed to fetch rate for %s: %w", currency, err)
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

func GetCurrency(code string) (*model.Currency, error) {
	db := GetDB()
	var currency model.Currency
	query := "SELECT code, rate FROM currencies WHERE code=$1"

	err := db.QueryRow(query, code).Scan(&currency.Code, &currency.Rate)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		log.Println("Error fetching currency:", err)
		return nil, err
	}

	return &currency, nil
}
