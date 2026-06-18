package review

import (
	"fmt"
	"time"
)

func normalizeReview(policy Policy, item Review) (Review, error) {
	normalized := Review{
		ID:              publicText(item.ID),
		At:              publicText(item.At),
		ItemKey:         normalizeToken(item.ItemKey),
		QueueClass:      normalizeToken(item.QueueClass),
		Risk:            normalizeToken(item.Risk),
		Status:          normalizeToken(item.Status),
		RequesterRole:   normalizeToken(item.RequesterRole),
		ReviewerRole:    normalizeToken(item.ReviewerRole),
		BackupAvailable: item.BackupAvailable,
		EvidenceRefs:    normalizeRefs(item.EvidenceRefs),
	}
	if err := validateReviewShape(normalized); err != nil {
		return Review{}, err
	}
	if err := validateReviewMembership(policy, normalized); err != nil {
		return Review{}, err
	}
	return normalized, validateReviewRefs(policy, normalized.EvidenceRefs)
}

func validateReviewShape(item Review) error {
	if item.At == "" {
		return fmt.Errorf("review timestamp is required")
	}
	if _, err := time.Parse(time.RFC3339, item.At); err != nil {
		return fmt.Errorf("review timestamp is invalid")
	}
	if item.ItemKey == "" {
		return fmt.Errorf("review item key is required")
	}
	if item.ReviewerRole == "" {
		return fmt.Errorf("reviewer role is required")
	}
	if len(item.EvidenceRefs) == 0 {
		return errMissingEvidenceRef
	}
	return nil
}

func validateReviewMembership(policy Policy, item Review) error {
	checks := []struct {
		label string
		value string
		have  []string
	}{
		{"queue class", item.QueueClass, policy.QueueClasses},
		{"risk", item.Risk, policy.AllowedRisks},
		{"status", item.Status, policy.AllowedStatuses},
		{"requester role", item.RequesterRole, policy.RequesterRoles},
		{"reviewer role", item.ReviewerRole, policy.ReviewerRoles},
	}
	for _, check := range checks {
		if !contains(normalizeList(check.have), check.value) {
			return fmt.Errorf("review %s %q is not allowed", check.label, check.value)
		}
	}
	return nil
}
