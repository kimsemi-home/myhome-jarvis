package cicache

func publicSafetyNonSkippable(graph graphFile) bool {
	for _, unit := range graph.Units {
		if unit.ID == "public-safety" {
			return unit.Cache == "" && len(unit.HashInputs) == 0
		}
	}
	return false
}
