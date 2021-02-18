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
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	symbolChunks := Symbols{}.getSymbolChunks(500)

	stockValidator := Validator{
		minPrice:  .01, // 10 cents
		maxPrice:  5.0, // $5
		minGrowth: .10, // 10 percent
	}

	var results Results

	for _, chunk := range symbolChunks {
		quotes := quote.List(chunk)
		for quotes.Next() {
			q := quotes.Quote()
			if stockValidator.isValid(q) {
				results = append(results, Result{}.New(q))
			}
		}
	}

	report := results.sortByPGrowth().asString(10)
	fmt.Println(report)

	if os.Getenv("DISCORD_BOT_TOKEN") != "" {
		dg, _ := discordgo.New("Bot " + os.Getenv("DISCORD_BOT_TOKEN"))
		dg.Open()
		dg.ChannelMessageSend(os.Getenv("DISCORD_CHANNEL_ID"), report)
		dg.Close()
	}

}
