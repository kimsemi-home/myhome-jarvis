package localfinancereadiness

import (
	"os"
	"path/filepath"
	"testing"
)

func fixturePath(parts ...string) string {
	values := append([]string{"..", "..", "fixtures", "local_finance_readiness"}, parts...)
	return filepath.Join(values...)
}

func TestFixtureReadinessManifestValidates(t *testing.T) {
	manifest, err := Read(fixturePath("manifest.json"))
	if err != nil {
		t.Fatal(err)
	}
	if len(manifest.Plans) != 4 || len(manifest.Stages) != 5 || manifest.Stages[3].Component != "finance-operator" || manifest.Stages[4].Component != "jarvis" {
		t.Fatalf("unexpected manifest: %#v", manifest)
	}
}

func TestWriteCapableScopeFails(t *testing.T) {
	file, err := os.Open(fixturePath("plans", "revenue.json"))
	if err != nil {
		t.Fatal(err)
	}
	plan, err := decodeOne[Plan](file)
	file.Close()
	if err != nil {
		t.Fatal(err)
	}
	plan.OAuthScopes = append(plan.OAuthScopes, "https://www.googleapis.com/auth/youtube.upload")
	plan.PlanHash = planHash(plan)
	ref := Ref{Component: plan.Component, PlanHash: plan.PlanHash}
	if validatePlan(plan, ref) == nil {
		t.Fatal("write-capable OAuth scope was accepted")
	}
}

func TestOutOfOrderStageFails(t *testing.T) {
	manifest, err := Read(fixturePath("manifest.json"))
	if err != nil {
		t.Fatal(err)
	}
	manifest.Stages[1].Day = 1
	manifest.AggregateHash = aggregateHash(manifest)
	if Validate(manifest) == nil {
		t.Fatal("out-of-order stage was accepted")
	}
}

func TestDirectChildSchedulesFail(t *testing.T) {
	manifest, err := Read(fixturePath("manifest.json"))
	if err != nil {
		t.Fatal(err)
	}
	manifest.DirectChildSchedulesEnabled = true
	manifest.AggregateHash = aggregateHash(manifest)
	if Validate(manifest) == nil {
		t.Fatal("direct child schedules were accepted")
	}
}
