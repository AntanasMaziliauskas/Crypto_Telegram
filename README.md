[![Build Status](https://travis-ci.com/AntanasMaziliauskas/Crypto_Telegram.svg?branch=master)](https://travis-ci.com/AntanasMaziliauskas/Crypto_Telegram)
[![Go Report Card](https://goreportcard.com/badge/github.com/AntanasMaziliauskas/Crypto_Telegram)](https://goreportcard.com/report/github.com/AntanasMaziliauskas/Crypto_Telegram)
# Crypto_Telegram

Crypto Telegram is designed to get the information of the specific crypto currency and compare it against the rules written in file. According to the rules, Telegram bot would notify users in the specific channel if the crypto currency price has increased of decreased.

## Crypto Currency Rates

Rates are received from [CoinLore][a]. It uses URL address to get information for specific coin.
Example: https://api.coinlore.com/api/ticker/?id=90
This is the example of JSON you get:

[
  {
    "id": "90",
    "symbol": "BTC",
    "name": "Bitcoin",
    "nameid": "bitcoin",
    "rank": 1,
    "price_usd": "6465.26",
    "percent_change_24h": "-1.27",
    "percent_change_1h": "0.19",
    "percent_change_7d": "-0.93",
    "market_cap_usd": "111737012373.28",
    "volume24": "3982512765.23",
    "volume24_native": "615986.77",
    "csupply": "17282687.00",
    "price_btc": "1.00",
    "tsupply": "17282687",
    "msupply": "21000000"
  }
]         

## Rules

info about file types and rules template

## Telegram Bot

info about what does telegram bot do

[a]: <https://www.coinlore.com/>
