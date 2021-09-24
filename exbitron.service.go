package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

type RVLTicker struct {
	At     string `json:"at"`
	Ticker struct {
		Low                string      `json:"low"`
		High               string      `json:"high"`
		Open               string      `json:"open"`
		Last               string      `json:"last"`
		Volume             string      `json:"volume"`
		Amount             string      `json:"amount"`
		Vol                string      `json:"vol"`
		AvgPrice           string      `json:"avg_price"`
		PriceChangePercent string      `json:"price_change_percent"`
		At                 interface{} `json:"at"`
	} `json:"ticker"`
}

func fetchApi() RVLTicker {
	var api = "https://www.exbitron.com/api/v2/peatio/public/markets/rvlusdt/tickers"
	var t RVLTicker
	var unmarshalErr *json.UnmarshalTypeError

	// Let's request the API
	resp, er := http.Get(api)
	if er != nil {
		panic(er)
	}

	defer resp.Body.Close()

	// log the status code
	fmt.Println("Status Code: ", resp.Status)
	if resp.Status != "200 OK" {
		fmt.Println("Error Requesting API")
	}

	decoder := json.NewDecoder(resp.Body)
	decoder.DisallowUnknownFields()
	err := decoder.Decode(&t)
	if err != nil {
		if errors.As(err, &unmarshalErr) {
			errorResponse("Wrong Type error "+unmarshalErr.Field, resp.StatusCode)
		} else {
			errorResponse("Bad Request", resp.StatusCode)
		}
	}

	// fmt.Println(t.Ticker.AvgPrice)

	// sendWebhook(t)
	return t

}

func errorResponse(message string, httpStatusCode int) {
	fmt.Printf("[ERROR] [%v] - %s", httpStatusCode, message)
}
