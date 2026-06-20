package repofactory

func contains(values []string, value string) bool {
	for _, item := range values {
		if item == value {
			return true
		}
	}
	return false
}

func containsAll(values []string, required []string) []string {
	missing := []string{}
	for _, item := range required {
		if !contains(values, item) {
			missing = append(missing, item)
		}
	}
	return missing
}

func templateRoles(files []TemplateFile) []string {
	roles := make([]string, 0, len(files))
	for _, file := range files {
		roles = append(roles, file.Role)
	}
	return roles
}

func gateKeys(gates []CreationGate) []string {
	keys := make([]string, 0, len(gates))
	for _, gate := range gates {
		keys = append(keys, gate.Key)
	}
	return keys
}
