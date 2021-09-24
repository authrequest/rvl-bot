package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
)

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
