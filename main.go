package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"

	telegram "github.com/go-telegram-bot-api/telegram-bot-api"
)

//Crypto structure used for API JSON
type Crypto struct {
	ID    string  `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price_usd"`
}

//CryptoRule structure used for JSON from file
type CryptoRule struct {
	ID    string  `json:"id"`
	Price float64 `json:"price"`
	Rule  string  `json:"rule"`
}

const URL = "https://api.coinlore.com/api/ticker/?id=90"
const File = "CryptoInfo.json"
const token = "717631082:AAEaOBNtLs8tJ-DnoWTbCk1Y2i6mawum3jk"

//APSIRASYTI
func main() {
	var (
		CryptoC     Crypto
		CryptoRules []CryptoRule
		err         error
	)

	if CryptoC, err = FromURL(URL); err != nil {
		log.Fatal(err)
	}
	if CryptoRules, err = FromFile(File); err != nil {
		log.Fatal(err)
	}

	bot, err := telegram.NewBotAPI(token)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Authorized on account %s", bot.Self.UserName)

	for _, r := range CheckAll(CryptoC, CryptoRules) {
		s := status(r)
		text := fmt.Sprintf("%s is running. %s price has %s!", bot.Self.UserName, CryptoC.Name, s)
		msg := telegram.NewMessageToChannel("@CryptTelegram", text)
		bot.Send(msg)
	}
}

//Checking the rule and returning string
func status(r CryptoRule) string {
	var s string

	if r.Rule == "gt" {
		s = "increased"
	}
	if r.Rule == "lw" {
		s = "decreased"
	}

	return s
}

//CheckAll funcion goes through the array of rules from the file and makes a list of rules that were approved with CheckOne funcion
func CheckAll(c Crypto, r []CryptoRule) []CryptoRule {
	var a []CryptoRule

	for _, x := range r {
		if CheckOne(c, x) == true {
			a = append(a, x)
		}
	}

	return a
}

//Checking data received from URL with the data from file according to the rule provided in the file, returning true/false
func CheckOne(c Crypto, r CryptoRule) bool {

	if r.ID == c.ID {
		if r.Rule == "gt" && r.Price > c.Price {
			return true
		}
		if r.Rule == "lw" && r.Price < c.Price {
			return true
		}
	}

	return false
}

//FromFile funcion read a file and unmarshals JSON from the file
func FromFile(FileName string) ([]CryptoRule, error) {
	var (
		CryptInfo []CryptoRule
		err       error
		jsonFile  []byte
	)

	if jsonFile, err = ioutil.ReadFile(FileName); err != nil {
		return CryptInfo, err
	}
	if err = json.Unmarshal(jsonFile, &CryptInfo); err != nil {
		return CryptInfo, err
	}

	return CryptInfo, err
}

// Function FromURL takes URL address of a JSON, unmarshals JSON and return the data
func FromURL(URLName string) (Crypto, error) {
	type CryptoJSON struct {
		ID    string `json:"id"`
		Name  string `json:"name"`
		Price string `json:"price_usd"`
	}

	var (
		err      error
		req      *http.Request
		res      *http.Response
		body     []byte
		CrypJson []CryptoJSON
	)

	spaceClient := http.Client{
		Timeout: time.Second * 10, // 2 secs
	}

	if req, err = http.NewRequest(http.MethodGet, URLName, nil); err != nil {
		return Crypto{}, err
	}
	if res, err = spaceClient.Do(req); err != nil {
		return Crypto{}, err
	}
	if body, err = ioutil.ReadAll(res.Body); err != nil {
		return Crypto{}, err
	}
	if err = json.Unmarshal(body, &CrypJson); err != nil {
		return Crypto{}, err
	}
	if len(CrypJson) != 1 {
		return Crypto{}, err
	}

	p, _ := strconv.ParseFloat(CrypJson[0].Price, 64)

	return Crypto{
		ID:    CrypJson[0].ID,
		Name:  CrypJson[0].Name,
		Price: p,
	}, nil
}
