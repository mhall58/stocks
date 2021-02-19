# Stocks Bot

It is said that history repeats itself. This discord bot identifies stock that have been good int he past and are hoping for a rebound, we hope. This is NOT sound financial advice. This is just a way to save time for searching for these type of stocks.

## How it works

1. Gets current stock quotes and selects stocks meeting certain criteria:
    - Stocks current price is above $0.10 amd below $5.0
    - Stocks stock is trending upward
      - higher price than yesturday
      - higher price than the stocks 50 day average price
    - Stock has historical potential of climbing more than 10%
      - For this we look at the difference between the 52 week high point and the current price
    - Stock must also be part of a valid US exchang: NASDAQ, AMEX, and NYSE
1. Recommends a sale price. This is based on the difference between the current price, the 52 week high devided by 2.
1. Sorts the list based on stocks with the highest possible growth
1. Posts the top 10 results to a discord channel
1. The code is setup to be ran on AWS Lambda


## Configuration (.ENV)
see `.env.example` for available config options.

## Future Enhancements

- Automatically disable discord spamming when developing locally
- make thresholds configurable as ENV variables.
- more easily run locally (right now requires you comment some things out)
- A warning message appearing with the tips. 
- protection from IPOs in recent time
- protection from 1 time spikes, we should switch from high to more of a historically sustained high price
- sector anlaysis so we can get tips in different sectors which might let us decide what to buy given the news
- advanced weighting
- simulations: would be good to run our system against the past to see how good these tips actually are. This could allow us to test tweaks.
