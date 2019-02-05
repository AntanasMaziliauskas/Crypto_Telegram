package rules_test

import (
	"testing"

	"github.com/AntanasMaziliauskas/Crypto_Telegram/rules"
	"github.com/AntanasMaziliauskas/Crypto_Telegram/types"
	"github.com/stretchr/testify/assert"
)

func TestUniqueRules(t *testing.T) {
	testTable := []struct {
		rule   []types.Rule
		expect []string
	}{
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
			expect: []string{
				"93",
				"80",
				"91",
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
					ID:     "93",
					Price:  3000.01,
					Rule:   "lw",
				},
				{
					RuleID: 2,
					ID:     "93",
					Price:  3000.01,
					Rule:   "lw",
				},
			},

			expect: []string{
				"93",
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
					ID:     "93",
					Price:  3000.01,
					Rule:   "lw",
				},
			},
			expect: []string{
				"93",
				"80",
			},
		},
	}

	for i, v := range testTable {
		result := rules.UniqueRules(v.rule)
		assert.Equal(t, result, v.expect, "Case %d", i)
	}
}
