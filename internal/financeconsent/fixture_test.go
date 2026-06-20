package financeconsent

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"
)

func testPolicy() Policy {
	return Policy{
		Context:                       "HouseholdFinanceConsentLedger",
		PrivateConsentLedger:          "data/private/finance/consent.jsonl",
		AppendOnly:                    true,
		PublicStatusRedacted:          true,
		FinanceMode:                   "read_only_review_only",
		ReadOnly:                      true,
		ReviewOnly:                    true,
		FixtureOnlyUntilConsent:       true,
		RealConnectorRequiresConsent:  true,
		SpouseScopeRequiresConsent:    true,
		HouseholdScopeRequiresConsent: true,
		ConsentKinds:                  requiredConsentKinds,
		ConsentStatuses:               requiredConsentStatuses,
		ReviewStatuses:                requiredReviewStatuses,
		AuthorityProfiles:             []string{"finance_review_only"},
		RequiredFields:                requiredFields,
		AllowedEvidencePrefixes:       []string{"data/private/", "generated/", "docs/"},
		PublicSummaryFields:           requiredSummaryFields,
		Commands: []string{
			"mhj finance-consent status",
			"mhj finance-consent record <json-payload>",
		},
	}
}

func writePolicy(t *testing.T, root string, policy Policy) {
	t.Helper()
	body, err := json.Marshal(policy)
	if err != nil {
		t.Fatal(err)
	}
	writeFile(t, root, PolicyRelativePath, string(body)+"\n")
}

func writeFile(t *testing.T, root string, rel string, body string) {
	t.Helper()
	path := filepath.Join(root, filepath.FromSlash(rel))
	if err := os.MkdirAll(filepath.Dir(path), 0o700); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(path, []byte(body), 0o600); err != nil {
		t.Fatal(err)
	}
}
