package dao

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"log"
	"strconv"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

type CurrencyDAO interface {
	GetConversionRate(from, to string) (float64, error)
	UpdateConversionRate(currency string, newRate float64) error
	GetAllRates() (map[string]float64, error)
}

type CurrencyDAOImpl struct {
	DB *dynamodb.Client
}

func NewCurrencyDAOImpl(db *dynamodb.Client) *CurrencyDAOImpl {
	return &CurrencyDAOImpl{DB: db}
}

func (dao *CurrencyDAOImpl) GetConversionRate(from, to string) (float64, error) {
	code := fmt.Sprintf("%s%s", from, to) // e.g., "USDINR"

	// Fetch rate from DynamoDB
	resp, err := dao.DB.GetItem(context.TODO(), &dynamodb.GetItemInput{
		TableName: aws.String("Currencies"),
		Key: map[string]types.AttributeValue{
			"code": &types.AttributeValueMemberS{Value: code},
		},
	})

	if err != nil {
		return 0, fmt.Errorf("failed to fetch conversion rate: %v", err)
	}

	// Check if item exists
	if resp.Item == nil {
		return 0, fmt.Errorf("conversion rate not found for: %s to %s", from, to)
	}

	// Extract rate
	rateAttr, ok := resp.Item["rate"].(*types.AttributeValueMemberN)
	if !ok {
		return 0, fmt.Errorf("rate attribute missing for %s", code)
	}

	rate, err := strconv.ParseFloat(rateAttr.Value, 64)
	if err != nil {
		return 0, fmt.Errorf("failed to parse rate: %v", err)
	}

	return rate, nil
}

func (dao *CurrencyDAOImpl) UpdateConversionRate(currency string, newRate float64) error {
	_, err := dao.DB.PutItem(context.TODO(), &dynamodb.PutItemInput{
		TableName: aws.String("Currencies"),
		Item: map[string]types.AttributeValue{
			"code": &types.AttributeValueMemberS{Value: currency},
			"rate": &types.AttributeValueMemberN{Value: fmt.Sprintf("%f", newRate)},
		},
	})

	if err != nil {
		return fmt.Errorf("failed to update rate for %s: %v", currency, err)
	}

	log.Printf("Updated %s to new rate: %.6f\n", currency, newRate)
	return nil
}

func (dao *CurrencyDAOImpl) GetAllRates() (map[string]float64, error) {
	resp, err := dao.DB.Scan(context.TODO(), &dynamodb.ScanInput{
		TableName: aws.String("Currencies"),
	})

	if err != nil {
		return nil, fmt.Errorf("failed to fetch rates: %v", err)
	}

	rates := make(map[string]float64)
	for _, item := range resp.Items {
		code := item["code"].(*types.AttributeValueMemberS).Value
		rate, _ := strconv.ParseFloat(item["rate"].(*types.AttributeValueMemberN).Value, 64)
		rates[code] = rate
	}

	return rates, nil
}

//package dao
//
//import (
//	"database/sql"
//	"fmt"
//)
//
//type CurrencyDAO interface {
//	GetConversionRate(from, to string) (float64, error)
//	UpdateConversionRate(currency string, newRate float64) error
//	GetAllRates() (map[string]float64, error)
//}
//
//type CurrencyDAOImpl struct {
//	DB *sql.DB
//}
//
//func NewCurrencyDAOImpl(db *sql.DB) *CurrencyDAOImpl {
//	return &CurrencyDAOImpl{DB: db}
//}
//
//func (dao *CurrencyDAOImpl) GetConversionRate(from, to string) (float64, error) {
//	var rate float64
//	code := fmt.Sprintf("%s%s", from, to) // e.g., "USDINR"
//
//	query := `SELECT rate FROM currencies WHERE code = $1`
//	err := dao.DB.QueryRow(query, code).Scan(&rate)
//
//	if err != nil {
//		if err == sql.ErrNoRows {
//			return 0, fmt.Errorf("conversion rate not found for: %s to %s", from, to)
//		}
//		return 0, fmt.Errorf("failed to fetch conversion rate: %v", err)
//	}
//
//	return rate, nil
//}
//
//func (dao *CurrencyDAOImpl) UpdateConversionRate(currency string, newRate float64) error {
//	query := `UPDATE currencies SET rate = $1 WHERE code = $2`
//	_, err := dao.DB.Exec(query, newRate, currency)
//	if err != nil {
//		return fmt.Errorf("failed to update rate for %s: %v", currency, err)
//	}
//
//	fmt.Printf("Updated %s to new rate: %.6f\n", currency, newRate)
//	return nil
//}
//
//func (dao *CurrencyDAOImpl) GetAllRates() (map[string]float64, error) {
//	rows, err := dao.DB.Query("SELECT code, rate FROM currencies")
//	if err != nil {
//		return nil, fmt.Errorf("failed to fetch rates: %v", err)
//	}
//	defer rows.Close()
//
//	rates := make(map[string]float64)
//	for rows.Next() {
//		var code string
//		var rate float64
//		if err := rows.Scan(&code, &rate); err != nil {
//			return nil, err
//		}
//		rates[code] = rate
//	}
//
//	return rates, nil
//}
