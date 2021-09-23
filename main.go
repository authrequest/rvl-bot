package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/DisgoOrg/disgohook"
	"github.com/DisgoOrg/disgohook/api"
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

func main() {
	// var access_key = "b3ce69ea07e684c3"
	// var secret_key = "005c5fd73fd10b09247d8f2f184a190f"
	fetchApi()
}

func sendWebhook() {
	webhook, err := disgohook.NewWebhookClientByToken(nil, nil, "750108819337511012/CDCJW-dfOvKx1Yv5QXrEA5Ykg1TxvVQHcVRis--E4TV-9mKGjzZ0VGY-8CHpoKmylp42")
	if err != nil {
		panic(err)
	}

	message, err := webhook.SendEmbeds(api.NewEmbedBuilder().SetTitle("Webhook Test").Build(),
		api.NewEmbedBuilder().SetThumbnail("https://ravencoinlite.info/wp-content/uploads/2021/09/RVL-transparent-bg.png").Build(),
		api.NewEmbedBuilder().SetField(0, "Ticker", "RVL/BTC", true).Build())
	fmt.Print(message)
}

func fetchApi() {
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

	fmt.Println(t.Ticker.AvgPrice)
	sendWebhook()

}

func errorResponse(message string, httpStatusCode int) {
	fmt.Println("[ERROR] [%+v]- %+v", httpStatusCode, message)
}
