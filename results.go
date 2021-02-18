package main

import (
	"fmt"
	"github.com/piquette/finance-go"
	"sort"
)

type Result struct {
	quote     *finance.Quote
	pGrowth   float64
	sellPrice float64
}

func (r Result) New(quote *finance.Quote) Result {
	x := Result{
		quote:     quote,
		pGrowth:   r.getPotentialGrowth(quote),
		sellPrice: r.getSellPrice(quote),
	}

	return x
}

func (r Result) getPotentialGrowth(quote *finance.Quote) float64 {
	return (quote.FiftyTwoWeekHigh - quote.RegularMarketPrice) / quote.FiftyTwoWeekHigh
}

func (r Result) getSellPrice(quote *finance.Quote) float64 {
	return (quote.FiftyTwoWeekHigh + quote.RegularMarketPrice) / 2
}

func (r Result) asString() string {
	return fmt.Sprintf(
		"Buy %s at $%.2f per share and limit sell it for $%.2f \n",
		r.quote.Symbol,
		r.quote.RegularMarketPrice,
		r.sellPrice,
	)
}

type Results []Result

func (results Results) sortByPGrowth() Results {
	sort.Slice(results, func(i, j int) bool {
		return results[i].pGrowth > results[j].pGrowth
	})
	return results
}

func (results Results) asString(cnt int) string {
	output := ""
	for i, result := range results {
		output = output + result.asString()
		if i+1 == cnt {
			break
		}
	}
	return output
}
