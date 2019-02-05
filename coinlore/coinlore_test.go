package coinlore_test

import (
	"testing"

	"github.com/AntanasMaziliauskas/Crypto_Telegram/coinlore"
	"github.com/stretchr/testify/assert"
)

func TestSliceContainsString(t *testing.T) {
	testTable := []struct {
		needle   string
		haystack []string
		expect   bool
	}{
		{
			needle: "80",
			haystack: []string{
				"90",
				"10",
				"80",
			},
			expect: true,
		},
		{
			needle: "105",
			haystack: []string{
				"90",
				"10",
				"80",
			},
			expect: false,
		},
		{
			needle: "80",
			haystack: []string{
				"90",
				"80",
				"80",
			},
			expect: true,
		},
	}

	for i, v := range testTable {
		result := coinlore.SliceContainsString(v.needle, v.haystack)
		assert.Equal(t, result, v.expect, "Case %d", i)
	}
}
