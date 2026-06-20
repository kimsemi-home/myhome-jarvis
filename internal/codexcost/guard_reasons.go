package codexcost

import "github.com/kimsemi-home/myhome-jarvis/internal/codexsustainability"

func guardReasons(
	cost Status,
	projected string,
	sustainability codexsustainability.Status,
) []string {
	reasons := []string{}
	if cost.BudgetState == "review_required" {
		reasons = append(reasons, "current_review_required")
	} else if cost.BudgetState == "warning" {
		reasons = append(reasons, "current_warning")
	}
	if projected == "review_required" {
		reasons = append(reasons, "projected_review_threshold")
	} else if projected == "warning" {
		reasons = append(reasons, "projected_warning_threshold")
	}
	switch sustainability.SustainabilityPosture {
	case "review_required":
		reasons = append(reasons, "sustainability_review_required")
	case "blocked":
		reasons = append(reasons, "sustainability_blocked")
	}
	return reasons
}
