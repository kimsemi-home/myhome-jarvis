package localfinanceevidence

import (
	"encoding/json"
	"errors"
)

const creditTemplateProofSchema = "myhome.ledger-credit-import-template-rehearsal/v1"

func validateCreditTemplateReport(value CreditTemplateReport, month string) error {
	if value.SchemaVersion != creditTemplateProofSchema || value.ExecutionMode != "fixture_only" ||
		value.CredentialsRead || value.ExternalNetwork || value.ExternalWrites || value.Month != month {
		return errors.New("Ledger credit import-template boundary is invalid")
	}
	if err := validateCreditTemplateImports(value.FirstImport, value.SecondImport); err != nil {
		return err
	}
	if err := validateCreditTemplateParts(value); err != nil {
		return err
	}
	copy := value
	copy.ReportHash = ""
	body, err := json.Marshal(copy)
	if err != nil || !hashPattern.MatchString(value.ReportHash) || value.ReportHash != digest(string(body)) {
		return errors.New("Ledger credit import-template report hash is invalid")
	}
	return nil
}
