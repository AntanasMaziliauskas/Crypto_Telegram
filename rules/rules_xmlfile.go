package rules

import (
	"encoding/xml"
	"io/ioutil"
	"log"

	"github.com/AntanasMaziliauskas/Crypto_Telegram/types"
)

type RulesFromXML struct {
	Path string
}
type XML struct {
	XMLName xml.Name     `xml:"rules"`
	Rules   []types.Rule `xml:"rule"`
}

func (r *RulesFromXML) Init() error {

	return nil
}

func (r *RulesFromXML) ReadRules() ([]types.Rule, error) {
	var (
		xmlRules XML
		err      error
		xmlFile  []byte
	)

	if xmlFile, err = ioutil.ReadFile(r.Path); err != nil {
		return nil, err
	}
	//var c RulesFromXML
	if err = xml.Unmarshal(xmlFile, &xmlRules); err != nil {
		return nil, err
	}
	//fmt.Println(xmlRules.Rules)
	return xmlRules.Rules, nil
}

func (r *RulesFromXML) Match(rules []types.Rule, data []types.LoreData) []types.Rule {
	var matched []types.Rule

	for _, d := range data {
		for _, x := range rules {
			if r.One(x, d) && !x.Notified {
				matched = append(matched, x)
			}
		}
	}
	if len(matched) == 0 {
		log.Printf("No rules matched.\n")
	}

	return matched
}

func (r *RulesFromXML) One(rule types.Rule, data types.LoreData) bool {
	if data.ID == rule.ID {
		if rule.Rule == "gt" && rule.Price < data.Price {
			return true
		}
		if rule.Rule == "lw" && rule.Price > data.Price {
			return true
		}
	}

	return false
}

func (r *RulesFromXML) SaveRules(updatedRules []types.Rule) error {
	var (
		err          error
		tofile       []byte
		rulesToWrite XML
	)
	rulesToWrite.Rules = updatedRules

	if tofile, err = xml.Marshal(rulesToWrite); err != nil {
		return err
	}
	return ioutil.WriteFile(r.Path, tofile, 0644)
}
