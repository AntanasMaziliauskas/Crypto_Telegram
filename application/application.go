package application

import (
	"fmt"
	"log"
	"time"

	"github.com/AntanasMaziliauskas/Crypto_Telegram/coinlore"
	"github.com/AntanasMaziliauskas/Crypto_Telegram/types"

	"github.com/AntanasMaziliauskas/Crypto_Telegram/rules"
	telegram "github.com/go-telegram-bot-api/telegram-bot-api"
)

//const Token = "717631082:AAEaOBNtLs8tJ-DnoWTbCk1Y2i6mawum3jk"
//const Channel = "@CryptTelegram"

/*
type message struct {
	Name   string
	Status string
	Price  float64
}*/

//App structure
type App struct {
	bot     *telegram.BotAPI
	msg     chan string
	Token   string
	Channel string
	Rules   rules.RulesService
	LoreAPI coinlore.CoinloreService
}

//Stop Funkcija --->>STOP
func (a *App) Stop() {

}

//Authenticate function authenticates Telegram bot with the given token and print out authentication message
func (a *App) Authenticate() error {
	var err error

	if a.bot, err = telegram.NewBotAPI(a.Token); err != nil {
		return err
	}
	//log.Printf("Authorized on account %s", a.Bot.Self.UserName) // return err - done

	return nil
}

//Init function authenticates Telegram Bot, reads rules from file and prepares a list of URLs for use
func (a *App) Init() error {
	var err error

	a.msg = make(chan string)

	//a.token = Token

	if err = a.LoreAPI.Init(); err != nil {
		return err
	}

	if err = a.Rules.Init(); err != nil {
		return err
	}

	if err = a.Authenticate(); err != nil {
		return err
	}

	return nil
}

// Go function starts two Go-Routines: PriceCheker and Sender
func (a *App) Go() {

	go a.PriceCheckTicker()

	go a.TelegramBotTicker()
}

//PriceChecker function is a ticker, that every two seconds received data from URL, checks data against rules,
//generates message accordingly, send the message to specific channel, edits rule list and saves it to file.
func (a *App) PriceCheckTicker() {
	var (
		err         error
		rulesList   []types.Rule
		dataFromAPI []types.LoreData
	)

	ticker := time.NewTicker(2 * time.Second)
	for {
		select {
		case <-ticker.C:
			log.Println("tick")

			if rulesList, err = a.Rules.ReadRules(); err != nil {
				log.Println(err)
			}
			fmt.Println("ReadRules: ", rulesList)

			ids := rules.UniqueRules(rulesList)
			//fmt.Println("UniqueRules: ", ids)

			if dataFromAPI, err = a.LoreAPI.FetchAll(ids); err != nil {
				log.Println(err)
			}

			//fmt.Println("FetchAll: ", dataFromAPI)

			matchedRules := a.Rules.Match(rulesList, dataFromAPI)
			//fmt.Println("Match: ", matchedRules)

			message := a.generateMsg(matchedRules, dataFromAPI)
			//fmt.Println("Message: ", message)

			//TODO: For
			for _, m := range message {
				a.msg <- m
			}

			updatedRules := a.updateRules(rulesList, matchedRules)
			fmt.Println("updatedRules: ", updatedRules)
			if err = a.Rules.SaveRules(updatedRules); err != nil {
				log.Println(err)
			}
		}
	}
}

//Sender function recceives message from the channel
func (a *App) TelegramBotTicker() {
	for {
		select {
		case text := <-a.msg:

			// msg := fmt.Sprintf("%s price has %s! It is %.2f USD now!", u.Name, Status(r), u.Price)
			a.sendToTelegram(text)

		}
	}
}

func (a *App) sendToTelegram(message string) error {
	text := telegram.NewMessageToChannel(a.Channel, message)
	if _, err := a.bot.Send(text); err != nil {
		return err
	}

	return nil
}

//generateMsg function creates a message for sending Telegram
func (a *App) generateMsg(matchedRules []types.Rule, dataFromAPI []types.LoreData) []string {
	var text []string

	for _, r := range matchedRules {
		for _, u := range dataFromAPI {
			if r.ID == u.ID {
				text = append(text, fmt.Sprintf("%s price has %s! It is %.2f USD now!", u.Name, Status(r), u.Price))
			}
		}
	}

	return text
}

//updateRules function updates existing rule list. It flags the rules that matched.
func (a *App) updateRules(rules []types.Rule, matched []types.Rule) []types.Rule {
	var updated []types.Rule

	if matched != nil {
		for _, re := range rules {
			for _, pr := range matched {
				if pr.RuleID == re.RuleID {
					re.Notified = true
				}
			}
			updated = append(updated, re)
		}
	} else {
		return rules
	}

	return updated
}

//Checking the rule and returning string -->> CRYPTO
func Status(r types.Rule) string {
	var s string

	if r.Rule == "gt" {
		s = "increased"
	}
	if r.Rule == "lw" {
		s = "decreased"
	}

	return s
}
