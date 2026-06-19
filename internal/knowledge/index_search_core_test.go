package knowledge

import (
	"encoding/json"
	"strings"
	"testing"
)

func TestSearchReturnsKnowledgeEvidenceWithoutSnippets(t *testing.T) {
	report := assertSearchEvidence(t, searchExpectation{
		query:   "KnowledgeIndex",
		concept: "KnowledgeIndex",
		mustRead: []string{
			"generated/concepts.generated.json",
			"docs/knowledge-index.md",
		},
	})
	if len(report.Hits) == 0 {
		t.Fatal("expected lexical hits")
	}
	if !containsString(report.LinearIssues, "KIM-14") {
		t.Fatalf("linear issues missing KIM-14: %#v", report.LinearIssues)
	}
	payload, err := json.Marshal(report)
	if err != nil {
		t.Fatal(err)
	}
	body := string(payload)
	for _, forbidden := range []string{
		repoRoot(t),
		"A local lexical index over SSOT",
		"raw private queue contents",
	} {
		if strings.Contains(body, forbidden) {
			t.Fatalf("search report leaked %q in %s", forbidden, body)
		}
	}
}
