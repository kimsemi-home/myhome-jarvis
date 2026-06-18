package translation

import (
	"fmt"
	"strings"
)

func recordLoss(policy Policy, status *Status, sourceContext string, targetContext string, level string, category string, lossStatus string) error {
	sourceContext = strings.TrimSpace(sourceContext)
	targetContext = strings.TrimSpace(targetContext)
	level = normalizeToken(level)
	category = normalizeToken(category)
	lossStatus = normalizeToken(lossStatus)
	if lossStatus == "" {
		lossStatus = "open"
	}
	if !contains(policy.AllowedContexts, sourceContext) || !contains(policy.AllowedContexts, targetContext) {
		return fmt.Errorf("translation loss context is not allowed")
	}
	if !contains(normalizeList(policy.LossLevels), level) ||
		!contains(normalizeList(policy.AllowedLossCategories), category) {
		return fmt.Errorf("translation loss level or category is not allowed")
	}
	recordContext(status, sourceContext, targetContext)
	status.ByLevel[level]++
	if lossStatus == "closed" {
		status.ClosedLossCount++
	} else {
		status.OpenLossCount++
	}
	if level == "l4_forbidden" || contains(normalizeList(policy.ForbiddenLossCategories), category) {
		status.ForbiddenLossCount++
	}
	return nil
}

func recordContext(status *Status, sourceContext string, targetContext string) {
	sourceContext = strings.TrimSpace(sourceContext)
	targetContext = strings.TrimSpace(targetContext)
	if sourceContext != "" {
		status.BySourceContext[sourceContext]++
	}
	if targetContext != "" {
		status.ByTargetContext[targetContext]++
	}
}
