package crypto

import (
	"fmt"
	"log"
	"time"

	//"github.com/AntanasMaziliauskas/Crypto_Telegram/crypto"

	telegram "github.com/go-telegram-bot-api/telegram-bot-api"
)

var Msg chan string

func Init() {
	Msg = make(chan string)
}

/*func botas() {
	bot, err := telegram.NewBotAPI(Token)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Authorized on account %s", bot.Self.UserName)
}*/

func PriceChecker(urls []string, rules []CryptoRule) {
	var (
		CryptoC Crypto
		err     error
		//r           string
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

				log.Printf("cryptoc :%#v, rules: %#v\n", CryptoC, rules)
				kint := CheckAll(CryptoC, rules)
				if len(kint) == 0 {
					log.Printf("URL: %s no rules matched.\n", url)
					continue
				}
				for _, r := range kint {
					text := fmt.Sprintf("%s price has %s! It is %v USD now!", CryptoC.Name, Status(r), CryptoC.Price)
					log.Printf("Sending text to %s\n", url)
					Msg <- text
					//siunciam text i kanala
				}
				rules = Notify(rules, kint)
			}
		}
	}
}

func Sender(bot *telegram.BotAPI) {
	for {
		select {
		case text := <-Msg:
			// send to telegram
			msg := telegram.NewMessageToChannel("@CryptTelegram", text)
			bot.Send(msg)
		}
	}
}
