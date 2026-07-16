package localfinanceevidence

import "math"

func expectedCreditPreview(value CreditImportPreview) bool {
	if value.TemplateID == "synthetic-card-export" && (value.TemplateVersion == 1 || value.TemplateVersion == 2) {
		return value.IssuerKey == "synthetic-card" && value.TransactionCount == 3 && value.DebitMinor == 20900 &&
			value.CreditMinor == 2200 && value.SuggestionCount == 3 && value.ExpectedBalance.OpeningMinor == 40000 &&
			value.ExpectedBalance.ClosingMinor == 58700
	}
	if value.TemplateID == "normalized-credit-csv" && value.TemplateVersion == 1 {
		return value.IssuerKey == "synthetic-normalized" && value.TransactionCount == 2 && value.DebitMinor == 12500 &&
			value.CreditMinor == 2200 && value.SuggestionCount == 0 && value.ExpectedBalance.OpeningMinor == 10000 &&
			value.ExpectedBalance.ClosingMinor == 20300
	}
	return false
}

func validCreditBalance(value CreditImportPreview) bool {
	opening := value.ExpectedBalance.OpeningMinor
	if value.DebitMinor > 0 && opening > math.MaxInt64-value.DebitMinor {
		return false
	}
	closing := opening + value.DebitMinor
	if value.CreditMinor > 0 && closing < math.MinInt64+value.CreditMinor {
		return false
	}
	closing -= value.CreditMinor
	return value.CalculatedClosingMinor == closing && value.ExpectedBalance.ClosingMinor == closing &&
		value.BalanceDeltaMinor == 0
}
