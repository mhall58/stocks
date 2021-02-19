package main

import "github.com/piquette/finance-go"

type Validator struct {
	minPrice  float64
	maxPrice  float64
	minGrowth float64
}

func (v Validator) isValid(quote *finance.Quote) bool {
	return v.hasValidExchange(quote) && v.priceIsBetweenMinAndMax(quote) && v.isTrendingUp(quote) && v.hasMinimumGrowthPotential(quote)
}

func (v Validator) priceIsBetweenMinAndMax(quote *finance.Quote) bool {
	return quote.RegularMarketPrice > v.minPrice && quote.RegularMarketPrice <= v.maxPrice
}

func (v Validator) isTrendingUp(quote *finance.Quote) bool {
	return quote.RegularMarketPreviousClose < quote.RegularMarketPrice && quote.RegularMarketPrice > quote.FiftyDayAverage
}

func (v Validator) hasMinimumGrowthPotential(quote *finance.Quote) bool {
	return Result{}.getPotentialGrowth(quote) >= v.minGrowth
}

func (v Validator) hasValidExchange(quote *finance.Quote) bool {
	return quote.FullExchangeName != "Other OTC"
}


