package review

import (
	"encoding/json"
	"testing"
)

const highRiskReviewJSON = `{"id":"rev_1","at":"2026-06-18T00:00:00Z","item_key":"major_ontology_change","queue_class":"major_ontology_change","risk":"high","status":"in_review","requester_role":"producer","reviewer_role":"governance_steward","backup_available":true,"evidence_refs":["generated/authority.generated.json"]}`

const missingReviewerJSON = `{"id":"rev_1","at":"2026-06-18T00:00:00Z","item_key":"security_boundary_change","queue_class":"authority_boundary_change","risk":"high","status":"requested","requester_role":"producer","reviewer_role":"","backup_available":false,"evidence_refs":["data/private/review/raw-notes.json"],"raw_review_notes":"private reviewer identity"}`

const missingEvidenceJSON = `{"id":"rev_2","at":"2026-06-18T00:01:00Z","item_key":"ssot_defect","queue_class":"ssot_defect","risk":"medium","status":"assigned","requester_role":"producer","reviewer_role":"independent_reviewer","backup_available":false,"evidence_refs":[],"raw_rationale":"private rationale"}`

func testPolicy() Policy {
	return Policy{
		Context:                "AgentCluster",
		Version:                "v1",
		GeneratedArtifact:      "generated/review.generated.json",
		PrivateReviewQueue:     "data/private/review/queue.jsonl",
		AppendOnly:             true,
		PublicStatusRedacted:   true,
		MaxOpenReviews:         5,
		MaxHighRiskOpenReviews: 0,
		MinBackupReviewers:     1,
		AllowedRisks:           requiredRisks,
		QueueClasses:           requiredQueueClasses,
		PriorityOrder:          requiredQueueClasses,
		AllowedStatuses:        requiredStatuses,
		RequesterRoles:         requiredRequesterRoles,
		ReviewerRoles:          requiredReviewerRoles,
		RequiredFields:         requiredReviewFields,
		PublicSummaryFields:    requiredSummaryFields,
		Commands:               []string{"mhj review status"},
		OverloadPolicy:         testOverloadPolicy(),
		AllowedEvidencePrefixes: []string{
			"data/private/", "generated/", "docs/", "cmd/", "internal/",
			"apps/flutter/", "lisp/", "crates/", "fixtures/", "harness/", ".github/",
		},
	}
}

func testOverloadPolicy() []OverloadRule {
	return []OverloadRule{
		{Key: "low_risk_only", AllowedWhenOverloaded: true},
		{Key: "deterministic_verification", AllowedWhenOverloaded: true},
		{Key: "evidence_collection", AllowedWhenOverloaded: true},
		{Key: "incident_response", AllowedWhenOverloaded: true},
		{Key: "revalidation", AllowedWhenOverloaded: true},
		{Key: "high_risk_change"},
		{Key: "major_ontology_change"},
		{Key: "security_boundary_change"},
		{Key: "production_change"},
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
