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

const Token = "717631082:AAEaOBNtLs8tJ-DnoWTbCk1Y2i6mawum3jk"
const Channel = "@CryptTelegram"

type message struct {
	Name   string
	Status string
	Price  float64
}

//App structure
type App struct {
	bot   *telegram.BotAPI
	msg   chan string
	token string
	//	rulesList    []types.Rule /
	//	matchedRules []types.Rule
	//	dataFromAPI  []types.LoreData
	Rules   rules.RulesService
	LoreAPI coinlore.CoinloreService
}

//Stop Funkcija --->>STOP
func (a *App) Stop() {

}

//Authenticate function authenticates Telegram bot with the given token and print out authentication message
func (a *App) Authenticate() error {
	var err error

	if a.bot, err = telegram.NewBotAPI(a.token); err != nil {
		return err
	}
	//log.Printf("Authorized on account %s", a.Bot.Self.UserName) // return err - done

	return nil
}

//Init function authenticates Telegram Bot, reads rules from file and prepares a list of URLs for use
func (a *App) Init() error {
	var err error

	a.msg = make(chan string)

	a.token = Token

	a.LoreAPI.Init() //err

	a.Rules.Init()

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

			ids := rules.UniqueRules(rulesList)

			if dataFromAPI, err = a.LoreAPI.FetchAll(ids); err != nil {
				log.Println(err)
			}

			matchedRules := a.Rules.Match(rulesList, dataFromAPI)

			message := a.generateMsg(matchedRules, dataFromAPI)

			a.msg <- message

			updatedRules := a.updateRules(rulesList, matchedRules)

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
	text := telegram.NewMessageToChannel(Channel, message)
	if _, err := a.bot.Send(text); err != nil {
		return err
	}

	return nil
}

//generateMsg function creates a message for sending Telegram
func (a *App) generateMsg(matchedRules []types.Rule, dataFromAPI []types.LoreData) string {
	for _, r := range matchedRules {
		for _, u := range dataFromAPI {
			if r.ID == u.ID {
				text := fmt.Sprintf("%s price has %s! It is %.2f USD now!", u.Name, Status(r), u.Price)
				return text
			}
		}
	}

	return ""
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
