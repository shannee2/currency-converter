package main

//
//import (
//	"fmt"
//	"grpc_currency_converter/dao"
//	"grpc_currency_converter/model"
//	"log"
//)
//
//func main() {
//	_, err := dao.InitDB()
//	if err != nil {
//		log.Fatal("Failed to initialize DB:", err)
//	}
//
//	// Adding a currency
//	newCurrency := model.Currency{Code: "GBP", Rate: 0.78}
//	err = dao.AddCurrency(newCurrency)
//	if err != nil {
//		log.Println("Failed to add currency:", err)
//	}
//
//	//"USD": 1.0,
//	//	"INR": 83.0,
//	//	"EUR": 0.91,
//	//	"GBP": 0.78,
//	//	"JPY": 150.5,
//
//	// Fetching currency
//	currency, err := dao.GetCurrency("USD")
//	if err != nil {
//		log.Println("Error fetching currency:", err)
//	} else if currency == nil {
//		fmt.Println("Currency not found")
//	} else {
//		fmt.Printf("Currency: %s, Rate: %.2f\n", currency.Code, currency.Rate)
//	}
//}
