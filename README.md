[![Build Status](https://travis-ci.com/AntanasMaziliauskas/Crypto_Telegram.svg?branch=master)](https://travis-ci.com/AntanasMaziliauskas/Crypto_Telegram)
[![Go Report Card](https://goreportcard.com/badge/github.com/AntanasMaziliauskas/Crypto_Telegram)](https://goreportcard.com/report/github.com/AntanasMaziliauskas/Crypto_Telegram)
# Crypto_Telegram

Crypto Telegram is designed to get the information of the specific crypto currency and compare it against the rules written in file. According to the rules, Telegram bot would notify users in the specific channel if the crypto currency price has increased of decreased.

## Crypto Currency Rates

- Rates are received from [CoinLore][a]. It uses URL address to get information for specific coin.
- URL example: https://api.coinlore.com/api/ticker/?id=90
- This is the example of JSON you get:
```sh
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
```
## Rules

You can provide two types of files for the list of rules:
 - JSON
 - XML
JSON example:
```
[
{"RuleID":0,"id":"90","price":3470.98,"rule":"lt","notified":false},
{"RuleID":1,"id":"90","price":3070.98,"rule":"gt","notified":false},
{"RuleID":2,"id":"91","price":100000.223,"rule":"lt","notified":false},
{"RuleID":3,"id":"92","price":100000.223,"rule":"lt","notified":false}
]
```
 XML example:
 ```
 <rules>
    <rule>
        <ruleid>0</ruleid>
        <id>90</id>
        <price>3470.98</price>
        <rule>lt</rule>
        <notified>false</notified>
    </rule>
    <rule>
        <ruleid>1</ruleid>
        <id>90</id>
        <price>3470.98</price>
        <rule>gt</rule>
        <notified>false</notified>
    </rule>
    <rule>
        <ruleid>2</ruleid>
        <id>91</id>
        <price>3470.98</price>
        <rule>lt</rule>
        <notified>false</notified>
    </rule>
    <rule>
        <ruleid>3</ruleid>
        <id>92</id>
        <price>3470.98</price>
        <rule>lt</rule>
        <notified>false</notified>
    </rule>
</rules>
 ```
When file is read, rules are being checked against the information received from the CoinLore. Every rule that is satiesfied is being placed into a new list of rules.

## Telegram Bot

Whenever there is a satisfied rule from the list Telegram bot provides information about the coin price and lets us know if it increased or decreased according to that rule.

[a]: <https://www.coinlore.com/>
