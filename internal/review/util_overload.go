package review

func mapOverloadRules(rules []OverloadRule) map[string]OverloadRule {
	mapped := map[string]OverloadRule{}
	for _, rule := range rules {
		rule.Key = normalizeToken(rule.Key)
		if rule.Key != "" {
			mapped[rule.Key] = rule
		}
	}
	return mapped
}
