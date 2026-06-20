package mergeevidence

func contains(values []string, value string) bool {
	for _, item := range values {
		if item == value {
			return true
		}
	}
	return false
}

func missingStrings(values []string, required []string) []string {
	var missing []string
	for _, item := range required {
		if !contains(values, item) {
			missing = append(missing, item)
		}
	}
	return missing
}

func gateKeys(gates []Gate) []string {
	keys := make([]string, 0, len(gates))
	for _, gate := range gates {
		keys = append(keys, gate.Key)
	}
	return keys
}
