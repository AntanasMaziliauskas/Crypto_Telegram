package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	crypto "github.com/AntanasMaziliauskas/Crypto_Telegram/Crypto"
	telegram "github.com/go-telegram-bot-api/telegram-bot-api"
)

//Program is designed to get the information of the specific crypto currency and compare it against data in file.
//According to the rules, Telegram bot would notify users in the specific channel if the crypto currency price has increased of decreased.
func main() {
	var (
		CryptoRules []crypto.CryptoRule
		err         error
	)

	app := crypto.Application{
		Token:      "717631082:AAEaOBNtLs8tJ-DnoWTbCk1Y2i6mawum3jk",
		URL:        "https://api.coinlore.com/api/ticker/?id=%s",
		File:       crypto.File,
		CryptoRule: CryptoRules,
	}

	// if err := app.Init(); err != nil { log.Fatal(err) }
	// crypto.Init()

	// if err := app.Start(); err != nil ...

	// <-stop

	// app.Stop()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM, syscall.SIGSTOP, syscall.SIGKILL)

	//Perkelti i atskira funkcija?
	bot, err := telegram.NewBotAPI(app.Token)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Authorized on account %s", bot.Self.UserName)

	// app.Init()
	app.CryptoRule, err = app.FromFile()

	/*	if CryptoRules, err = crypto.FromFile(crypto.File); err != nil {
		log.Fatal(err)
	}*/
	r := app.URLID()

	go crypto.PriceChecker(r, app.CryptoRule)
	go crypto.Sender(bot)

	<-stop

	/*

		output :=

		app := telegramticker.Application{
			Token: "asdasdas",
			Source: URL,
			Msg: make(chan string),
			Output:	crypto.Output{
				File: "notify.json"
			}
		}

		// init: a.Rules = a.Output.ReadRules()

		a.rules := a.Output.ReadRules()


		a.Output.SaveRules()


	*/
}
