package main

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
	"github.com/piquette/finance-go/quote"
	"log"
	"os"
)



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

	var results Results

	for _, chunk := range symbolChunks {
		quotes := quote.List(chunk)

		for i := 0; i <= quotes.Count(); i++ {
			quotes.Next()

			if quotes.Err() != nil {
				break
			}

			quote := quotes.Quote()

			if quote == nil {
				continue
			}

			if stockValidator.isValid(quote) {
				results = append(results, Result{}.New(quote))
			}

		}

	}

	report := results.sortByPGrowth().asString(10)
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
