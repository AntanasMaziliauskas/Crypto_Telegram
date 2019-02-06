package application_test

import (
	"testing"

	"github.com/AntanasMaziliauskas/Crypto_Telegram/application"
	"github.com/AntanasMaziliauskas/Crypto_Telegram/types"
	"github.com/stretchr/testify/assert"
)

func TestStatus(t *testing.T) {
	testTable := []struct {
		rule   types.Rule
		expect string
	}{
		{

			rule: types.Rule{
				RuleID: 0,
				ID:     "93",
				Price:  3000.01,
				Rule:   "gt",
			},
			expect: "increased",
		},
		{
			rule: types.Rule{
				RuleID: 0,
				ID:     "93",
				Price:  3000.01,
				Rule:   "lt",
			},
			expect: "decreased",
		},
		{
			rule: types.Rule{
				RuleID: 0,
				ID:     "93",
				Price:  3000.01,
				Rule:   "lt",
			},
			expect: "decreased",
		},
	}
	for i, v := range testTable {
		result := application.Status(v.rule)
		assert.Equal(t, result, v.expect, "Case %d", i)
	}

}
