package coinlore

import "github.com/AntanasMaziliauskas/Crypto_Telegram/types"

//URL is the unprepared address of the Cryto Currency information.
//It needs ID to be added to make it valid
const URL = "https://api.coinlore.com/api/ticker/?id=%s"

//CoinloreService interface
type CoinloreService interface {
	Init() error
	FetchAll(ids []string) ([]types.LoreData, error)
	FetchOne(id string) (types.LoreData, error)
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
