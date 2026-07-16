package localfinancereadiness

import (
	"os"
	"testing"
)

func TestMissingLedgerOAuthBoundaryFails(t *testing.T) {
	file, err := os.Open(fixturePath("plans", "ledger.json"))
	if err != nil {
		t.Fatal(err)
	}
	plan, err := decodeOne[Plan](file)
	file.Close()
	if err != nil {
		t.Fatal(err)
	}
	checks := make([]string, 0, len(plan.Checks)-1)
	for _, check := range plan.Checks {
		if check != "oauth-token-origin-pinned" {
			checks = append(checks, check)
		}
	}
	plan.Checks = checks
	plan.PlanHash = planHash(plan)
	ref := Ref{Component: plan.Component, PlanHash: plan.PlanHash}
	if validatePlan(plan, ref) == nil {
		t.Fatal("missing Ledger OAuth origin pin was accepted")
	}
}
