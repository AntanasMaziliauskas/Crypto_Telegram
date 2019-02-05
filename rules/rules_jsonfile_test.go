package rules_test

import (
	"testing"

	"github.com/AntanasMaziliauskas/Crypto_Telegram/rules"
	"github.com/AntanasMaziliauskas/Crypto_Telegram/types"
	"github.com/stretchr/testify/assert"
)

//
func TestImplements(t *testing.T) {
	assert.Implements(t, (*rules.RulesService)(nil), &rules.RulesFromJSON{})
}

func TestMatch(t *testing.T) {

	ts := rules.RulesFromJSON{}

	testTable := []struct {
		rule   []types.Rule
		data   []types.LoreData
		expect []types.Rule
	}{
		{
			rule: []types.Rule{
				{
					RuleID: 0,
					ID:     "90",
					Price:  3000.01,
					Rule:   "gt",
				},
				{
					RuleID: 1,
					ID:     "80",
					Price:  3000.01,
					Rule:   "lw",
				},
				{
					RuleID: 2,
					ID:     "91",
					Price:  3000.01,
					Rule:   "lw",
				},
			},
			data: []types.LoreData{
				{
					ID:    "90",
					Name:  "Bitcoin",
					Price: 3396.06,
				},
				{
					ID:    "80",
					Name:  "Ethereum",
					Price: 3396.06,
				},
				{
					ID:    "91",
					Name:  "ClubCoin",
					Price: 3396.06,
				},
			},
			expect: []types.Rule{
				{
					RuleID: 0,
					ID:     "90",
					Price:  3000.01,
					Rule:   "gt",
				},
			},
		},
		{
			rule: []types.Rule{
				{
					RuleID: 0,
					ID:     "90",
					Price:  3000.01,
					Rule:   "gt",
				},
				{
					RuleID: 1,
					ID:     "80",
					Price:  3000.01,
					Rule:   "lw",
				},
				{
					RuleID: 2,
					ID:     "91",
					Price:  3000.01,
					Rule:   "lw",
				},
			},
			data: []types.LoreData{
				{
					ID:    "90",
					Name:  "Bitcoin",
					Price: 3396.06,
				},
				{
					ID:    "80",
					Name:  "Ethereum",
					Price: 3000,
				},
				{
					ID:    "91",
					Name:  "ClubCoin",
					Price: 10.06,
				},
			},
			expect: []types.Rule{
				{
					RuleID: 0,
					ID:     "90",
					Price:  3000.01,
					Rule:   "gt",
				},
				{
					RuleID: 1,
					ID:     "80",
					Price:  3000.01,
					Rule:   "lw",
				},
				{
					RuleID: 2,
					ID:     "91",
					Price:  3000.01,
					Rule:   "lw",
				},
			},
		},
		{
			rule: []types.Rule{
				{
					RuleID: 0,
					ID:     "93",
					Price:  3000.01,
					Rule:   "gt",
				},
				{
					RuleID: 1,
					ID:     "80",
					Price:  3000.01,
					Rule:   "lw",
				},
				{
					RuleID: 2,
					ID:     "91",
					Price:  3000.01,
					Rule:   "lw",
				},
			},
			data: []types.LoreData{
				{
					ID:    "94",
					Name:  "Factom",
					Price: 3396.06,
				},
				{
					ID:    "100",
					Name:  "Safe Exchange Coin",
					Price: 3000,
				},
				{
					ID:    "11",
					Name:  "Einsteinium",
					Price: 10.06,
				},
			},
			expect: nil,
		},
	}
	for i, v := range testTable {
		result := ts.Match(v.rule, v.data)
		assert.Equal(t, result, v.expect, "Case %d", i)
	}
}

func TestOne(t *testing.T) {

	ts := rules.RulesFromJSON{}

	testTable := []struct {
		rule   types.Rule
		data   types.LoreData
		expect bool
	}{
		{
			rule: types.Rule{
				ID:    "90",
				Price: 3000.06,
				Rule:  "gt",
			},
			data: types.LoreData{
				ID:    "90",
				Name:  "Bitcoin",
				Price: 3396.06,
			},
			expect: true,
		},
		{
			rule: types.Rule{
				ID:    "90",
				Price: 2222.22,
				Rule:  "lw",
			},
			data: types.LoreData{
				ID:    "80",
				Name:  "Ethereum",
				Price: 104.38,
			},

			expect: false,
		},
		{
			rule: types.Rule{
				ID:    "80",
				Price: 110.22,
				Rule:  "lw",
			},
			data: types.LoreData{
				ID:    "80",
				Name:  "Ethereum",
				Price: 104.38,
			},

			expect: true,
		},
	}
	for i, v := range testTable {
		result := ts.One(v.rule, v.data)
		assert.Equal(t, result, v.expect, "Case %d", i)
	}
}
