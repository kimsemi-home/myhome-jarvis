package codeshape

import "time"

func StatusForRoot(root string) (Status, error) {
	policy, err := ReadPolicy(root)
	if err != nil {
		return Status{}, err
	}
	status := Status{
		PolicyPath:   PolicyRelativePath,
		MaxFileLines: policy.MaxFileLines,
		TopDebt:      []FileFinding{},
		Regressions:  []FileFinding{},
		OK:           true,
		CheckedAt:    time.Now().UTC().Format(time.RFC3339),
	}
	if err := scanRoots(root, policy, &status); err != nil {
		return Status{}, err
	}
	sortFindings(&status)
	status.OK = status.BudgetRegressionCount == 0
	return status, nil
}
