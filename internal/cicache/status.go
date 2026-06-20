package cicache

import "strings"

const (
	graphPath    = "generated/verification_graph.generated.json"
	workflowPath = ".github/workflows/quality.yml"
)

func StatusForRoot(root string) (Status, error) {
	graph, workflow, err := readInputs(root)
	if err != nil {
		return Status{}, err
	}
	status := Status{
		GraphPath:                 graphPath,
		WorkflowPath:              workflowPath,
		OK:                        true,
		PublicSafe:                true,
		GeneratedArtifactCount:    len(graph.GeneratedArtifacts),
		GeneratedCoverageRequired: true,
		RawEvidencePublicAllowed:  false,
		PrivatePayloadsAllowed:    false,
		PushOnlyCacheSaveRequired: true,
		CacheHitSkipsVerification: true,
		CacheMissRunsVerification: true,
		PublicSafetyNonSkippable:  publicSafetyNonSkippable(graph),
	}
	for _, unit := range graph.Units {
		unitStatus := inspectUnit(unit, graph.GeneratedArtifacts, workflow)
		status.Units = append(status.Units, unitStatus)
		applyUnitStatus(&status, unitStatus)
	}
	status.GeneratedCoverageOK = ssotGeneratedCoverageOK(status.Units)
	status.OK = status.PublicSafetyNonSkippable && status.GeneratedCoverageOK &&
		status.InvalidCachedUnitCount == 0 && status.WorkflowContractIssueCount == 0
	return status, nil
}

func applyUnitStatus(status *Status, unit UnitStatus) {
	if unit.CacheKey == "" {
		status.UncachedUnitCount++
		return
	}
	status.CachedUnitCount++
	if !unit.Valid {
		status.InvalidCachedUnitCount++
	}
	if !unit.CacheRestoreConfigured || !unit.CacheHitSkipsVerification ||
		!unit.CacheMissRunsVerification || !unit.PushOnlyCacheSaveConfigured {
		status.WorkflowContractIssueCount++
	}
}

func ssotGeneratedCoverageOK(units []UnitStatus) bool {
	for _, unit := range units {
		if strings.EqualFold(unit.ID, "ssot") {
			return unit.GeneratedArtifactCoverageOK
		}
	}
	return false
}
