package Crypto

import (
	"fmt"
	"log"
	"time"

	telegram "github.com/go-telegram-bot-api/telegram-bot-api"
)

var Msg chan string

func Init() {
	Msg = make(chan string)
}

func PriceChecker(urls []string, rules []CryptoRule) {
	var (
		CryptoC Crypto
		err     error
	)

	ticker := time.NewTicker(2 * time.Second)
	for {
		select {
		case <-ticker.C:
			log.Println("tick")
			for _, url := range urls {
				//GoRoutine...Einu per URL sarasa ir tikrinu crypto rates
				if CryptoC, err = FromURL(url); err != nil {
					log.Println(err)
					continue
				}
				//	log.Printf("cryptoc :%#v, rules: %#v\n", CryptoC, rules)
				kint := CheckAll(CryptoC, rules)
				if len(kint) == 0 {
					//		log.Printf("URL: %s no rules matched.\n", url)
					continue
				}
				for _, r := range kint {
					text := fmt.Sprintf("%s price has %s! It is %.2f USD now!", CryptoC.Name, Status(r), CryptoC.Price)
					log.Printf("Sending text to %s\n", url)
					Msg <- text
				}
				rules = Notify(rules, kint)
				WriteToFile(rules)
			}
		}
	}
}

func Sender(bot *telegram.BotAPI) {
	for {
		select {
		case text := <-Msg:
			msg := telegram.NewMessageToChannel("@CryptTelegram", text)
			bot.Send(msg)
		}
	}
}
