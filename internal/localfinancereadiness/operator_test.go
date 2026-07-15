package localfinancereadiness

import (
	"encoding/json"
	"os"
	"testing"
)

func TestFixtureReadinessAggregateHashIsCurrent(t *testing.T) {
	body, err := os.ReadFile(fixturePath("manifest.json"))
	if err != nil {
		t.Fatal(err)
	}
	var manifest Manifest
	if err := json.Unmarshal(body, &manifest); err != nil {
		t.Fatal(err)
	}
	if actual := aggregateHash(manifest); manifest.AggregateHash != actual {
		t.Fatalf("fixture readiness aggregate hash = %s", actual)
	}
}

func TestMutableOperatorStageFails(t *testing.T) {
	file, err := os.Open(fixturePath("plans", "operator.json"))
	if err != nil {
		t.Fatal(err)
	}
	plan, err := decodeOne[OperatorPlan](file)
	file.Close()
	if err != nil {
		t.Fatal(err)
	}
	plan.Stages[0].DueDay = 1
	plan.PlanHash = operatorPlanHash(plan)
	ref := Ref{Component: plan.Component, PlanHash: plan.PlanHash}
	if validateOperatorPlan(plan, ref) == nil {
		t.Fatal("mutable Finance Operator stage was accepted")
	}
}
