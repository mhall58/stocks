package main

import (
	"encoding/json"
	"fmt"
	"github.com/piquette/finance-go"
	"github.com/piquette/finance-go/quote"
	"io/ioutil"
	"sort"
)

type Stock struct {
	Company string `json:"Company Name"`
	Symbol  string `json:"ACT Symbol"`
}

type Result struct {
	quote              *finance.Quote
	hypotheticalGrowth float64
}

func main() {

	jsonFile, _ := ioutil.ReadFile("json/nyse.json")

	//todo make these configurable
	maxPrice := 5.0
	minPrice := .01
	minGrowth := .10
	chunkSize := 500

	var stocks []Stock

	err := json.Unmarshal(jsonFile, &stocks)
	if err != nil {
		panic(err)
	}

	var symbols []string

	for _, v := range stocks {
		symbols = append(symbols, v.Symbol)
	}

	var symbolChunks [][]string

	for i := 0; i < len(symbols); i += chunkSize {

		end := i + chunkSize

		if end > len(symbols) {
			end = len(symbols)
		}

		symbolChunks = append(symbolChunks, symbols[i:end])
	}

	var results []Result
	resultCnt := 0

	for _, chunk := range symbolChunks {
		quotes := quote.List(chunk)

		for i := 0; i <= quotes.Count(); i++ {
				quotes.Next()

			if quotes.Err() != nil {
				break
			}

			q := quotes.Quote()

			if q == nil {
				continue
			}

			//if !q.IsTradeable {
			//	continue
			//}

			//stock should be over min
			if q.RegularMarketPrice < minPrice {
				continue
			}

			//stock should be under max
			if q.RegularMarketPrice > maxPrice {
				continue
			}

			//stock should be trending up.
			if q.RegularMarketPreviousClose > q.RegularMarketPrice {
				continue
			}

			//stock should be trending up.
			if q.RegularMarketPreviousClose > q.RegularMarketPrice {
				continue
			}

			hypotheticalGrowth := (q.FiftyTwoWeekHigh - q.RegularMarketPrice) / q.FiftyTwoWeekHigh

			//stock should have historical potential
			if hypotheticalGrowth < minGrowth {
				continue
			}

			results = append(results, Result{
				quote:              q,
				hypotheticalGrowth: hypotheticalGrowth,
			})

			resultCnt++
		}

	}

	fmt.Println("\r\n Matt's top ten stock picks!")

	// sort descending
	sort.Slice(results, func(i, j int) bool {
		return results[i].hypotheticalGrowth > results[j].hypotheticalGrowth
	})

	for i, v := range results {
		if i == 10 {
			break
		}

		fmt.Println(
			fmt.Sprintf(
				"SYMBOL: %s, Price: %f, 52W-High: %f, Hypthetical Growth %f",
				v.quote.Symbol,
				v.quote.RegularMarketPrice,
				v.quote.FiftyTwoWeekHigh,
				v.hypotheticalGrowth,
			))
	}

}
