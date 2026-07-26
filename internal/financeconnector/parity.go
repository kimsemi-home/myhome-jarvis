package financeconnector

import "sort"

func AssessParity(transactions []SourceTransaction) ParityReport {
	report := ParityReport{
		Provider:                    ProviderTossMyData,
		Records:                     len(transactions),
		ExpectedFields:              len(CanonicalFieldContract),
		ThresholdPercent:            ParityThreshold,
		FieldContract:               append([]ParityField(nil), CanonicalFieldContract...),
		UnsupportedProviderFeatures: []string{"write_actions", "transfer", "trade", "payment", "raw_credentials"},
	}
	if len(transactions) == 0 {
		return report
	}

	for _, field := range CanonicalFieldContract {
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
	report.RecordsWithRequiredFields = countRequiredCompleteRecords(transactions)
	report.ParityPercent = float64(report.MappedFields) / float64(report.ExpectedFields) * 100
	sort.Strings(report.MissingRequiredFields)
	report.Pass = report.ParityPercent >= ParityThreshold &&
		report.RecordsWithRequiredFields == report.Records
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

func countRequiredCompleteRecords(transactions []SourceTransaction) int {
	count := 0
	for _, transaction := range transactions {
		complete := true
		for _, field := range CanonicalFieldContract {
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
