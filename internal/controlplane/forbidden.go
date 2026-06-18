package controlplane

import "strings"

func rejectForbiddenPolicyText(policy Policy, value string) error {
	lower := strings.ToLower(value)
	for _, field := range policy.ForbiddenPublicFields {
		marker := strings.ToLower(strings.TrimSpace(field))
		if marker != "" && strings.Contains(lower, marker) {
			return errForbiddenPublicMarker()
		}
	}
	return nil
}

func containsForbiddenManifestMarker(policy Policy, line string) bool {
	if rejectSensitiveText(line) != nil {
		return true
	}
	return rejectForbiddenPolicyText(policy, line) != nil
}
