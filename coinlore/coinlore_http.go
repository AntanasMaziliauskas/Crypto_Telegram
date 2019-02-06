package coinlore

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"

	"github.com/AntanasMaziliauskas/Crypto_Telegram/types"
)

//CoinloreAPI struct
type CoinloreAPI struct {
	sourceURL string
}

// Init function gives URL value to c.sourceURL
func (c *CoinloreAPI) Init() error {
	c.sourceURL = URL

	return nil
}

//FetchAll function received a list of Crypto Currency IDs, uses prepareURL function
//to generate URL which then used to get JSON, unmarshals JSON and returns the data received
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
		CrypJSON    []CryptoJSON
		DataFromURL []types.LoreData
	)

	for _, u := range ids {

		url := c.prepareURL(u)

		spaceClient := http.Client{
			Timeout: time.Second * 10,
		}

		if req, err = http.NewRequest(http.MethodGet, url, nil); err != nil {

			return nil, err
		}
		if res, err = spaceClient.Do(req); err != nil {

			return nil, err
		}
		if body, err = ioutil.ReadAll(res.Body); err != nil {

			return nil, err
		}
		if err = json.Unmarshal(body, &CrypJSON); err != nil {

			return nil, err
		}
		if len(CrypJSON) < 1 {

			return DataFromURL, err
		}
		p, _ := strconv.ParseFloat(CrypJSON[0].Price, 64)

		DataFromURL = append(DataFromURL, types.LoreData{
			ID:    CrypJSON[0].ID,
			Name:  CrypJSON[0].Name,
			Price: p,
		})
	}

	return DataFromURL, err
}

//FetchOne function takes one ID of a specific Crypto Currency and uses prepareURL
//function to make whole URL address, then retrieves JSON, unmarshals JSON and return the data
func (c *CoinloreAPI) FetchOne(id string) (types.LoreData, error) {
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
		CrypJSON    CryptoJSON
		DataFromURL types.LoreData
	)

	url := c.prepareURL(id)

	spaceClient := http.Client{
		Timeout: time.Second * 10,
	}

	if req, err = http.NewRequest(http.MethodGet, url, nil); err != nil {

		return DataFromURL, err
	}
	if res, err = spaceClient.Do(req); err != nil {

		return DataFromURL, err
	}
	if body, err = ioutil.ReadAll(res.Body); err != nil {

		return DataFromURL, err
	}
	if err = json.Unmarshal(body, &CrypJSON); err != nil {

		return DataFromURL, err
	}

	p, _ := strconv.ParseFloat(CrypJSON.Price, 64)

	DataFromURL = types.LoreData{
		ID:    CrypJSON.ID,
		Name:  CrypJSON.Name,
		Price: p,
	}

	return DataFromURL, err
}

//prepareURL function received id of a Crypto Currency and adds it to the URL
func (c *CoinloreAPI) prepareURL(id string) string {
	url := fmt.Sprintf(c.sourceURL, id)

	return url
}
