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

//App struct
type App struct {
	bot     *telegram.BotAPI
	msg     chan string
	Token   string
	Channel string
	Rules   rules.RulesService
	LoreAPI coinlore.CoinloreService
}

//Stop Funkcija --->>STOP TOTO: kaip turi atrodyti?
func (a *App) Stop() {

}

//Authenticate function authenticates Telegram bot with the given token
func (a *App) Authenticate() error {
	var err error

	if a.bot, err = telegram.NewBotAPI(a.Token); err != nil {
		return err
	}

	return nil
}

//Init function launches LoreAPI.Init and Rules.Init, also authenticates Telegram Bot
func (a *App) Init() error {
	var err error

	a.msg = make(chan string)

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

// Go function starts two Go-Routines: PriceChekeTicker and TelegramBotTicker
func (a *App) Go() {

	go a.PriceCheckTicker()

	go a.TelegramBotTicker()
}

//PriceCheckTicker function every two seconds launches these functions:
//ReadRules, Unique Rules, FetchAll, Match, generateMsg, updateRules and SaveRules.
func (a *App) PriceCheckTicker() {
	var (
		err         error
		rulesList   []types.Rule
		dataFromAPI []types.LoreData
	)

	ticker := time.NewTicker(1 * time.Minute)
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

			for _, m := range message {
				a.msg <- m
			}

			updatedRules := a.updateRules(rulesList, matchedRules)

			if err = a.Rules.SaveRules(updatedRules); err != nil {
				log.Println(err)
			}
		}
	}
}

//TelegramBotTicker function receives message from the channel
//Uses sendToTelegram function to send a message.
func (a *App) TelegramBotTicker() {
	for {
		select {
		case text := <-a.msg:
			a.sendToTelegram(text)
		}
	}
}

//sendToTelegram function sends a message to the channel
func (a *App) sendToTelegram(message string) error {
	text := telegram.NewMessageToChannel(a.Channel, message)
	if _, err := a.bot.Send(text); err != nil {
		return err
	}

	fmt.Printf("Message that was sent to channel %s: %s\n", a.Channel, message)

	return nil
}

//generateMsg function creates a message for sending to Telegram
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

//updateRules function updates existing rule list. It changes the notified field for the rules that matched.
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

//Status function is checking if rule is GT or LT and returning 'increased/decreased'
func Status(r types.Rule) string {
	var s string

	if r.Rule == "gt" {
		s = "increased"
	}
	if r.Rule == "lt" {
		s = "decreased"
	}

	return s
}
