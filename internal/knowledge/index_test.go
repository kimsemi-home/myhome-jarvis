package knowledge

import (
	"encoding/json"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestVerifyGeneratedRegistry(t *testing.T) {
	report, err := Verify(repoRoot(t))
	if err != nil {
		t.Fatal(err)
	}
	if !report.OK {
		t.Fatalf("verify failed: %#v", report)
	}
	if report.ContextCount != 9 {
		t.Fatalf("context count = %d", report.ContextCount)
	}
	if report.ConceptCount != 14 {
		t.Fatalf("concept count = %d", report.ConceptCount)
	}
	if report.EventCount != 2 {
		t.Fatalf("event count = %d", report.EventCount)
	}
	if report.HarnessCount != 3 {
		t.Fatalf("harness count = %d", report.HarnessCount)
	}
}

func TestSearchReturnsKnowledgeEvidenceWithoutSnippets(t *testing.T) {
	report, err := Search(repoRoot(t), "KnowledgeIndex")
	if err != nil {
		t.Fatal(err)
	}
	if !hasConcept(report.Concepts, "KnowledgeIndex") {
		t.Fatalf("unexpected concepts: %#v", report.Concepts)
	}
	if len(report.Hits) == 0 {
		t.Fatal("expected lexical hits")
	}
	if !containsString(report.MustRead, "generated/concepts.generated.json") {
		t.Fatalf("must read missing concepts artifact: %#v", report.MustRead)
	}
	if !containsString(report.MustRead, "docs/knowledge-index.md") {
		t.Fatalf("must read missing docs/knowledge-index.md: %#v", report.MustRead)
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

func TestSearchConnectorReadinessReturnsCatalogEvidence(t *testing.T) {
	report, err := Search(repoRoot(t), "connector readiness")
	if err != nil {
		t.Fatal(err)
	}
	if !hasConcept(report.Concepts, "ConnectorCatalog") {
		t.Fatalf("expected ConnectorCatalog concept, got %#v", report.Concepts)
	}
	if !containsString(report.MustRead, "generated/connectors.generated.json") {
		t.Fatalf("must read missing connector artifact: %#v", report.MustRead)
	}
	if !containsString(report.MustRead, "docs/connectors.md") {
		t.Fatalf("must read missing docs/connectors.md: %#v", report.MustRead)
	}
}

func TestSearchAgentClusterReturnsPolicyEvidence(t *testing.T) {
	report, err := Search(repoRoot(t), "agent cluster learning loop")
	if err != nil {
		t.Fatal(err)
	}
	if !hasConcept(report.Concepts, "AgentClusterPolicy") {
		t.Fatalf("expected AgentClusterPolicy concept, got %#v", report.Concepts)
	}
	if !containsString(report.MustRead, "generated/agent_cluster.generated.json") {
		t.Fatalf("must read missing agent cluster artifact: %#v", report.MustRead)
	}
	if !containsString(report.MustRead, "docs/agent-cluster.md") {
		t.Fatalf("must read missing docs/agent-cluster.md: %#v", report.MustRead)
	}
}

func TestSearchDomainEventReturnsEventEvidence(t *testing.T) {
	report, err := Search(repoRoot(t), "DomainEvent")
	if err != nil {
		t.Fatal(err)
	}
	if !hasConcept(report.Concepts, "CheckpointRecorded") {
		t.Fatalf("expected CheckpointRecorded concept, got %#v", report.Concepts)
	}
	if !hasEvent(report.Events, "CheckpointRecorded") {
		t.Fatalf("expected CheckpointRecorded event, got %#v", report.Events)
	}
	if !hasEvent(report.Events, "KnowledgeLookupRecorded") {
		t.Fatalf("expected KnowledgeLookupRecorded event, got %#v", report.Events)
	}
	if !containsString(report.MustRead, "internal/orchestrator/checkpoint.go") {
		t.Fatalf("must read missing checkpoint implementation: %#v", report.MustRead)
	}
}

func TestRegistryFailuresDetectDuplicateAlias(t *testing.T) {
	root := t.TempDir()
	writeTestTarget(t, root, "generated/a.json")
	writeTestTarget(t, root, "generated/b.json")
	registry := Registry{
		BoundedContexts: []BoundedContext{{Name: "AgentOps"}},
		Concepts: []Concept{{
			CanonicalName:    "One",
			BoundedContext:   "AgentOps",
			AllowedAliases:   []string{"loop"},
			GeneratedTargets: []string{"generated/a.json"},
		}, {
			CanonicalName:    "Two",
			BoundedContext:   "AgentOps",
			AllowedAliases:   []string{"loop"},
			GeneratedTargets: []string{"generated/b.json"},
		}},
		GeneratedArtifactContracts: []ArtifactContract{{Name: "a", Path: "generated/a.json"}},
		PlanningRules: PlanningRules{
			KnowledgeIndexRequiredBeforePlanning: true,
		},
		KnowledgeIndexSchema: IndexSchema{Kind: "local-lexical", IndexRoots: []string{"generated"}},
	}
	failures := registryFailures(root, registry)
	if !containsFailure(failures, "alias") {
		t.Fatalf("expected alias failure, got %#v", failures)
	}
}

func TestRegistryFailuresDetectInvalidDDDKind(t *testing.T) {
	root := t.TempDir()
	writeTestTarget(t, root, "generated/a.json")
	registry := Registry{
		BoundedContexts: []BoundedContext{{Name: "AgentOps"}},
		DDDPatterns:     []string{"Entity"},
		Concepts: []Concept{{
			CanonicalName:    "One",
			BoundedContext:   "AgentOps",
			DDDKind:          "Bogus",
			AllowedAliases:   []string{"one"},
			GeneratedTargets: []string{"generated/a.json"},
		}},
		DomainEvents: []DomainEvent{{
			Name:           "OneRecorded",
			BoundedContext: "AgentOps",
			EmittedBy:      "One",
			PayloadFields:  []string{"one"},
		}},
		HarnessCaseContracts: []HarnessCase{{
			Name:           "one_harness",
			BoundedContext: "AgentOps",
			Command:        "mhj harness home",
			EvidenceTarget: "generated/a.json",
		}},
		GeneratedArtifactContracts: []ArtifactContract{{Name: "a", Path: "generated/a.json"}},
		PlanningRules: PlanningRules{
			KnowledgeIndexRequiredBeforePlanning: true,
		},
		KnowledgeIndexSchema: IndexSchema{Kind: "local-lexical", IndexRoots: []string{"generated"}},
	}
	failures := registryFailures(root, registry)
	if !containsFailure(failures, "ddd_kind") {
		t.Fatalf("expected ddd_kind failure, got %#v", failures)
	}
}

func containsString(values []string, wanted string) bool {
	for _, value := range values {
		if value == wanted {
			return true
		}
	}
	return false
}

func hasConcept(values []ConceptSummary, wanted string) bool {
	for _, value := range values {
		if value.CanonicalName == wanted {
			return true
		}
	}
	return false
}

func hasEvent(values []DomainEventSummary, wanted string) bool {
	for _, value := range values {
		if value.Name == wanted {
			return true
		}
	}
	return false
}

func containsFailure(values []string, needle string) bool {
	for _, value := range values {
		if strings.Contains(value, needle) {
			return true
		}
	}
	return false
}

func writeTestTarget(t *testing.T, root string, rel string) {
	t.Helper()
	path := filepath.Join(root, filepath.FromSlash(rel))
	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(path, []byte("{}\n"), 0o644); err != nil {
		t.Fatal(err)
	}
}

func repoRoot(t *testing.T) string {
	t.Helper()
	dir, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}
	for {
		if _, err := os.Stat(filepath.Join(dir, "go.mod")); err == nil {
			return dir
		}
		next := filepath.Dir(dir)
		if next == dir {
			t.Fatal("could not find repo root")
		}
		dir = next
	}
}
