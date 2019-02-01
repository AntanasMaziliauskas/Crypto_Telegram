package crypto

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"

	telegram "github.com/go-telegram-bot-api/telegram-bot-api"
)

type CryptoStruct struct {
	Msg           chan string
	RealURL       string
	URL           string
	File          string
	RulesFromFile []CryptoRule
	URLList       []string
	Bot           *telegram.BotAPI
	DataFromURL   []Crypto
	SatRules      []CryptoRule
}

//Crypto structure used for API JSON
type Crypto struct {
	ID    string  `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price_usd"`
}

//CryptoRule structure used for JSON from file
type CryptoRule struct {
	RuleID   int     `json:ruleid`
	ID       string  `json:"id"`
	Price    float64 `json:"price"`
	Rule     string  `json:"rule"`
	Notified bool    `json:"notified"`
}

const URL = "https://api.coinlore.com/api/ticker/?id=%s"
const File = "CryptoInfo.json"
const Token = "717631082:AAEaOBNtLs8tJ-DnoWTbCk1Y2i6mawum3jk"

//Checking the rule and returning string -->> CRYPTO
func Status(r CryptoRule) string {
	var s string

	if r.Rule == "gt" {
		s = "increased"
	}
	if r.Rule == "lw" {
		s = "decreased"
	}

	return s
}

//WriteToFile function writes the data into the file -->> CRYPTO
func (a *CryptoStruct) WriteToFile() error {

	tofile, _ := json.Marshal(a.RulesFromFile)
	if err := ioutil.WriteFile(a.File, tofile, 0644); err != nil {
		return err
	}

	return nil
}

//Notify function reads through the rules list and changes Notified value if the rule has been possted already.
//Creates new list of rules that passes back. -->> CRYPTO
func (a *CryptoStruct) Notify() {
	var edited []CryptoRule

	if a.SatRules != nil {
		//edited = nil
		for _, re := range a.RulesFromFile {
			for _, pr := range a.SatRules {
				if pr.RuleID == re.RuleID {
					re.Notified = true
				}
			}
			edited = append(edited, re)
		}
		a.RulesFromFile = edited
	}

}

//TODO: Apsirasyti GenerateMsg funkcija
func (a *CryptoStruct) GenerateMsg() {
	for _, r := range a.SatRules {
		for _, u := range a.DataFromURL {
			if r.ID == u.ID {
				text := fmt.Sprintf("%s price has %s! It is %.2f USD now!", u.Name, Status(r), u.Price)
				log.Printf("Sending telegram for rule: %s\n", r)
				a.Msg <- text
			}
		}
	}
}

//CheckAll funcion goes through the array of rules from the file and makes a list of rules that were approved with CheckOne funcion
func (a *CryptoStruct) CheckAll() {
	//	var a []CryptoRule
	//TODO: Perdaryti reikia
	a.SatRules = nil
	for _, u := range a.DataFromURL {
		for _, x := range a.RulesFromFile {
			if CheckOne(u, x) && x.Notified != true {
				a.SatRules = append(a.SatRules, x)
			}
		}
	}
	if len(a.SatRules) == 0 {
		log.Printf("URL: %s no rules matched.\n", a.RealURL)
	}

}

//Checking data received from URL with the data from file according to the rule provided in the file, returning true/false
func CheckOne(c Crypto, r CryptoRule) bool {

	if r.ID == c.ID {
		if r.Rule == "gt" && r.Price < c.Price {
			return true
		}
		if r.Rule == "lw" && r.Price > c.Price {
			return true
		}
	}

	return false
}

//FromFile funcion read a file and unmarshals JSON from the file
func (a *CryptoStruct) ReadRules() error {

	// Init() { ..
	// a.CryptoRules = a.Crypto.ReadRules()
	// }

	var (
		CryptInfo []CryptoRule
		err       error
		jsonFile  []byte
	)

	if jsonFile, err = ioutil.ReadFile(a.File); err != nil {
		return err
	}
	if err = json.Unmarshal(jsonFile, &CryptInfo); err != nil {
		return err
	}
	a.RulesFromFile = CryptInfo
	return err
}

//APRASAS
func (a *CryptoStruct) PrepareURL() {
	//var urls []string

	for _, v := range a.RulesFromFile {
		url := fmt.Sprintf(a.URL, v.ID)
		if !SliceContainsString(url, a.URLList) {
			a.URLList = append(a.URLList, url)
		}

	}
}

// SliceContainsString will return true if needle has been found in haystack.
func SliceContainsString(needle string, haystack []string) bool {
	for _, v := range haystack {
		if v == needle {
			return true
		}
	}
	return false
}

// Function FromURL takes URL address of a JSON, unmarshals JSON and return the data
func (a *CryptoStruct) FromURL() error {
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

	for _, u := range a.URLList {

		spaceClient := http.Client{
			Timeout: time.Second * 10, // 2 secs
		}

		if req, err = http.NewRequest(http.MethodGet, u, nil); err != nil {
			//	fmt.Println("#1")
			return err
		}
		if res, err = spaceClient.Do(req); err != nil {
			//	fmt.Println("#2")
			return err
		}
		if body, err = ioutil.ReadAll(res.Body); err != nil {
			//	fmt.Println("#3")
			return err
		}
		if err = json.Unmarshal(body, &CrypJson); err != nil {
			//	fmt.Println("#4")
			return err
		}
		if len(CrypJson) < 1 {
			//	fmt.Println("#4")
			return errors.New("cryp json empty")
		}
		//fmt.Println("#praejau")
		//fmt.Println(len(CrypJson))
		p, _ := strconv.ParseFloat(CrypJson[0].Price, 64)

		a.DataFromURL = append(a.DataFromURL, Crypto{
			ID:    CrypJson[0].ID,
			Name:  CrypJson[0].Name,
			Price: p,
		})
	}
	return err
}
