package monetization

import (
	"errors"
	"os"
	"path/filepath"
	"time"
)

func StatusForRoot(root string) (Status, error) {
	policy, err := ReadPolicy(root)
	if err != nil {
		return Status{}, err
	}
	status := newStatus(policy)
	file, err := os.Open(filepath.Join(root, filepath.FromSlash(policy.PrivateExperimentLedger)))
	if errors.Is(err, os.ErrNotExist) {
		return status, nil
	}
	if err != nil {
		return Status{}, err
	}
	defer file.Close()
	status.Exists = true
	if err := scanLedger(file, policy, &status); err != nil {
		return Status{}, err
	}
	status.MonetizationDebtCount = debtCount(status)
	return status, nil
}

func newStatus(policy Policy) Status {
	return Status{
		PolicyPath:          PolicyRelativePath,
		LedgerPath:          policy.PrivateExperimentLedger,
		ByState:             map[string]int{},
		ByDecisionKind:      map[string]int{},
		ByReviewStatus:      map[string]int{},
		ByExpectedValueBand: map[string]int{},
		CheckedAt:           time.Now().UTC().Format(time.RFC3339),
	}
}
