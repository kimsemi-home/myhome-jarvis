package repofactory

func packetGates(
	status Status,
	gates []CreationGate,
	safety PublicSafetyEvidence,
) []GateEvidence {
	evidence := make([]GateEvidence, 0, len(gates))
	for _, gate := range gates {
		evidence = append(evidence, GateEvidence{
			Key:                gate.Key,
			Required:           gate.Required,
			BlocksRepoCreation: gate.BlocksRepoCreation,
			EvidenceKind:       gate.Evidence,
			State:              packetGateState(status, gate, safety),
		})
	}
	return evidence
}

func packetGateState(
	status Status,
	gate CreationGate,
	safety PublicSafetyEvidence,
) string {
	switch gate.Key {
	case "generated_ci":
		return readyIf(status.GeneratedCIPresent)
	case "private_data_policy":
		return readyIf(status.PrivateDataPolicyPresent)
	case "bootstrap_checklist":
		return readyIf(status.BootstrapChecklistReady)
	case "authority_review":
		return "blocked_pending_human_review"
	case "public_safety_evidence":
		return safety.EvidenceState
	default:
		if gate.Required {
			return "required_before_creation"
		}
		return "optional"
	}
}

func readyIf(ready bool) string {
	if ready {
		return "ready"
	}
	return "missing"
}
