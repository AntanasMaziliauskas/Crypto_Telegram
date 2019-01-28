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

/*
[
	{"id":"80", "price":90000.233, "rule":"gt"},
]
*/

//Crypto structure used for JSON
type Crypto struct {
	ID    string  `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price_usd"`
}

//APSIRASYTI
type CryptoRules struct {
	ID    string  `json:"id"`
	Price float64 `json:"price"`
	Rule  string  `json:"rule"`
}

const URL = "https://api.coinlore.com/api/ticker/?id=90"
const File = "CryptoInfo.json"

//APSIRASYTI
func main() {

	var CryptoC Crypto
	var CryptoRules []CryptoRules
	var err error
	//!!Pasikeisti
	//	if _, err := http.Get("https://api.telegram.org/bot717631082:AAEaOBNtLs8tJ-DnoWTbCk1Y2i6mawum3jk/sendMessage?chat_id=@CryptTelegram&text=Hello"); err != nil {
	//		fmt.Println(err)
	//	}
	//KAS CIA VYKSTA APRASYTI
	if CryptoC, err = FromURL(URL); err != nil {
		//	fmt.Println(errStart.Error())
		log.Fatal(err)
	}

	//for _, value := range CryptoC {
	fmt.Println(CryptoC)
	//}
	//KAS CIA VYKSTA APRASYTI
	if CryptoRules, err = FromFile(File); err != nil {
		//	fmt.Println(errStart.Error())
		log.Fatal(err)
	}

	//	json.Unmarshal(byteValue, &CryptInfo)
	var Price float64
	for _, x := range CryptoRules {
		fmt.Printf("%s \n", x.Rule)
		Price = x.Price
	}
	//Tikrinu duomenis pagal taisykle is failo
	if Price > CryptoC.Price {
		//Authenticatinu BOTa Uztenka viena karta prie MAIN
		bot, err := telegram.NewBotAPI("717631082:AAEaOBNtLs8tJ-DnoWTbCk1Y2i6mawum3jk")
		log.Printf("Authorized on account %s", bot.Self.UserName)

		if err != nil {
			log.Panic(err)
		}

		//Botas printina zinute i kanala
		text := fmt.Sprintf("%s is running. %s price has increased!", bot.Self.UserName, CryptoC.Name)
		msg := telegram.NewMessageToChannel("@CryptTelegram", text)
		bot.Send(msg)
	}

}

//func CheckIt()

//Nuskaitau JSON is failo su crypto currency info ATSKIRA FUNKCIJA
func FromFile(FileName string) ([]CryptoRules, error) {
	var CryptInfo []CryptoRules
	var err error
	var jsonFile []byte

	if jsonFile, err = ioutil.ReadFile(FileName); err != nil {
		// if we os.Open returns an error then handle it
		return CryptInfo, err
	}

	if err = json.Unmarshal(jsonFile, &CryptInfo); err != nil {
		return CryptInfo, err
	}
	return CryptInfo, err
}

// Function FromURL takes URL address of a JSON, unmarshals JSON and return the data
func FromURL(URLName string) (Crypto, error) {
	spaceClient := http.Client{
		Timeout: time.Second * 10, // 2 secs
	}
	type CryptoJSON struct {
		ID    string `json:"id"`
		Name  string `json:"name"`
		Price string `json:"price_usd"`
	}
	exrates := Crypto{}
	CrypJson := []CryptoJSON{}

	var err error
	var req *http.Request
	var res *http.Response
	var body []byte

	//	717631082:AAEa0BNtLs8tJ-DnoWTbCk1Y2i6mawum3jk

	if req, err = http.NewRequest(http.MethodGet, URLName, nil); err != nil {
		return exrates, err
		//fmt.Println(err.Error())
	}
	if res, err = spaceClient.Do(req); err != nil {
		return exrates, err
		//log.Fatal(getErr)
	}
	if body, err = ioutil.ReadAll(res.Body); err != nil {
		return exrates, err
		//
	}
	//exrates := ExRates{}
	// json.Unmarshal(content, &friends)
	if err = json.Unmarshal(body, &CrypJson); err != nil {
		return exrates, err
	}
	//!!PASIDARYTI
	// if len != 1 {fail}
	if len(CrypJson) != 1 {
		log.Fatal(err)
	}
	exrates.ID = CrypJson[0].ID
	exrates.Name = CrypJson[0].Name

	f := CrypJson[0].Price

	if s, err := strconv.ParseFloat(f, 64); err == nil {
		//	fmt.Println(s) // 3.14159265
		exrates.Price = s
	}
	//realrate := Crypro{Price: flat(exrates[0].price_usd)}

	return exrates, nil
}
