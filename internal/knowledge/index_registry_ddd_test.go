package knowledge

import "testing"

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
		KnowledgeIndexSchema: IndexSchema{
			Kind: "local-lexical", IndexRoots: []string{"generated"},
		},
	}
	failures := registryFailures(root, registry)
	if !containsFailure(failures, "ddd_kind") {
		t.Fatalf("expected ddd_kind failure, got %#v", failures)
	}
}
