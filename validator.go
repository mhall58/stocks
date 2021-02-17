package main

import "github.com/piquette/finance-go"

type Validator struct {
	minPrice float64
	maxPrice float64
	minGrowth float64
}

func (v Validator) isValid(quote *finance.Quote) bool {
	return v.priceIsBetweenMinAndMax(quote) && v.isTrendingUpToday(quote) && v.hasMinimumGrowthPotential(quote)
}

func (v Validator) priceIsBetweenMinAndMax(quote *finance.Quote) bool {
	return quote.RegularMarketPrice > v.minPrice && quote.RegularMarketPrice <= v.maxPrice
}

func (v Validator) isTrendingUpToday(quote *finance.Quote) bool {
	return quote.RegularMarketPreviousClose < quote.RegularMarketPrice
}

func (v Validator) hasMinimumGrowthPotential(quote *finance.Quote) bool {
	return v.GetGrowthPotential(quote) >= v.minGrowth
}

func (v Validator) GetGrowthPotential(quote *finance.Quote)float64 {
	return (quote.FiftyTwoWeekHigh - quote.RegularMarketPrice) / quote.FiftyTwoWeekHigh
}