package incidents

func requiredIncidentFields() []string {
	return []string{"at", "kind", "stage", "status", "owner_role", "evidence_refs"}
}

func requiredSummaryFields() []string {
	return []string{
		"count",
		"open_count",
		"incident_debt_count",
		"missing_owner_count",
		"missing_evidence_ref_count",
		"stale_quarantine_count",
		"by_stage",
		"by_owner_role",
		"checked_at",
	}
}
