package application

import (
	"log"
)

//Init function authenticates Telegram Bot, reads rules from file and prepares a list of URLs for use
func (a *App) Init() {

	a.Authenticate()

	if err := a.CS.ReadRules(); err != nil {
		log.Println(err)
	}

	a.CS.PrepareURL()

}
