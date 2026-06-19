package knowledge

import "testing"

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
		KnowledgeIndexSchema: IndexSchema{
			Kind: "local-lexical", IndexRoots: []string{"generated"},
		},
	}
	failures := registryFailures(root, registry)
	if !containsFailure(failures, "alias") {
		t.Fatalf("expected alias failure, got %#v", failures)
	}
}
