package rules

import (
	"github.com/AntanasMaziliauskas/Crypto_Telegram/coinlore"
	"github.com/AntanasMaziliauskas/Crypto_Telegram/types"
)

//RulesService is the interface that wraps Init, ReadRules, SaveRules, Match and One methods
type RulesService interface {
	Init() error
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
