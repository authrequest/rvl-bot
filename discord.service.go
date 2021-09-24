package main

import (
	"strconv"

	"github.com/DisgoOrg/disgohook"
	"github.com/DisgoOrg/disgohook/api"
	"github.com/dustin/go-humanize"
)

func sendWebhook(t RVLTicker, hashrate string, difficulty string, supply string, addressInfo GetAddress) {
	webhook, err := disgohook.NewWebhookClientByToken(nil, nil, "750108819337511012/CDCJW-dfOvKx1Yv5QXrEA5Ykg1TxvVQHcVRis--E4TV-9mKGjzZ0VGY-8CHpoKmylp42")
	if err != nil {
		panic(err)
	}

	s, _ := strconv.ParseInt(supply, 10, 64)

	var bool = true
	webhook.SendEmbeds(
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
			}, &api.EmbedField{
				Name:   "Current Hashrate",
				Value:  hashrate,
				Inline: &bool,
			}, &api.EmbedField{
				Name:   "Current Difficulty",
				Value:  difficulty,
				Inline: &bool,
			}, &api.EmbedField{
				Name:   "Current Supply",
				Value:  humanize.Comma(s),
				Inline: &bool,
			}, &api.EmbedField{
				Name:   "Address",
				Value:  addressInfo.Address,
				Inline: &bool,
			}).
			SetFooter("LayersTech Exchange Go Monitor", "https://ravencoinlite.info/wp-content/uploads/2021/09/RVL-transparent-bg.png").
			Build())
}
