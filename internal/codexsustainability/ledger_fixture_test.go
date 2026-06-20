package codexsustainability

func debtLedgerFixture() string {
	return `{"at":"2026-06-19T00:00:00Z","record_kind":"trend_baseline","metric":"elapsed_cycle_minutes","amount":30,"trend_baseline_version":"trend-2026w25","trend_measured_at":"2026-06-19T00:00:00Z","evidence_refs":["docs/assistant-vision.md"]}` + "\n" +
		`{"at":"2026-06-19T01:00:00Z","record_kind":"cycle_sample","metric":"elapsed_cycle_minutes","amount":45,"evidence_refs":["generated/verification_graph.generated.json"]}` + "\n" +
		`{"at":"2026-06-19T02:00:00Z","record_kind":"usage_sample","metric":"codex_tokens","amount":700000,"evidence_refs":["generated/codex_cost.generated.json"],"raw_prompt":"private prompt"}` + "\n" +
		`{"at":"2026-06-19T03:00:00Z","record_kind":"usage_sample","metric":"accepted_change_count","amount":1,"evidence_refs":["docs/codex-cost-governor.md"]}` + "\n" +
		`{"at":"2026-06-19T04:00:00Z","record_kind":"usage_sample","metric":"validation_failure_count","amount":1,"evidence_refs":[".github/workflows/quality.yml"]}` + "\n" +
		`{"at":"2026-06-19T05:00:00Z","record_kind":"feature_proposal","proposal_id":"opt-1","cost_per_accepted_change":1,"median_cycle_minutes":1,"cache_savings_units":0,"defect_rework_rate":0,"monetization_ref":"KIM-140","evidence_refs":[]}` + "\n"
}
