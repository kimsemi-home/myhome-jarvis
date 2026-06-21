package storagearchive

import "testing"

func assertStorageArchiveSourceKeys(t *testing.T, keys []string) {
	t.Helper()
	for _, key := range []string{
		"codex_cost_attribution",
		"monetization",
		"finance_consent",
		"authority_review",
		"authority_approval",
		"external_evidence",
	} {
		if !containsKey(keys, key) {
			t.Fatalf("archive sources = %#v", keys)
		}
	}
}
