package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strings"
)

const PolygonPath = "https://api.polygon.io"

type Stock struct {
	Ticker string `json:"ticker"`
	Name   string `json:"name"`
	Price  float64
}

type Values struct {
	Open float64 `json:"open"`
}

func SearchTicker(ticker string, apiKey string) []Stock {
	resp, err := http.Get(PolygonPath + "/v3/reference/tickers?" +
		apiKey + "&ticker=" + strings.ToUpper(ticker))

	if err != nil {
		log.Fatal(err)
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("Errore nella richiesta", err)
	}
	data := struct {
		Results []Stock `json:"results"`
	}{}

	json.Unmarshal(body, &data)
	return data.Results
}

func GetDailyValues(ticker string, apiKey string) Values {
	resp, err := http.Get(PolygonPath + "/v1/open-close/" + strings.ToUpper(ticker) + "/2023-10-10/?" + apiKey)
	if err != nil {
		log.Fatal(err)
	}
	body, err := io.ReadAll(resp.Body)

	if err != nil {
		log.Fatal(err)
	}

	data := Values{}
	json.Unmarshal(body, &data)
	return data
}
