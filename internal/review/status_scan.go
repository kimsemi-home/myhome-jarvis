package review

import (
	"bufio"
	"encoding/json"
	"errors"
	"io"
	"strings"
)

func scanReviewQueue(policy Policy, status Status, reader io.Reader) (Status, error) {
	status.Exists = true
	scanner := bufio.NewScanner(reader)
	scanner.Buffer(make([]byte, 0, 64*1024), 1024*1024)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}
		var item Review
		if err := json.Unmarshal([]byte(line), &item); err != nil {
			status.InvalidReviewCount++
			continue
		}
		normalized, err := normalizeReview(policy, item)
		if err != nil {
			countRejectedReview(&status, err)
			continue
		}
		recordReview(&status, normalized)
	}
	if err := scanner.Err(); err != nil {
		return Status{}, err
	}
	return finalizeStatus(policy, status), nil
}

func countRejectedReview(status *Status, err error) {
	switch {
	case errors.Is(err, errMissingEvidenceRef):
		status.MissingEvidenceCount++
	case strings.Contains(err.Error(), "reviewer"):
		status.MissingReviewerCount++
	default:
		status.InvalidReviewCount++
	}
}

func recordReview(status *Status, review Review) {
	status.Count++
	status.ByRisk[review.Risk]++
	status.ByStatus[review.Status]++
	status.ByQueueClass[review.QueueClass]++
	if review.ReviewerRole != "" {
		status.ByReviewerRole[review.ReviewerRole]++
	}
	if review.BackupAvailable {
		status.BackupAvailableCount++
	}
	if reviewOpen(review.Status) {
		status.OpenCount++
		if review.Risk == "high" {
			status.HighRiskOpenCount++
		}
	}
	status.LastObservedAt = laterRFC3339(status.LastObservedAt, review.At)
}
