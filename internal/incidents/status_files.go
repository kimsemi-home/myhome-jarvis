package incidents

import (
	"os"
	"path/filepath"
)

func openIncidentLedger(root string, policy Policy) (*os.File, error) {
	return os.Open(filepath.Join(root, filepath.FromSlash(policy.PrivateIncidentLedger)))
}
