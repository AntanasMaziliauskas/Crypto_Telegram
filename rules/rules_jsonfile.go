package rules

import (
	"encoding/json"
	"io/ioutil"
	"log"

	"github.com/AntanasMaziliauskas/Crypto_Telegram/types"
)

//RulesFromJSON structure
type RulesFromJSON struct {
	Path string
}

//Init function does nothing
func (r *RulesFromJSON) Init() error {

	return nil
}

//ReadRules function read throught the JSON file and unmarshals it
func (r *RulesFromJSON) ReadRules() ([]types.Rule, error) {
	var (
		CryptInfo []types.Rule
		err       error
		jsonFile  []byte
	)

	if jsonFile, err = ioutil.ReadFile(r.Path); err != nil {
		return nil, err
	}
	if err = json.Unmarshal(jsonFile, &CryptInfo); err != nil {
		return nil, err
	}

	return CryptInfo, nil
}

// Match function goes through the rules list with data from URL
// looking for any matching rule. Makes a new list of matched rules
func (r *RulesFromJSON) Match(rules []types.Rule, data []types.LoreData) []types.Rule {
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

//One function checks if data from API matched specific rule
func (r *RulesFromJSON) One(rule types.Rule, data types.LoreData) bool {
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
func (r *RulesFromJSON) SaveRules(updatedRules []types.Rule) error {
	var (
		err    error
		tofile []byte
	)

	if tofile, err = json.Marshal(updatedRules); err != nil {
		return err
	}

	return ioutil.WriteFile(r.Path, tofile, 0644)
}
