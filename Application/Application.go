package Application

import (
	"log"
	"os"
	"os/signal"
	"syscall"
)

type Application struct {
}

//STOP Funkcija --->>STOP
func Stop() {
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM, syscall.SIGSTOP, syscall.SIGKILL)
}

//BOT Authentication funkcija --->> INIT
func Authenticate() {
	bot, err := telegram.NewBotAPI(app.Token)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Authorized on account %s", bot.Self.UserName)
}

//FromFile funkcija ?? --->> INIT

//URL ID funkcija ?? --->> INIT

//PriceChecker funkcija --->> GO

//Zinutes issiuntimo i Telegrama funkcija --->>GO
