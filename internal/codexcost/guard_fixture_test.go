package codexcost

import (
	"encoding/json"
	"fmt"
	"testing"
	"time"

	"github.com/kimsemi-home/myhome-jarvis/internal/codexsustainability"
)

func writeSustainabilityPolicy(t *testing.T, root string) {
	t.Helper()
	policy := codexsustainability.Policy{
		Context:                              "CodexSustainabilityEvidenceLoop",
		Version:                              "v1",
		GeneratedArtifact:                    codexsustainability.PolicyRelativePath,
		PrivateEvidenceLedger:                "data/private/codex-sustainability/evidence.jsonl",
		AppendOnly:                           true,
		PublicStatusRedacted:                 true,
		TrendBaselinesVersioned:              true,
		EvidenceMaxAgeHours:                  168,
		TrendBaselineMaxAgeHours:             168,
		CostPerAcceptedChangeReviewThreshold: 500000,
		RecordKinds: []string{
			"usage_sample", "cycle_sample", "trend_baseline", "feature_proposal",
		},
		Metrics: []string{
			"codex_tokens", "codex_coin", "github_actions_minutes",
			"elapsed_cycle_minutes", "rework_count", "cache_hit_count",
			"cache_miss_count", "validation_failure_count",
			"human_review_debt", "accepted_change_count", "cache_savings_units",
		},
		RequiredFields: []string{"at", "record_kind", "evidence_refs"},
		ProposalRequiredFields: []string{
			"proposal_id", "evidence_refs", "cost_per_accepted_change",
			"median_cycle_minutes", "cache_savings_units", "defect_rework_rate",
			"monetization_ref",
		},
		AllowedEvidencePrefixes: []string{
			"generated/", "docs/", ".github/", "data/private/",
		},
		PublicSummaryFields: []string{
			"record_count", "trend_posture", "sustainability_posture",
			"evidence_freshness", "review_gate_count", "cost_per_accepted_change",
			"cache_savings_units", "rework_count",
		},
		Commands: []string{
			"mhj codex-sustainability status",
			"mhj codex-sustainability record-quality",
			"mhj codex-sustainability record-proposal <json-payload>",
		},
	}
	body, err := json.Marshal(policy)
	if err != nil {
		t.Fatal(err)
	}
	writeFile(t, root, codexsustainability.PolicyRelativePath, string(body)+"\n")
}

func writeSustainableLedger(t *testing.T, root string) {
	t.Helper()
	now := time.Now().UTC().Format(time.RFC3339)
	body := fmt.Sprintf(
		`{"at":%q,"record_kind":"trend_baseline","metric":"elapsed_cycle_minutes","amount":30,"trend_baseline_version":"test","trend_measured_at":%q,"evidence_refs":["docs/codex-sustainability.md"]}`+"\n"+
			`{"at":%q,"record_kind":"cycle_sample","metric":"elapsed_cycle_minutes","amount":10,"evidence_refs":["docs/codex-sustainability.md"]}`+"\n"+
			`{"at":%q,"record_kind":"feature_proposal","proposal_id":"opt","cost_per_accepted_change":100000,"median_cycle_minutes":10,"cache_savings_units":1,"defect_rework_rate":0,"monetization_ref":"KIM-132","evidence_refs":["docs/codex-sustainability.md"]}`+"\n",
		now, now, now, now,
	)
	writeFile(t, root, "data/private/codex-sustainability/evidence.jsonl", body)
}
