package Crypto

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
)

type Application struct {
	Crypto     Crypto
	CryptoRule []CryptoRule
	URL        string
	File       string
	Token      string
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
func WriteToFile(edited []CryptoRule) error {

	tofile, _ := json.Marshal(edited)
	if err := ioutil.WriteFile(File, tofile, 0644); err != nil {
		return err
	}

	return nil
}

//Notify function reads through the rules list and changes Notified value if the rule has been possted already.
//Creates new list of rules that passes back. -->> CRYPTO
func Notify(real []CryptoRule, printed []CryptoRule) []CryptoRule {
	var edited []CryptoRule

	if printed == nil {
		return real
	}
	for _, pr := range printed {
		for _, re := range real {
			if pr.RuleID == re.RuleID {
				re.Notified = true
			}
			edited = append(edited, re)
		}
	}

	return edited
}

//CheckAll funcion goes through the array of rules from the file and makes a list of rules that were approved with CheckOne funcion
func CheckAll(c Crypto, r []CryptoRule) []CryptoRule {
	var a []CryptoRule

	for _, x := range r {
		if CheckOne(c, x) && x.Notified != true {
			a = append(a, x)
		}
	}

	return a
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
func (a *Application) FromFile() ([]CryptoRule, error) {

	// Init() { ..
	// a.CryptoRules = a.Crypto.ReadRules()
	// }

	var (
		CryptInfo []CryptoRule
		err       error
		jsonFile  []byte
	)

	if jsonFile, err = ioutil.ReadFile(a.File); err != nil {
		return CryptInfo, err
	}
	if err = json.Unmarshal(jsonFile, &CryptInfo); err != nil {
		return CryptInfo, err
	}

	return CryptInfo, err
}

//APRASAS
func (a *Application) URLID() []string {
	var urls []string

	for _, v := range a.CryptoRule {
		url := fmt.Sprintf(URL, v.ID)
		if !SliceContainsString(url, urls) {
			urls = append(urls, url)
		}

	}
	return urls
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
