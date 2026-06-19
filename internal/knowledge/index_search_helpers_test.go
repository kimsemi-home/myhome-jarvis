package knowledge

import "testing"

type searchExpectation struct {
	query    string
	concept  string
	mustRead []string
}

func assertSearchEvidence(t *testing.T, exp searchExpectation) SearchReport {
	t.Helper()
	report, err := Search(repoRoot(t), exp.query)
	if err != nil {
		t.Fatal(err)
	}
	if !hasConcept(report.Concepts, exp.concept) {
		t.Fatalf("expected %s concept, got %#v", exp.concept, report.Concepts)
	}
	for _, path := range exp.mustRead {
		if !containsString(report.MustRead, path) {
			t.Fatalf("must read missing %s: %#v", path, report.MustRead)
		}
	}
	return report
}
