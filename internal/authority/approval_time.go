package authority

import (
	"fmt"
	"strings"
	"time"
)

func normalizePacketCheckedAt(value string, now time.Time) (string, error) {
	value = strings.TrimSpace(value)
	checkedAt, err := time.Parse(time.RFC3339, value)
	if err != nil {
		return "", fmt.Errorf("decision packet checked_at must be RFC3339")
	}
	checkedAt = checkedAt.UTC()
	if checkedAt.After(now.Add(5 * time.Minute)) {
		return "", fmt.Errorf("decision packet checked_at is in the future")
	}
	if now.Sub(checkedAt) > approvalPacketFreshnessHours*time.Hour {
		return "", fmt.Errorf("decision packet is stale")
	}
	return checkedAt.Format(time.RFC3339), nil
}

func normalizeApprovalExpiry(value string, now time.Time) (string, error) {
	value = strings.TrimSpace(value)
	expiresAt, err := time.Parse(time.RFC3339, value)
	if err != nil {
		return "", fmt.Errorf("approval expires_at must be RFC3339")
	}
	expiresAt = expiresAt.UTC()
	if !expiresAt.After(now) {
		return "", fmt.Errorf("approval lease must expire in the future")
	}
	if expiresAt.Sub(now) > 7*24*time.Hour {
		return "", fmt.Errorf("approval lease must be seven days or less")
	}
	return expiresAt.Format(time.RFC3339), nil
}
