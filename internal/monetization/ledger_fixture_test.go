package monetization

func ledgerFixture() string {
	return `{"at":"2026-06-19T00:00:00Z","experiment_id":"shorts-001","hypothesis_key":"shorts_factory","state":"review_required","decision_kind":"scale_requested","review_status":"requested","expected_value_band":"high","cost_estimate_units":420,"cost_unit_kind":"codex_tokens","evidence_refs":["generated/assistant_vision.generated.json"],"private_revenue_notes":"private revenue note"}` + "\n" +
		`{"at":"2026-06-19T01:00:00Z","experiment_id":"shorts-001","hypothesis_key":"shorts_factory","state":"running","decision_kind":"hypothesis_created","review_status":"approved","expected_value_band":"medium","cost_estimate_units":350,"cost_unit_kind":"codex_tokens","evidence_refs":["docs/assistant-vision.md"]}` + "\n" +
		`{"at":"2026-06-19T02:00:00Z","experiment_id":"blog-001","hypothesis_key":"affiliate_blog","state":"backlog","decision_kind":"hypothesis_created","review_status":"not_required","expected_value_band":"unknown","cost_estimate_units":10,"cost_unit_kind":"codex_coin","evidence_refs":["docs/assistant-vision.md"]}` + "\n" +
		`{"at":"2026-06-19T03:00:00Z","experiment_id":"bad-001","hypothesis_key":"missing_evidence","state":"backlog","decision_kind":"hypothesis_created","review_status":"not_required","expected_value_band":"low","cost_estimate_units":10,"cost_unit_kind":"codex_coin","evidence_refs":[]}` + "\n" +
		`{"at":"2026-06-19T04:00:00Z","experiment_id":"bad-002","hypothesis_key":"missing_cost","state":"backlog","decision_kind":"hypothesis_created","review_status":"not_required","expected_value_band":"low","cost_unit_kind":"codex_coin","evidence_refs":["docs/assistant-vision.md"]}` + "\n" +
		`{"at":"2026-06-19T05:00:00Z","experiment_id":"bad-003","hypothesis_key":"bad_state","state":"launched","decision_kind":"hypothesis_created","review_status":"not_required","expected_value_band":"low","cost_estimate_units":1,"cost_unit_kind":"codex_coin","evidence_refs":["docs/assistant-vision.md"]}` + "\n"
}
