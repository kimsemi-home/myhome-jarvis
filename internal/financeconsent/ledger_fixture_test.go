package financeconsent

func activeLedgerFixture() string {
	return `{"at":"2026-06-19T00:00:00Z","consent_kind":"finance_connector","subject_scope":"user","status":"granted","review_status":"approved","authority_profile":"finance_review_only","evidence_refs":["data/private/review/finance-connector.md"],"private_notes":"private consent note"}` + "\n" +
		`{"at":"2026-06-19T00:01:00Z","consent_kind":"spouse_scope","subject_scope":"spouse","status":"granted","review_status":"approved","authority_profile":"finance_review_only","evidence_refs":["data/private/review/spouse-scope.md"]}` + "\n" +
		`{"at":"2026-06-19T00:02:00Z","consent_kind":"household_scope","subject_scope":"household","status":"granted","review_status":"approved","authority_profile":"finance_review_only","evidence_refs":["data/private/review/household-scope.md"]}` + "\n"
}

func debtLedgerFixture() string {
	return `{"at":"2026-06-19T00:00:00Z","consent_kind":"finance_connector","subject_scope":"user","status":"requested","review_status":"requested","authority_profile":"finance_review_only","evidence_refs":[]}` + "\n" +
		`{"at":"2026-06-19T00:01:00Z","consent_kind":"spouse_scope","subject_scope":"spouse","status":"revoked","review_status":"approved","authority_profile":"finance_review_only","evidence_refs":["data/private/review/spouse-scope.md"]}` + "\n" +
		`{"at":"2026-06-19T00:02:00Z","consent_kind":"invalid","subject_scope":"household","status":"granted","review_status":"approved","authority_profile":"finance_review_only","evidence_refs":["data/private/review/household-scope.md"]}` + "\n"
}
