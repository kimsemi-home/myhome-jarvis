package localfinanceevidence

import (
	"errors"
	"regexp"
	"sort"
)

var hashPattern = regexp.MustCompile(`^[a-f0-9]{64}$`)

var requiredCapabilities = map[string]string{
	"ledger":    "monthly-credit-summary",
	"portfolio": "read-only-holdings-snapshot",
	"revenue":   "youtube-monthly-revenue",
	"shorts":    "youtube-private-upload-plan",
}

func Validate(manifest Manifest) error {
	if manifest.SchemaVersion != ManifestSchema ||
		!regexp.MustCompile(`^\d{4}-\d{2}$`).MatchString(manifest.Month) ||
		manifest.ExternalWritesEnabled || len(manifest.Receipts) != len(requiredCapabilities) {
		return errors.New("local finance evidence manifest is invalid")
	}
	seen := map[string]bool{}
	for _, receipt := range manifest.Receipts {
		if err := validateReceipt(receipt); err != nil || seen[receipt.Component] {
			return errors.New("local finance evidence receipt is invalid")
		}
		seen[receipt.Component] = true
	}
	if len(seen) != len(requiredCapabilities) || manifest.AggregateHash != aggregateHash(manifest) {
		return errors.New("local finance aggregate hash is invalid")
	}
	return nil
}

func validateReceipt(receipt Receipt) error {
	expected, ok := requiredCapabilities[receipt.Component]
	checks := append([]string{}, receipt.Checks...)
	sort.Strings(checks)
	if !ok || receipt.SchemaVersion != ReceiptSchema || receipt.Capability != expected ||
		(receipt.ExecutionMode != "fixture_only" && receipt.ExecutionMode != "plan_only") ||
		receipt.ExternalWritesEnabled || !hashPattern.MatchString(receipt.ArtifactHash) ||
		len(checks) == 0 || receipt.ReceiptHash != receiptHash(receipt) {
		return errors.New("indirect evidence receipt is invalid")
	}
	return nil
}
