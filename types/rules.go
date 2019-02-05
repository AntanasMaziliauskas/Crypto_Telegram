package types

// Rule is fetched from file
type Rule struct {
	//	XMLName  string  `json:"" xml:"rules"`
	RuleID   int     `json:"ruleid" xml:"ruleid"`
	ID       string  `json:"id" xml:"id"`
	Price    float64 `json:"price" xml:"price"`
	Rule     string  `json:"rule" xml:"rule"`
	Notified bool    `json:"notified" xml:"notified"`
}

// LoreData is fetched from URL
type LoreData struct {
	ID    string  `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price_usd"`
}

//RuleXML is fetched from XML
type RuleXML struct {
	RuleID   string ``
	ID       string `xml:"id"`
	Price    string `xml:"price"`
	Rule     string `xml:"rule"`
	Notified string `xml:"notified"`
}
