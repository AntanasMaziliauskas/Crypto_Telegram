package main

import (
	"flag"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/AntanasMaziliauskas/Crypto_Telegram/coinlore"
	"github.com/AntanasMaziliauskas/Crypto_Telegram/rules"

	"github.com/AntanasMaziliauskas/Crypto_Telegram/application"
)

//const Token = "717631082:AAEaOBNtLs8tJ-DnoWTbCk1Y2i6mawum3jk"

//Program is designed to get the information of the specific crypto currency and compare it against data in file.
//According to the rules, Telegram bot would notify users in the specific channel if the crypto currency price has increased of decreased.
func main() {

	path := flag.String("path", "rules.xml", "a string")
	//typ := flag.Bool("type", true, "a bool")
	token := flag.String("token", "717631082:AAEaOBNtLs8tJ-DnoWTbCk1Y2i6mawum3jk", "a string")
	channel := flag.String("channel", "@CryptTelegram", "a string")
	flag.Parse()

	fileHandler := &rules.RulesFromXML{
		Path: *path,
	}
	/*if !*typ {
		fileHandler = &rules.RulesFromXML{
			Path: *path,
		}
	}*/

	app := application.App{
		Token:   *token,
		Channel: *channel,
		Rules:   fileHandler,
		LoreAPI: &coinlore.CoinloreAPI{},
	}

	if err := app.Init(); err != nil {
		log.Fatal(err)
	}

	app.Go()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM, syscall.SIGSTOP, syscall.SIGKILL)
	<-stop

}
