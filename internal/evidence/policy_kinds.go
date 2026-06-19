package evidence

import "fmt"

func validateGraphKinds(policy Policy) error {
	nodeKinds := normalizeList(policy.NodeKinds)
	edgeKinds := normalizeList(policy.EdgeKinds)
	for _, kind := range []string{"learning_observation", "evidence_artifact"} {
		if !contains(nodeKinds, kind) {
			return fmt.Errorf("evidence graph policy must include learning and artifact nodes")
		}
	}
	if !contains(edgeKinds, "supports") {
		return fmt.Errorf("evidence graph policy must include supports edges")
	}
	return nil
}

func validatePolicySummary(policy Policy) error {
	summaryFields := normalizeList(policy.PublicSummaryFields)
	for _, field := range []string{
		"node_count",
		"edge_count",
		"by_node_kind",
		"by_edge_kind",
		"checked_at",
	} {
		if !contains(summaryFields, field) {
			return fmt.Errorf("evidence graph public summary missing %q", field)
		}
	}
	if !contains(normalizeList(policy.AllowedEvidencePrefixes), "data/private/") {
		return fmt.Errorf("evidence graph evidence refs must allow data/private")
	}
	if !contains(policy.Commands, "mhj evidence status") {
		return fmt.Errorf("evidence graph status command is missing")
	}
	return nil
}
