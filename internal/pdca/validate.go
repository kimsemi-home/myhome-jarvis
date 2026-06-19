package pdca

import (
	"fmt"
	"strings"
)

func validatePolicy(policy Policy) error {
	if policy.Context != "AgentCluster" {
		return fmt.Errorf("PDCA policy context = %q", policy.Context)
	}
	if !strings.HasPrefix(policy.PrivateCycleLedger, "data/private/") ||
		!strings.HasSuffix(policy.PrivateCycleLedger, ".jsonl") {
		return fmt.Errorf("PDCA cycle ledger must stay in data/private JSONL")
	}
	if !policy.AppendOnly || !policy.PublicStatusRedacted || policy.RawCyclePublicAllowed {
		return fmt.Errorf("PDCA policy must be private append-only with redacted public status")
	}
	if err := validateSteps(policy.Steps); err != nil {
		return err
	}
	if err := requireAll("required field", policy.RequiredFields, requiredFields); err != nil {
		return err
	}
	if err := requireAll("status", policy.AllowedStatuses, requiredStatuses); err != nil {
		return err
	}
	if err := requireAll("evidence source", policy.EvidenceSources, requiredSources); err != nil {
		return err
	}
	if err := requireAll("public summary", policy.PublicSummaryFields, requiredSummary); err != nil {
		return err
	}
	if !contains(policy.Commands, "mhj pdca status") {
		return fmt.Errorf("PDCA status command is missing")
	}
	return nil
}

func validateSteps(steps []Step) error {
	if len(steps) != len(requiredSteps) {
		return fmt.Errorf("PDCA must declare four ordered steps")
	}
	for i, id := range requiredSteps {
		step := steps[i]
		if step.ID != id || step.Role == "" || step.Artifact == "" || step.Command == "" {
			return fmt.Errorf("PDCA step %q is invalid", id)
		}
	}
	return nil
}
