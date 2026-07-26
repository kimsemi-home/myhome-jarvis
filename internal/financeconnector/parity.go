package financeconnector

import "sort"

func AssessParity(transactions []SourceTransaction) ParityReport {
	return assessParity(transactions, CanonicalFieldContract)
}

func assessParity(transactions []SourceTransaction, contract []ParityField) ParityReport {
	report := ParityReport{
		Provider: ProviderTossMyData, Records: len(transactions),
		ExpectedFields: len(contract), ThresholdPercent: ParityThreshold,
		FieldContract: append([]ParityField(nil), contract...),
		UnsupportedProviderFeatures: []string{
			"write_actions", "transfer", "trade", "payment", "raw_credentials",
		},
	}
	if len(transactions) == 0 {
		return report
	}
	for _, field := range contract {
		mapped := false
		for _, transaction := range transactions {
			if transaction.fieldValue(field.Name) {
				mapped = true
				break
			}
		}
		if mapped {
			report.MappedFields++
		} else if field.Required {
			report.MissingRequiredFields = append(report.MissingRequiredFields, field.Name)
		}
	}
	report.RecordsWithRequiredFields = countRequiredCompleteRecords(transactions, contract)
	report.ParityPercent = float64(report.MappedFields) / float64(report.ExpectedFields) * 100
	sort.Strings(report.MissingRequiredFields)
	report.Pass = report.ParityPercent >= ParityThreshold && report.RecordsWithRequiredFields == report.Records
	return report
}

func VerifyFixture(root string) (ParityReport, error) {
	transactions, err := LoadFixture(root)
	if err != nil {
		return ParityReport{}, err
	}
	report := AssessParity(transactions)
	report.FixturePath = FixtureRelativePath
	return report, nil
}

func countRequiredCompleteRecords(transactions []SourceTransaction, contract []ParityField) int {
	count := 0
	for _, transaction := range transactions {
		complete := true
		for _, field := range contract {
			if field.Required && !transaction.fieldValue(field.Name) {
				complete = false
				break
			}
		}
		if complete {
			count++
		}
	}
	return count
}
