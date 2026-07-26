package financeconnector

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
)

const unknownMerchant = "미제공"

func currency(value string) string {
	if strings.TrimSpace(value) == "" {
		return "KRW"
	}
	return strings.TrimSpace(value)
}

func normalizedCurrency(value string) string {
	value = strings.TrimSpace(value)
	if value == "" {
		return "KRW"
	}
	return value
}

func parseMyDataAmount(raw json.RawMessage) (int64, error) {
	text := strings.Trim(strings.TrimSpace(string(raw)), "\"")
	if amount, err := strconv.ParseInt(text, 10, 64); err == nil {
		return amount, nil
	}
	amount, err := strconv.ParseFloat(text, 64)
	if err != nil {
		return 0, fmt.Errorf("invalid amount")
	}
	return int64(amount + 0.5), nil
}

func myDataDirection(code string) (string, error) {
	switch strings.TrimSpace(code) {
	case "03", "04", "06", "09", "98":
		return "credit", nil
	case "02", "05", "07", "10", "99":
		return "debit", nil
	default:
		return "", fmt.Errorf("unsupported MyData transaction type %q", code)
	}
}
