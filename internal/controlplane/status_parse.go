package controlplane

import "encoding/json"

func parseManifestLine(policy Policy, status *Status, line string) (Manifest, bool) {
	if containsForbiddenManifestMarker(policy, line) {
		status.InvalidManifestCount++
		return Manifest{}, false
	}
	var manifest Manifest
	if err := json.Unmarshal([]byte(line), &manifest); err != nil {
		status.InvalidManifestCount++
		return Manifest{}, false
	}
	if err := validateManifest(policy, manifest); err != nil {
		status.InvalidManifestCount++
		return Manifest{}, false
	}
	return manifest, true
}
