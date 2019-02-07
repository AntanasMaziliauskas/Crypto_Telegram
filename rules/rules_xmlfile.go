package rules

import (
	"encoding/xml"
	"io/ioutil"
	"log"

	"github.com/AntanasMaziliauskas/Crypto_Telegram/types"
)

//RulesFromXML structure holds path of file and implements RulesService interface
type RulesFromXML struct {
	Path string
}

//XML structure holds the name of XML content and the structure of file content
type XML struct {
	XMLName xml.Name     `xml:"rules"`
	Rules   []types.Rule `xml:"rule"`
}

//Init function does nothing
func (r *RulesFromXML) Init() error {

	return nil
}

//ReadRules function reads throught the XML file and unmarshals it
func (r *RulesFromXML) ReadRules() ([]types.Rule, error) {
	var (
		xmlRules XML
		err      error
		xmlFile  []byte
	)

	if xmlFile, err = ioutil.ReadFile(r.Path); err != nil {
		return nil, err
	}
	if err = xml.Unmarshal(xmlFile, &xmlRules); err != nil {
		return nil, err
	}

	return xmlRules.Rules, nil
}

// Match function goes through the rules list with data from URL
// looking for any matching rule. Makes a new list of matched rules
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

//One function checks if data from API matches specific rule
func (r *RulesFromXML) One(rule types.Rule, data types.LoreData) bool {
	if data.ID == rule.ID {
		if rule.Rule == "gt" && rule.Price < data.Price {
			return true
		}
		if rule.Rule == "lt" && rule.Price > data.Price {
			return true
		}
	}

	return false
}

//SaveRules function saves the updated rules to a file
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
