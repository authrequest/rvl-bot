package main

import (
	"time"
)

func main() {
	// var access_key = "b3ce69ea07e684c3"
	// var secret_key = "005c5fd73fd10b09247d8f2f184a190f"
	for range time.Tick(time.Second * 10) {
		hashrate := getHashrate()
		difficulty := getDifficulty()
		supply := getSupply()
		t := fetchApi()
		sendWebhook(t, hashrate, difficulty, supply)
	}
}
