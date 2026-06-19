package learning

import (
	"fmt"
	"strings"
)

func validatePolicy(policy Policy) error {
	if policy.Context != "AgentCluster" {
		return fmt.Errorf("learning policy context = %q", policy.Context)
	}
	if !strings.HasPrefix(policy.PrivateLedger, "data/private/") ||
		!strings.HasSuffix(policy.PrivateLedger, ".jsonl") {
		return fmt.Errorf("learning ledger must stay in data/private JSONL")
	}
	if !policy.AppendOnly || !policy.PrivateJournalRequired ||
		!policy.PublicStatusRedacted || policy.RawObservationPublicAllowed {
		return fmt.Errorf("learning policy must be private append-only with redacted public status")
	}
	for _, required := range []string{
		"kind",
		"source",
		"summary",
		"evidence_refs",
		"owner",
		"next_action",
	} {
		if !contains(normalizeList(policy.RequiredFields), required) {
			return fmt.Errorf("learning policy missing required field %q", required)
		}
	}
	if !contains(normalizeList(policy.AllowedKinds), "loop_gap") ||
		!contains(normalizeList(policy.AllowedKinds), "evidence_debt") {
		return fmt.Errorf("learning policy must track loop_gap and evidence_debt")
	}
	if !contains(normalizeList(policy.Commands), "mhj learning status") ||
		!contains(normalizeList(policy.Commands), "mhj learning record <json-payload>") {
		return fmt.Errorf("learning policy commands are incomplete")
	}
	return nil
}
