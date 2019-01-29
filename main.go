package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	crypto "github.com/AntanasMaziliauskas/Crypto_Telegram/crypto"
	telegram "github.com/go-telegram-bot-api/telegram-bot-api"
)

//Program is designed to get the information of the specific crypto currency and compare it against data in file.
//According to the rules, Telegram bot would notify users in the specific channel if the crypto currency price has increased of decreased.
func main() {
	var (
		//	CryptoC     crypto.Crypto
		CryptoRules []crypto.CryptoRule
		err         error
		//r           []strin
	)
	crypto.Init()
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM, syscall.SIGSTOP, syscall.SIGKILL)

	bot, err := telegram.NewBotAPI(crypto.Token)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Authorized on account %s", bot.Self.UserName)

	if CryptoRules, err = crypto.FromFile(crypto.File); err != nil {
		log.Fatal(err)
	}

	r := crypto.URLID(CryptoRules)
	//GoRoutine...Einu per URL sarasa ir tikrinu crypto rates
	/*	if CryptoC, err = crypto.FromURL(r); err != nil {
			log.Fatal(err)
		}

		for _, r := range crypto.CheckAll(CryptoC, CryptoRules) {
			text := fmt.Sprintf("%s is running. %s price has %s!", bot.Self.UserName, CryptoC.Name, crypto.Status(r))
			//siunciam text i kanala
		}
	*/
	//GoRoutine... kai gaunam zinute, siunciam telegrama
	//	msg := telegram.NewMessageToChannel("@CryptTelegram", text)
	//	bot.Send(msg)
	go crypto.PriceChecker(r, CryptoRules)
	go crypto.Sender(bot)

	// os.Signal
	<-stop
}
