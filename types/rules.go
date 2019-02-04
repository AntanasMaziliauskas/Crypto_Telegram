package types

// Rule is fetched from file
type Rule struct {
	RuleID   int     `json:"ruleid"`
	ID       string  `json:"id"`
	Price    float64 `json:"price"`
	Rule     string  `json:"rule"`
	Notified bool    `json:"notified"`
}

// LoreData is fetched from URL
type LoreData struct {
	ID    string  `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price_usd"`
}
