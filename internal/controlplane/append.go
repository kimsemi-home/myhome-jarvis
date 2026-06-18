package controlplane

import (
	"encoding/json"
	"os"
	"path/filepath"
)

func AppendManifest(root string, request ManifestRequest) (RecordResult, error) {
	policy, err := ReadPolicy(root)
	if err != nil {
		return RecordResult{}, err
	}
	manifest, err := normalizeManifest(policy, request)
	if err != nil {
		return RecordResult{}, err
	}
	path := filepath.Join(root, filepath.FromSlash(policy.PrivateManifestLedger))
	if err := os.MkdirAll(filepath.Dir(path), 0o700); err != nil {
		return RecordResult{}, err
	}
	if err := appendManifestLine(path, manifest); err != nil {
		return RecordResult{}, err
	}
	return RecordResult{
		ID:           manifest.ID,
		ManifestPath: policy.PrivateManifestLedger,
		DecisionKind: manifest.DecisionKind,
		LeaseStatus:  manifest.LeaseStatus,
		RecordedAt:   manifest.At,
	}, nil
}

func appendManifestLine(path string, manifest Manifest) error {
	file, err := os.OpenFile(path, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0o600)
	if err != nil {
		return err
	}
	defer file.Close()
	data, err := json.Marshal(manifest)
	if err != nil {
		return err
	}
	_, err = file.Write(append(data, '\n'))
	return err
}
