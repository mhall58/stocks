package main

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
	"github.com/piquette/finance-go"
	"github.com/piquette/finance-go/quote"
	"log"
	"os"
	"sort"
)

type Result struct {
	quote              *finance.Quote
	hypotheticalGrowth float64
}

func main() {

	//todo make these configurable
	maxPrice := 5.0
	minPrice := .01
	minGrowth := .10
	chunkSize := 500

	symbolChunks := Symbols{}.getSymbolChunks(chunkSize)

	stockValidator := Validator{
		minPrice:  minPrice,
		maxPrice:  maxPrice,
		minGrowth: minGrowth,
	}

	var results []Result

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

			if stockValidator.isValid(q) {
				results = append(results, Result{
					quote:              q,
					hypotheticalGrowth: stockValidator.GetGrowthPotential(q),
				})
			}

		}

	}

	// sort descending
	sort.Slice(results, func(i, j int) bool {
		return results[i].hypotheticalGrowth > results[j].hypotheticalGrowth
	})

	report := ""

	for i, v := range results {
		if i == 10 {
			break
		}

		report = report + fmt.Sprintf(
			"SYMBOL: %s, Price: %f, 52W-High: %f, Hypthetical Growth %f \n",
			v.quote.Symbol,
			v.quote.RegularMarketPrice,
			v.quote.FiftyTwoWeekHigh,
			v.hypotheticalGrowth,
		)

	}

	fmt.Println(report)

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	if os.Getenv("DISCORD_BOT_TOKEN") != "" {
		dg, _ := discordgo.New("Bot " + os.Getenv("DISCORD_BOT_TOKEN"))
		dg.Open()
		dg.ChannelMessageSend(os.Getenv("DISCORD_CHANNEL_ID"), report)
		dg.Close()
	}

}
