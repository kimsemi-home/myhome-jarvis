package localfinanceevidence

import "errors"

func validateCreditTemplateImports(first, second CreditTemplateImport) error {
	if first.TemplateID != "synthetic-card-export" || second.TemplateID != first.TemplateID ||
		first.TemplateVersion != 1 || second.TemplateVersion != 2 || first.FileSHA256 == second.FileSHA256 ||
		first.ProfileSHA256 == second.ProfileSHA256 || first.FingerprintSetHash != second.FingerprintSetHash ||
		!validCreditTemplateHashes(first) || !validCreditTemplateHashes(second) ||
		first.Read != 3 || first.Inserted != 3 || first.Duplicates != 0 || first.SuggestionsRecorded != 3 ||
		second.Read != 3 || second.Inserted != 0 || second.Duplicates != 3 || second.SuggestionsRecorded != 3 ||
		first.NormalizedDebitMinor != 20900 || second.NormalizedDebitMinor != 20900 ||
		first.NormalizedCreditMinor != 2200 || second.NormalizedCreditMinor != 2200 {
		return errors.New("Ledger credit import-template versions are invalid")
	}
	return nil
}

func validCreditTemplateHashes(value CreditTemplateImport) bool {
	return hashPattern.MatchString(value.FileSHA256) && hashPattern.MatchString(value.ProfileSHA256) &&
		hashPattern.MatchString(value.FingerprintSetHash) && value.RunID == "import-"+value.FileSHA256[:16]
}
