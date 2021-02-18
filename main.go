package main

import (
	"context"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
	"github.com/piquette/finance-go/quote"
	"log"
	"os"
)

type Request struct {
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	//tips, _ := getStockTips()
	//fmt.Println(tips)
	lambda.Start(HandleRequest)
}

func HandleRequest(ctx context.Context, request Request) (string, error) {
	tips, err := getStockTips()
	postToDiscordChanel(tips)
	return tips, err
}

func getStockTips() (string, error) {

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

	return report, nil
}

func postToDiscordChanel(message string) {
	if os.Getenv("DISCORD_BOT_TOKEN") != "" {
		dg, _ := discordgo.New("Bot " + os.Getenv("DISCORD_BOT_TOKEN"))
		dg.Open()
		dg.ChannelMessageSend(os.Getenv("DISCORD_CHANNEL_ID"), message)
		dg.Close()
	}
}
