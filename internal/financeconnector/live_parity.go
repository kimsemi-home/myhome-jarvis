package financeconnector

func AssessLiveParity(transactions []SourceTransaction) ParityReport {
	contract := liveContract(transactions)
	report := assessParity(transactions, contract)
	if len(contract) < len(CanonicalFieldContract) {
		report.UnsupportedProviderFeatures = append(report.UnsupportedProviderFeatures, "card_id_not_available_from_bank_transactions")
	}
	return report
}

func liveContract(transactions []SourceTransaction) []ParityField {
	for _, transaction := range transactions {
		if transaction.CardID != "" {
			return append([]ParityField(nil), CanonicalFieldContract...)
		}
	}
	contract := make([]ParityField, 0, len(CanonicalFieldContract)-1)
	for _, field := range CanonicalFieldContract {
		if field.Name != "card_id" {
			contract = append(contract, field)
		}
	}
	return contract
}
