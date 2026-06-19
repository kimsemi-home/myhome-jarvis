package domain

import "strings"

func recordCurrency(currencies map[string]bool, currency string) {
	if strings.TrimSpace(currency) != "" {
		currencies[currency] = true
	}
}

func recordOwnerCurrency(
	currencies map[string]map[string]bool,
	owner string,
	currency string,
) {
	if strings.TrimSpace(currency) == "" {
		return
	}
	if _, ok := currencies[owner]; !ok {
		currencies[owner] = map[string]bool{}
	}
	currencies[owner][currency] = true
}

func summaryCurrency(currencies map[string]bool) string {
	switch len(currencies) {
	case 0:
		return ""
	case 1:
		for currency := range currencies {
			return currency
		}
	}
	return "mixed"
}

func firstCurrency(values ...string) string {
	for _, value := range values {
		if strings.TrimSpace(value) != "" {
			return value
		}
	}
	return ""
}
