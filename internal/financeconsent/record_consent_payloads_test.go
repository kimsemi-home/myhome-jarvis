package financeconsent

func activeConsentPayloads() []string {
	return []string{
		`{"consent_kind":"finance_connector","subject_scope":"owner_read_only","status":"granted","review_status":"approved","authority_profile":"finance_review_only","evidence_refs":["docs/finance-consent.md"]}`,
		`{"consent_kind":"spouse_scope","subject_scope":"spouse_read_only","status":"granted","review_status":"approved","authority_profile":"finance_review_only","evidence_refs":["docs/finance-consent.md"]}`,
		`{"consent_kind":"household_scope","subject_scope":"household_read_only","status":"granted","review_status":"approved","authority_profile":"finance_review_only","evidence_refs":["docs/finance-consent.md"]}`,
	}
}

func mustFinanceConsentJSON(t testingT, value any) []byte {
	t.Helper()
	body, err := jsonMarshal(value)
	if err != nil {
		t.Fatal(err)
	}
	return body
}
