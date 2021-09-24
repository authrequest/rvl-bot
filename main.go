package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

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
	run()
}

func run() {
	for range time.Tick(time.Second * 10) {
		hashrate := getHashrate()
		difficulty := getDifficulty()
		t := fetchApi()
		sendWebhook(t, hashrate, difficulty)
	}
}
func getHashrate() string {
	// Set
	req, err := http.NewRequest("GET", "http://explorer.ravencoinlite.org/api/getnetworkhashps", nil)
	if err != nil {
		log.Fatalln(err)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	hashrate := strings.Trim(string(b), "\"")
	s, err := strconv.ParseFloat(hashrate, 32)
	if err != nil {
		log.Fatalln(err)
	}

	return strconv.FormatFloat(s/1000000000, 'f', 3, 32) + " GH/s"

}

func getDifficulty() string {
	// Set
	req, err := http.NewRequest("GET", "http://explorer.ravencoinlite.org/api/getdifficulty", nil)
	if err != nil {
		log.Fatalln(err)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	difficulty := strings.Trim(string(b), "\"")
	return difficulty
}

func sendWebhook(t RVLTicker, hashrate string, difficulty string) {
	webhook, err := disgohook.NewWebhookClientByToken(nil, nil, "750108819337511012/CDCJW-dfOvKx1Yv5QXrEA5Ykg1TxvVQHcVRis--E4TV-9mKGjzZ0VGY-8CHpoKmylp42")
	if err != nil {
		panic(err)
	}

	var bool = true
	message, err := webhook.SendEmbeds(
		api.NewEmbedBuilder().
			SetTitle("RVL Exchange Price").
			SetURL("https://www.exbitron.com/markets/rvlusdt").
			SetThumbnail("https://ravencoinlite.info/wp-content/uploads/2021/09/RVL-transparent-bg.png").
			SetFields(&api.EmbedField{
				Name:   "Exchange",
				Value:  "Exbitron",
				Inline: &bool,
			}, &api.EmbedField{
				Name:   "Average Price",
				Value:  t.Ticker.AvgPrice,
				Inline: &bool,
			}, &api.EmbedField{
				Name:   "Price Change",
				Value:  t.Ticker.PriceChangePercent,
				Inline: &bool,
			}, &api.EmbedField{
				Name:   "High",
				Value:  t.Ticker.High,
				Inline: &bool,
			}, &api.EmbedField{
				Name:   "Low",
				Value:  t.Ticker.Low,
				Inline: &bool,
			}, &api.EmbedField{
				Name:   "Volume",
				Value:  t.Ticker.Volume,
				Inline: &bool,
			}).
			SetFooter("LayersTech Exchange Go Monitor", "https://ravencoinlite.info/wp-content/uploads/2021/09/RVL-transparent-bg.png").
			Build())

	fmt.Print(message)
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
