package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type TradedCompanies struct {
	Company string `json:"name"`
	Symbol  string `json:"ticker"`
	Exchange  string `json:"exchange"`
}

type Symbols struct {
}

func (s Symbols) getSymbolChunks(chunkSize int) [][]string {
	var symbolChunks [][]string

	symbols := s.getSymbols()

	for i := 0; i < len(symbols); i += chunkSize {
		end := i + chunkSize
		if end > len(symbols) {
			end = len(symbols)
		}

		symbolChunks = append(symbolChunks, symbols[i:end])
	}

	return symbolChunks
}

func (s Symbols) getSymbols() []string {
	var symbols []string

	for _, v := range s.getCompanies() {
		symbols = append(symbols, v.Symbol)
	}

	return symbols
}

func (Symbols) getCompanies() []TradedCompanies {
	var companies []TradedCompanies
	resp,_ := http.Get("https://dumbstockapi.com/stock?exchange=AMEX,NASDAQ,NYSE")

	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	json.Unmarshal(body, &companies)

	return companies
}
