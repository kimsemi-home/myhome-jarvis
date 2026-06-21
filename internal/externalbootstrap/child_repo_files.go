package externalbootstrap

import (
	"encoding/json"
	"os"
	"path/filepath"
)

func checkChildRepoFiles(root string, packet Packet, status *ChildRepoStatus) {
	for _, rel := range childRequiredFiles(packet) {
		info, err := os.Stat(filepath.Join(root, filepath.FromSlash(rel)))
		if err != nil || info.IsDir() || info.Size() == 0 {
			status.addFinding(rel, "missing_file", "required child repo file is missing")
		}
	}
}

func checkChildRepoPrivateData(root string, status *ChildRepoStatus) {
	if _, err := os.Stat(filepath.Join(root, "data", "private")); err == nil {
		status.PrivateDataAbsent = false
		status.addFinding("data/private", "public_safety_private_data",
			"public child repo must not contain private data")
		return
	}
	status.PrivateDataAbsent = true
}

func readChildJSON[T any](root string, rel string) (T, bool) {
	var value T
	body, err := os.ReadFile(filepath.Join(root, filepath.FromSlash(rel)))
	if err != nil {
		return value, false
	}
	if err := json.Unmarshal(body, &value); err != nil {
		return value, false
	}
	return value, true
}
