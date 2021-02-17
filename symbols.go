package main

import (
	"encoding/json"
	"io/ioutil"
)

type TradedCompanies struct {
	Company string `json:"Company Name"`
	Symbol  string `json:"ACT Symbol"`
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
	jsonFile, _ := ioutil.ReadFile("json/nyse.json")

	err := json.Unmarshal(jsonFile, &companies)
	if err != nil {
		panic(err)
	}

	return companies
}
