package codexcost

import "fmt"

func normalizeGuardRequest(policy Policy, request GuardRequest) (GuardRequest, error) {
	guard := GuardRequest{
		Scope:            normalizeToken(request.Scope),
		UnitKind:         normalizeToken(request.UnitKind),
		EstimatedUnits:   request.EstimatedUnits,
		EstimatedMinutes: request.EstimatedMinutes,
		EvidenceRefs:     normalizeRefs(request.EvidenceRefs),
	}
	if guard.EstimatedUnits <= 0 || guard.EstimatedMinutes <= 0 {
		return GuardRequest{}, fmt.Errorf("codex cost guard estimates must be positive")
	}
	if !contains(normalizeList(policy.LoopScopes), guard.Scope) {
		return GuardRequest{}, fmt.Errorf("codex cost guard scope %q is not allowed", guard.Scope)
	}
	if !contains(normalizeList(policy.UnitKinds), guard.UnitKind) {
		return GuardRequest{}, fmt.Errorf("codex cost guard unit kind %q is not allowed", guard.UnitKind)
	}
	if len(guard.EvidenceRefs) == 0 {
		return GuardRequest{}, errMissingEvidenceRef
	}
	for _, ref := range guard.EvidenceRefs {
		if err := validateRef(policy, ref); err != nil {
			return GuardRequest{}, err
		}
	}
	return guard, nil
}
