package rules_test


func TestReadRules(t *testing.T){

	ts:= rules.RulesFromXML{}

	testTable := struct {
		expectRules []types.Rule
	}
	{
		expectRules: []types.Rule{
		{
			RuleID: 0,
			ID:     "90",
			Price:  3470.98,
			Rule:   "lw",
		},
		{
			RuleID: 1,
			ID:     "90",
			Price:  3470.98,
			Rule:   "gt",
		},
		{
			RuleID: 2,
			ID:     "91",
			Price:  3470.98,
			Rule:   "lw",
		},
		{
			RuleID: 3,
			ID:     "92",
			Price:  3470.98,
			Rule:   "lw",
		},
	},
},
	}
}