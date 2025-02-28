package dao

import (
	"database/sql"
	"fmt"
)

type CurrencyDAO interface {
	GetConversionRate(from, to string) (float64, error)
	UpdateConversionRate(currency string, newRate float64) error
	GetAllRates() (map[string]float64, error)
}

type CurrencyDAOImpl struct {
	DB *sql.DB
}

func NewCurrencyDAOImpl(db *sql.DB) *CurrencyDAOImpl {
	return &CurrencyDAOImpl{DB: db}
}

func (dao *CurrencyDAOImpl) GetConversionRate(from, to string) (float64, error) {
	var rate float64
	code := fmt.Sprintf("%s%s", from, to) // e.g., "USDINR"

	query := `SELECT rate FROM currencies WHERE code = $1`
	err := dao.DB.QueryRow(query, code).Scan(&rate)

	if err != nil {
		if err == sql.ErrNoRows {
			return 0, fmt.Errorf("conversion rate not found for: %s to %s", from, to)
		}
		return 0, fmt.Errorf("failed to fetch conversion rate: %v", err)
	}

	return rate, nil
}

func (dao *CurrencyDAOImpl) UpdateConversionRate(currency string, newRate float64) error {
	query := `UPDATE currencies SET rate = $1 WHERE code = $2`
	_, err := dao.DB.Exec(query, newRate, currency)
	if err != nil {
		return fmt.Errorf("failed to update rate for %s: %v", currency, err)
	}

	fmt.Printf("Updated %s to new rate: %.6f\n", currency, newRate)
	return nil
}

func (dao *CurrencyDAOImpl) GetAllRates() (map[string]float64, error) {
	rows, err := dao.DB.Query("SELECT code, rate FROM currencies")
	if err != nil {
		return nil, fmt.Errorf("failed to fetch rates: %v", err)
	}
	defer rows.Close()

	rates := make(map[string]float64)
	for rows.Next() {
		var code string
		var rate float64
		if err := rows.Scan(&code, &rate); err != nil {
			return nil, err
		}
		rates[code] = rate
	}

	return rates, nil
}
