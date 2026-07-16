package localfinanceevidence

import (
	"encoding/json"
	"errors"
	"slices"
)

func validateCreditBatch(value CreditBatchPreview, linked CreditImportPreview) error {
	if value.SchemaVersion != "myhome.ledger-import-batch-preview/v1" || value.ExecutionMode != "batch_preview_only" ||
		value.CredentialsRead || value.ExternalNetwork || value.ExternalWrites ||
		value.ManifestSHA256 != "f378379d8cf2aadf06f3a9e984796d634e0e27760b71372d3a8e96afced90efa" ||
		value.StatementCount != 2 || value.ReadyCount != 2 || !value.AllReady ||
		value.RawFileNamesReported || value.RawRowsReported || len(value.Statements) != 2 ||
		!slices.Equal(value.Checks, requiredCreditBatchChecks()) {
		return errors.New("Ledger batch preview boundary is invalid")
	}
	expectedNames := []string{
		"d8852d20dfb991da62854549f8ebe761d1cba30a5fc62dc658b18684f3432079",
		"e817e6bee3b048c5ce3abde65a8a1576681ebd6c86e306382145274f15874df6",
	}
	sources, identities, linkedPreview := map[string]bool{}, map[string]bool{}, false
	for index, statement := range value.Statements {
		preview := statement.Preview
		if statement.SourceNameSHA256 != expectedNames[index] || validateCreditPreviewContract(preview) != nil ||
			sources[preview.SourceSHA256] || identities[preview.FingerprintSetHash] {
			return errors.New("Ledger batch statement proof is invalid")
		}
		sources[preview.SourceSHA256], identities[preview.FingerprintSetHash] = true, true
		if preview.TemplateID == linked.TemplateID && preview.TemplateVersion == linked.TemplateVersion {
			linkedPreview = preview.PreviewHash == linked.PreviewHash && preview.SourceSHA256 == linked.SourceSHA256 &&
				preview.ProfileSHA256 == linked.ProfileSHA256
		}
	}
	if !linkedPreview || value.PreviewSetHash != creditBatchPreviewSetHash(value.Statements) {
		return errors.New("Ledger batch preview set is invalid")
	}
	copy := value
	copy.BatchHash = ""
	body, err := json.Marshal(copy)
	if err != nil || !hashPattern.MatchString(value.BatchHash) || value.BatchHash != digest(string(body)) {
		return errors.New("Ledger batch preview hash is invalid")
	}
	return nil
}

func requiredCreditBatchChecks() []string {
	return []string{
		"bounded-manifest", "duplicate-content-denied", "duplicate-normalized-set-denied",
		"filename-hash-only", "no-database-write", "previews-hash-bound", "root-containment-enforced",
	}
}
