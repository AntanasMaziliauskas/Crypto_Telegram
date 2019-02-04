package rules

import (
	"github.com/AntanasMaziliauskas/Crypto_Telegram/coinlore"
	"github.com/AntanasMaziliauskas/Crypto_Telegram/types"
)

const File = "CryptoInfo.json"

type RulesService interface {
	Init()
	ReadRules() ([]types.Rule, error)
	SaveRules([]types.Rule) error
	Match([]types.Rule, []types.LoreData) []types.Rule
	One(types.Rule, types.LoreData) bool
}

//UniqueRules picks out all of the different crypto currency ID's from rules list
func UniqueRules(rules []types.Rule) []string {
	var ids []string

	for _, v := range rules {
		if !coinlore.SliceContainsString(v.ID, ids) {
			ids = append(ids, v.ID)
		}
	}

	return ids
}
