package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// Define a struct for the response containing only the "result" field
type CurrencyResponse struct {
	Result float64 `json:"result"`
}

func main() {
	url := "http://apilayer.net/api/convert?access_key=8bbb3ae162150cfab8f05beb07603221&from=USD&to=EUR&amount=10"

	// Make the GET request
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error making request:", err)
		return
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response:", err)
		return
	}

	// Parse the JSON response
	var currencyResp CurrencyResponse
	err = json.Unmarshal(body, &currencyResp)
	if err != nil {
		fmt.Println("Error parsing JSON:", err)
		return
	}

	// Print the extracted result
	fmt.Printf("Converted Amount: %.2f\n", currencyResp.Result+1)
}
