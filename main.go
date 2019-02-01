package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/AntanasMaziliauskas/Crypto_Telegram/application"
	"github.com/AntanasMaziliauskas/Crypto_Telegram/crypto"
)

//Program is designed to get the information of the specific crypto currency and compare it against data in file.
//According to the rules, Telegram bot would notify users in the specific channel if the crypto currency price has increased of decreased.
func main() {

	app := application.App{
		Token: crypto.Token,
		CS: crypto.CryptoStruct{
			File: crypto.File,
			URL:  crypto.URL,
			Msg:  make(chan string),
		},
	}

	app.Init()

	app.Go()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM, syscall.SIGSTOP, syscall.SIGKILL)
	<-stop
}
