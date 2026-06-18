package review

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
	file, err := os.Open(filepath.Join(root, filepath.FromSlash(policy.PrivateReviewQueue)))
	if errors.Is(err, os.ErrNotExist) {
		status.CapacityState, status.ActiveRule = capacityState(policy, status)
		return status, nil
	}
	if err != nil {
		return Status{}, err
	}
	defer file.Close()
	return scanReviewQueue(policy, status, file)
}

func newStatus(policy Policy) Status {
	return Status{
		PolicyPath:             PolicyRelativePath,
		QueuePath:              policy.PrivateReviewQueue,
		MaxOpenReviews:         policy.MaxOpenReviews,
		MaxHighRiskOpenReviews: policy.MaxHighRiskOpenReviews,
		ByRisk:                 map[string]int{},
		ByStatus:               map[string]int{},
		ByReviewerRole:         map[string]int{},
		ByQueueClass:           map[string]int{},
		CheckedAt:              time.Now().UTC().Format(time.RFC3339),
	}
}
