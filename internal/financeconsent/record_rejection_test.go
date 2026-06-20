package financeconsent

import "testing"

func TestRecordConsentRejectsInvalidPayloads(t *testing.T) {
	root := t.TempDir()
	writePolicy(t, root, testPolicy())
	for _, payload := range []string{
		`{"consent_kind":"bad","subject_scope":"owner","status":"granted","review_status":"approved","authority_profile":"finance_review_only","evidence_refs":["docs/finance-consent.md"]}`,
		`{"consent_kind":"finance_connector","subject_scope":"owner","status":"requested","review_status":"requested","authority_profile":"finance_review_only","evidence_refs":["docs/finance-consent.md"]}`,
		`{"consent_kind":"finance_connector","subject_scope":"person_name","status":"granted","review_status":"approved","authority_profile":"finance_review_only","evidence_refs":["docs/finance-consent.md"]}`,
		`{"consent_kind":"finance_connector","subject_scope":"owner","status":"granted","review_status":"approved","authority_profile":"finance_review_only","evidence_refs":["/tmp/evidence.md"]}`,
		`{"consent_kind":"finance_connector","subject_scope":"owner","status":"granted","review_status":"approved","authority_profile":"finance_review_only","private_notes":"private","evidence_refs":["docs/finance-consent.md"]}`,
	} {
		if _, err := RecordConsent(root, []byte(payload)); err == nil {
			t.Fatalf("expected error for %s", payload)
		}
	}
}
