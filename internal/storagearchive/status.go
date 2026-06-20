package storagearchive

import (
	"path/filepath"

	"github.com/kimsemi-home/myhome-jarvis/internal/domain"
)

const PolicyRelativePath = "generated/storage.generated.json"

func StatusForRoot(root string) (Status, error) {
	policy, err := domain.ReadStoragePolicy(
		filepath.Join(root, filepath.FromSlash(PolicyRelativePath)),
	)
	if err != nil {
		return Status{}, err
	}
	if err := ValidatePolicy(policy); err != nil {
		return Status{}, err
	}
	manifest, err := readManifestSummary(root, policy)
	if err != nil {
		return Status{}, err
	}
	return statusFromPolicy(policy, manifest), nil
}
