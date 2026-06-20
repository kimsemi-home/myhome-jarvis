package commandcenter

func hasBlockedGate(gates []GateSummary, wanted string) bool {
	for _, gate := range gates {
		if gate.Key == wanted {
			return true
		}
	}
	return false
}
