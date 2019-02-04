package crypto_test

import (
	"testing"

	"github.com/AntanasMaziliauskas/Crypto_Telegram/crypto"
	"github.com/influxdata/influxdb/pkg/testing/assert"
)

type Test struct {
	CS crypto.CryptoStruct
}

func TestCheckOne(t *testing.T) {
	testTable := []struct {
		crypt  crypto.Crypto
		rule   crypto.CryptoRule
		expect bool
	}{
		{
			crypt: crypto.Crypto{
				ID:    "90",
				Name:  "Bitcoin",
				Price: 3396.06,
			},
			rule: crypto.CryptoRule{
				ID:    "90",
				Price: 3000.01,
				Rule:  "gt",
			},
			expect: true,
		},
		{
			crypt: crypto.Crypto{
				ID:    "80",
				Name:  "Ethereum",
				Price: 104.38,
			},
			rule: crypto.CryptoRule{
				ID:    "90",
				Price: 2222.22,
				Rule:  "lw",
			},
			expect: false,
		},
		{
			crypt: crypto.Crypto{
				ID:    "80",
				Name:  "Ethereum",
				Price: 104.38,
			},
			rule: crypto.CryptoRule{
				ID:    "80",
				Price: 110.22,
				Rule:  "lw",
			},
			expect: true,
		},
	}
	for i, v := range testTable {
		result := crypto.CheckOne(v.crypt, v.rule)
		assert.Equal(t, result, v.expect, "Case %d", i)
	}
}

/*func TestCheckAll(t *testing.T) {
	testTable := []struct {
		crypt  crypto.Crypto
		rule   []crypto.CryptoRule
		expect []crypto.CryptoRule
	}{
		{
			crypt: crypto.Crypto{
				ID:    "90",
				Name:  "Bitcoin",
				Price: 3396.06,
			},
			rule: []crypto.CryptoRule{
				{
					ID:       "90",
					Price:    3000.01,
					Rule:     "gt",
					Notified: true,
				},
				{
					ID:    "90",
					Price: 3000.01,
					Rule:  "lw",
				},
			},
			expect: nil,
		},
		{
			crypt: crypto.Crypto{
				ID:    "90",
				Name:  "Bitcoin",
				Price: 3396.06,
			},
			rule: []crypto.CryptoRule{
				{
					ID:    "80",
					Price: 3000.01,
					Rule:  "gt",
				},
				{
					ID:    "90",
					Price: 4000.02,
					Rule:  "gt",
				},
			},
			expect: nil,
		},
		{
			crypt: crypto.Crypto{
				ID:    "80",
				Name:  "Ethereum",
				Price: 104.38,
			},
			rule: []crypto.CryptoRule{
				{
					ID:    "80",
					Price: 100.01,
					Rule:  "gt",
				},
				{
					ID:    "80",
					Price: 110.10,
					Rule:  "lw",
				},
			},
			expect: []crypto.CryptoRule{
				{
					ID:    "80",
					Price: 100.01,
					Rule:  "gt",
				},
				{
					ID:    "80",
					Price: 110.10,
					Rule:  "lw",
				},
			},
		},
	}
	for i, v := range testTable {
		result := crypto.CheckAll(v.crypt, v.rule)
		assert.Equal(t, result, v.expect, "Case %d", i)
	}
}*/
func TestNotify(t *testing.T) {
	testTable := []struct {
		real   []crypto.CryptoRule
		sent   []crypto.CryptoRule
		edited []crypto.CryptoRule
	}{
		{
			real: []crypto.CryptoRule{
				{
					RuleID:   0,
					ID:       "90",
					Price:    3000.01,
					Rule:     "gt",
					Notified: true,
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
			sent: []crypto.CryptoRule{
				{
					RuleID:   1,
					ID:       "80",
					Price:    3000.01,
					Rule:     "lw",
					Notified: true,
				},
			},
			edited: []crypto.CryptoRule{
				{
					RuleID:   0,
					ID:       "90",
					Price:    3000.01,
					Rule:     "gt",
					Notified: true,
				},
				{
					RuleID:   1,
					ID:       "80",
					Price:    3000.01,
					Rule:     "lw",
					Notified: true,
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
			real: []crypto.CryptoRule{
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
			sent: []crypto.CryptoRule{
				{
					RuleID:   1,
					ID:       "80",
					Price:    3000.01,
					Rule:     "lw",
					Notified: true,
				},
			},
			edited: []crypto.CryptoRule{
				{
					RuleID:   0,
					ID:       "90",
					Price:    3000.01,
					Rule:     "gt",
					Notified: false,
				},
				{
					RuleID:   1,
					ID:       "80",
					Price:    3000.01,
					Rule:     "lw",
					Notified: true,
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
			real: []crypto.CryptoRule{
				{
					RuleID:   0,
					ID:       "90",
					Price:    3000.01,
					Rule:     "gt",
					Notified: true,
				},
				{
					RuleID:   1,
					ID:       "80",
					Price:    3000.01,
					Rule:     "lw",
					Notified: true,
				},
				{
					RuleID:   2,
					ID:       "91",
					Price:    3000.01,
					Rule:     "lw",
					Notified: true,
				},
			},
			sent: nil,
			edited: []crypto.CryptoRule{
				{
					RuleID:   0,
					ID:       "90",
					Price:    3000.01,
					Rule:     "gt",
					Notified: true,
				},
				{
					RuleID:   1,
					ID:       "80",
					Price:    3000.01,
					Rule:     "lw",
					Notified: true,
				},
				{
					RuleID:   2,
					ID:       "91",
					Price:    3000.01,
					Rule:     "lw",
					Notified: true,
				},
			},
		},
	}
	for i, v := range testTable {
		result := crypto.Notify()
		assert.Equal(t, result, v.edited, "Case %d", i)
	}
}

/*
var Zirafe struct {
	Ugis int
}

func TestZirafa(t *testing.T) {
	x := Zirafe{180}
	x.Paaugo(5)
	assert.Equal(t, x.Ugis, 185)
}
*/
