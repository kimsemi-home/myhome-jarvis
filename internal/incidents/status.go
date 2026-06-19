package incidents

import (
	"errors"
	"os"
	"time"
)

const PolicyRelativePath = "generated/incidents.generated.json"

func StatusForRoot(root string) (Status, error) {
	policy, err := ReadPolicy(root)
	if err != nil {
		return Status{}, err
	}
	checkedAt := time.Now().UTC()
	status := newStatus(policy, checkedAt)
	file, err := openIncidentLedger(root, policy)
	if errors.Is(err, os.ErrNotExist) {
		return status, nil
	}
	if err != nil {
		return Status{}, err
	}
	defer file.Close()

	status.Exists = true
	if err := scanIncidentLedger(file, policy, checkedAt, &status); err != nil {
		return Status{}, err
	}
	status.IncidentDebtCount = incidentDebt(status)
	return status, nil
}
