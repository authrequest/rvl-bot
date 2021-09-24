package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
)

func getHashrate() string {
	resp, err := http.Get("http://explorer.ravencoinlite.org/api/getnetworkhashps")
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
	resp, err := http.Get("http://explorer.ravencoinlite.org/api/getdifficulty")
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

func getSupply() string {
	resp, err := http.Get("http://explorer.ravencoinlite.org/ext/getmoneysupply")
	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	s, err := strconv.ParseFloat(strings.Trim(string(b), " "), 32)
	if err != nil {
		log.Fatalln(err)
	}

	supply := strconv.FormatFloat(s, 'f', 0, 64)
	fmt.Println(supply)
	return supply
}
