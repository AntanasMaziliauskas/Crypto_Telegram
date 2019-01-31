package Crypto_test

import (
	"testing"

	Crypto "github.com/AntanasMaziliauskas/Crypto_Telegram/Crypto"
	"github.com/influxdata/influxdb/pkg/testing/assert"
)

func TestCheckOne(t *testing.T) {
	testTable := []struct {
		crypt  Crypto.Crypto
		rule   Crypto.CryptoRule
		expect bool
	}{
		{
			crypt: Crypto.Crypto{
				ID:    "90",
				Name:  "Bitcoin",
				Price: 3396.06,
			},
			rule: Crypto.CryptoRule{
				ID:    "90",
				Price: 3000.01,
				Rule:  "gt",
			},
			expect: true,
		},
		{
			crypt: Crypto.Crypto{
				ID:    "80",
				Name:  "Ethereum",
				Price: 104.38,
			},
			rule: Crypto.CryptoRule{
				ID:    "90",
				Price: 2222.22,
				Rule:  "lw",
			},
			expect: false,
		},
		{
			crypt: Crypto.Crypto{
				ID:    "80",
				Name:  "Ethereum",
				Price: 104.38,
			},
			rule: Crypto.CryptoRule{
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
func TestCheckAll(t *testing.T) {
	testTable := []struct {
		crypt  Crypto.Crypto
		rule   []Crypto.CryptoRule
		expect []Crypto.CryptoRule
	}{
		{
			crypt: Crypto.Crypto{
				ID:    "90",
				Name:  "Bitcoin",
				Price: 3396.06,
			},
			rule: []Crypto.CryptoRule{
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
			crypt: Crypto.Crypto{
				ID:    "90",
				Name:  "Bitcoin",
				Price: 3396.06,
			},
			rule: []Crypto.CryptoRule{
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
			crypt: Crypto.Crypto{
				ID:    "80",
				Name:  "Ethereum",
				Price: 104.38,
			},
			rule: []Crypto.CryptoRule{
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
			expect: []Crypto.CryptoRule{
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
}
func TestNotify(t *testing.T) {
	testTable := []struct {
		real   []Crypto.CryptoRule
		sent   []Crypto.CryptoRule
		edited []Crypto.CryptoRule
	}{
		{
			real: []Crypto.CryptoRule{
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
			sent: []Crypto.CryptoRule{
				{
					RuleID:   1,
					ID:       "80",
					Price:    3000.01,
					Rule:     "lw",
					Notified: true,
				},
			},
			edited: []Crypto.CryptoRule{
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
			real: []Crypto.CryptoRule{
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
			sent: []Crypto.CryptoRule{
				{
					RuleID:   1,
					ID:       "80",
					Price:    3000.01,
					Rule:     "lw",
					Notified: true,
				},
			},
			edited: []Crypto.CryptoRule{
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
			real: []Crypto.CryptoRule{
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
			edited: []Crypto.CryptoRule{
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
		result := crypto.Notify(v.real, v.sent)
		assert.Equal(t, result, v.edited, "Case %d", i)
	}
}
