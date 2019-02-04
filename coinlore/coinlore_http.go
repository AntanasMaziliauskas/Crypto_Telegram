package coinlore

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"

	"github.com/AntanasMaziliauskas/Crypto_Telegram/types"
)

type CoinloreAPI struct {
	sourceURL string
}

func (c *CoinloreAPI) Init() {
	c.sourceURL = URL
}

// Function FromURL takes URL address of a JSON, unmarshals JSON and return the data
func (c *CoinloreAPI) FetchAll(ids []string) ([]types.LoreData, error) {
	type CryptoJSON struct {
		ID    string `json:"id"`
		Name  string `json:"name"`
		Price string `json:"price_usd"`
	}

	var (
		err         error
		req         *http.Request
		res         *http.Response
		body        []byte
		CrypJson    []CryptoJSON
		DataFromURL []types.LoreData
	)

	for _, u := range ids {

		url := c.prepareURL(u)

		spaceClient := http.Client{
			Timeout: time.Second * 10, // 2 secs
		}

		if req, err = http.NewRequest(http.MethodGet, url, nil); err != nil {
			//	fmt.Println("#1")
			return nil, err
		}
		if res, err = spaceClient.Do(req); err != nil {
			//	fmt.Println("#2")
			return nil, err
		}
		if body, err = ioutil.ReadAll(res.Body); err != nil {
			//	fmt.Println("#3")
			return nil, err
		}
		if err = json.Unmarshal(body, &CrypJson); err != nil {
			//	fmt.Println("#4")
			return nil, err
		}
		if len(CrypJson) < 1 {
			//	fmt.Println("#4")
			return nil, errors.New("Json empty")
		}
		//fmt.Println("#praejau")
		//fmt.Println(len(CrypJson))
		p, _ := strconv.ParseFloat(CrypJson[0].Price, 64)

		DataFromURL = append(DataFromURL, types.LoreData{
			ID:    CrypJson[0].ID,
			Name:  CrypJson[0].Name,
			Price: p,
		})
	}
	return DataFromURL, err
}

//prepareURL function prepares URL for reading
func (c *CoinloreAPI) prepareURL(id string) string {
	url := fmt.Sprintf(c.sourceURL, id)

	return url
}
