package repofactory

func packetGateReadyCount(gates []GateEvidence) int {
	count := 0
	for _, gate := range gates {
		if gate.State == "ready" {
			count++
		}
	}
	return count
}

func packetBlockingGates(gates []GateEvidence) int {
	count := 0
	for _, gate := range gates {
		if gate.BlocksRepoCreation && gate.State != "ready" {
			count++
		}
	}
	return count
}

func packetMissingEvidence(gates []GateEvidence) []string {
	missing := []string{}
	for _, gate := range gates {
		if gate.Required && gate.State != "ready" {
			missing = append(missing, gate.Key)
		}
	}
	return missing
}
