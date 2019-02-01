package application

import (
	"log"

	"github.com/AntanasMaziliauskas/Crypto_Telegram/crypto"
	telegram "github.com/go-telegram-bot-api/telegram-bot-api"
)

//App structure
type App struct {
	Token      string
	CryptoRule []crypto.CryptoRule
	err        error
	CS         crypto.CryptoStruct
}

//Stop Funkcija --->>STOP
func (a *App) Stop() {

}

//Authenticate function authenticates Telegram bot with the given token and print out authentication message
func (a *App) Authenticate() {
	a.CS.Bot, a.err = telegram.NewBotAPI(a.Token)
	if a.err != nil {
		log.Fatal(a.err)
	}
	log.Printf("Authorized on account %s", a.CS.Bot.Self.UserName)
}
