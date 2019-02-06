package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/AntanasMaziliauskas/Crypto_Telegram/coinlore"
	"github.com/AntanasMaziliauskas/Crypto_Telegram/rules"

	"github.com/AntanasMaziliauskas/Crypto_Telegram/application"
)

//Program is designed to get the information of the specific crypto currency and compare it against data in file.
//According to the rules, Telegram bot would notify users in the specific channel if the crypto currency price has increased of decreased.
func main() {
	var fileHandler rules.RulesService

	path := flag.String("path", "rules", "a path with the file name for reading rules")
	typ := flag.Bool("type", false, "File type. For JSON type set 'false', for XML type set 'true'; default 'false'")
	token := flag.String("token", "717631082:AAEaOBNtLs8tJ-DnoWTbCk1Y2i6mawum3jk", "a token that is used to authenticate Telegram bot")
	channel := flag.String("channel", "@CryptTelegram", "Telegram channel name that's being used for sending messages")
	flag.Parse()

	fileHandler = &rules.RulesFromJSON{
		Path: fmt.Sprintf("%s.json", *path),
	}
	if *typ {
		fileHandler = &rules.RulesFromXML{
			Path: fmt.Sprintf("%s.xml", *path),
		}
	}

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
