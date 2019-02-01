package crypto

import (
	"log"
	"time"

	telegram "github.com/go-telegram-bot-api/telegram-bot-api"
)

//PriceChecker function is a ticker, that every two seconds received data from URL, checks data against rules,
//generates message accordingly, send the message to specific channel, edits rule list and saves it to file.
func (a *CryptoStruct) PriceChecker() {
	var (
		err error
	)

	ticker := time.NewTicker(2 * time.Second)
	for {
		select {
		case <-ticker.C:
			log.Println("tick")

			if err = a.FromURL(); err != nil {
				log.Println(err)
			}
			// AR SITUS VERTA DETI PRIE APP METODO?
			a.CheckAll()

			a.GenerateMsg()

			a.Notify()
			if err = a.WriteToFile(); err != nil {
				log.Println(err)
			}
		}
	}
}

//Sender function recceives message from the channel
func (a *CryptoStruct) Sender() {
	for {
		select {
		case text := <-a.Msg:
			msg := telegram.NewMessageToChannel("@CryptTelegram", text)
			a.Bot.Send(msg)
		}
	}
}
