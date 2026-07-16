package localfinanceevidence

import (
	"encoding/json"
	"testing"
)

func resealCreditTemplateAndReport(t *testing.T, report *CreditReport) {
	t.Helper()
	report.ImportTemplate.ReportHash = ""
	unsignedTemplate, err := json.Marshal(report.ImportTemplate)
	if err != nil {
		t.Fatal(err)
	}
	report.ImportTemplate.ReportHash = digest(string(unsignedTemplate))
	report.ReportHash = ""
	unsignedReport, err := json.Marshal(report)
	if err != nil {
		t.Fatal(err)
	}
	report.ReportHash = digest(string(unsignedReport))
}
