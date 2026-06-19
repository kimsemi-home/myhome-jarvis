package evidencequality

var requiredQualityLevels = []string{"high", "medium", "low", "blocked"}

var requiredMappingLevels = []string{"high", "medium", "low", "unknown"}

var requiredPurposes = []string{
	"root_cause",
	"confidence_assessment",
	"incident_review",
	"release_gate",
	"conformance",
	"revalidation",
}

var requiredReassessmentReasons = []string{
	"age",
	"schema_version_change",
	"ontology_version_change",
	"counter_evidence",
	"security_incident",
	"quarantine",
	"translation_loss",
}

var requiredSnapshotFields = []string{
	"at",
	"evidence_ref",
	"purpose",
	"quality_level",
	"schema_version",
	"ontology_version",
	"mapping_confidence",
	"assessed_by",
	"reassessment_reasons",
}

var requiredSummaryFields = []string{
	"snapshot_count",
	"invalid_snapshot_count",
	"reassessment_debt_count",
	"missing_evidence_count",
	"stale_snapshot_count",
	"low_quality_count",
	"blocked_quality_count",
	"mapping_drift_count",
	"by_quality_level",
	"by_mapping_confidence",
	"checked_at",
}
