package crypto_test

import (
	"testing"

	crypto "github.com/AntanasMaziliauskas/Crypto_Telegram/Crypto"
	"github.com/influxdata/influxdb/pkg/testing/assert"
)

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
func TestCheckAll(t *testing.T) {
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
					ID:    "90",
					Price: 3000.01,
					Rule:  "gt",
				},
				{
					ID:    "90",
					Price: 3000.01,
					Rule:  "lw",
				},
			},
			expect: []crypto.CryptoRule{
				{
					ID:    "90",
					Price: 3000.01,
					Rule:  "gt",
				},
			},
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
}
