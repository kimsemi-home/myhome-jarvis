package localfinancereadiness

import (
	"os"
	"testing"
)

func TestRevenueOAuthBoundaryTamperFails(t *testing.T) {
	file, err := os.Open(fixturePath("plans", "revenue.json"))
	if err != nil {
		t.Fatal(err)
	}
	plan, err := decodeOne[Plan](file)
	file.Close()
	if err != nil {
		t.Fatal(err)
	}
	for index, check := range plan.Checks {
		if check == "oauth-official-origin-pinned" {
			plan.Checks = append(plan.Checks[:index], plan.Checks[index+1:]...)
			break
		}
	}
	plan.PlanHash = planHash(plan)
	if validatePlan(plan, Ref{Component: plan.Component, PlanHash: plan.PlanHash}) == nil {
		t.Fatal("Revenue plan without official OAuth origin pin was accepted")
	}
}
