package localfinancereadiness

import (
	"os"
	"slices"
	"testing"
)

func TestIncompletePortfolioKISBoundaryFails(t *testing.T) {
	file, err := os.Open(fixturePath("plans", "portfolio.json"))
	if err != nil {
		t.Fatal(err)
	}
	plan, err := decodeOne[Plan](file)
	file.Close()
	if err != nil {
		t.Fatal(err)
	}
	plan.Checks = slices.DeleteFunc(plan.Checks, func(check string) bool {
		return check == "kis-exact-official-origin"
	})
	plan.PlanHash = planHash(plan)
	ref := Ref{Component: plan.Component, PlanHash: plan.PlanHash}
	if validatePlan(plan, ref) == nil {
		t.Fatal("incomplete Portfolio KIS boundary was accepted")
	}
}
