package financeconnector

func AssessLiveParity(transactions []SourceTransaction) ParityReport {
	contract := make([]ParityField, 0, len(CanonicalFieldContract)-1)
	for _, field := range CanonicalFieldContract {
		if field.Name != "card_id" {
			contract = append(contract, field)
		}
	}
	report := assessParity(transactions, contract)
	report.UnsupportedProviderFeatures = append(
		report.UnsupportedProviderFeatures,
		"card_id_not_available_from_bank_transactions",
	)
	return report
}
