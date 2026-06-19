package knowledge

import "testing"

func TestSearchHumanReviewCapacityReturnsQueueEvidence(t *testing.T) {
	assertSearchEvidence(t, searchExpectation{
		query:   "human review capacity backup steward review debt",
		concept: "HumanReviewCapacity",
		mustRead: []string{
			"generated/review.generated.json",
			"docs/human-review-capacity.md",
		},
	})
}

func TestSearchIncidentLifecycleReturnsIncidentEvidence(t *testing.T) {
	assertSearchEvidence(t, searchExpectation{
		query:   "incident lifecycle quarantine owner",
		concept: "IncidentLifecycle",
		mustRead: []string{
			"generated/incidents.generated.json",
			"docs/incident-lifecycle.md",
		},
	})
}

func TestSearchDomainEventReturnsEventEvidence(t *testing.T) {
	report := assertSearchEvidence(t, searchExpectation{
		query:    "DomainEvent",
		concept:  "CheckpointRecorded",
		mustRead: []string{"internal/orchestrator/checkpoint.go"},
	})
	for _, event := range []string{"CheckpointRecorded", "KnowledgeLookupRecorded"} {
		if !hasEvent(report.Events, event) {
			t.Fatalf("expected %s event, got %#v", event, report.Events)
		}
	}
}
