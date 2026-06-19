package controlplane

import (
	"encoding/json"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestAppendManifestAndStatusSummarizeRedactedCounts(t *testing.T) {
	root := t.TempDir()
	writePolicy(t, root, testPolicy())

	result, err := AppendManifest(root, ManifestRequest{
		DecisionKind:     "loop_once",
		PolicyVersion:    "control-plane:v1",
		OntologyVersion:  "concepts:v1",
		AuthorityProfile: "local_readonly",
		SelectedRoute:    "loop_once",
		ReviewerRole:     "go_review_gate",
		VerifierRole:     "deterministic_quality_gate",
		LeaseSeconds:     120,
		LeaseStatus:      "finished",
		EvidenceRefs:     []string{"generated/control_plane.generated.json", "generated/planner.generated.json"},
		OutputRef:        "data/private/checkpoints/20260618T000000.000000000Z.json",
	})
	if err != nil {
		t.Fatal(err)
	}
	if result.ManifestPath != "data/private/control-plane/manifests.jsonl" || result.DecisionKind != "loop_once" {
		t.Fatalf("result = %#v", result)
	}

	status, err := StatusForRoot(root)
	if err != nil {
		t.Fatal(err)
	}
	if !status.Exists || status.Count != 1 || status.ManifestDebtCount != 0 {
		t.Fatalf("status = %#v", status)
	}
	if status.ByDecisionKind["loop_once"] != 1 || status.ByAuthorityProfile["local_readonly"] != 1 || status.ByLeaseStatus["finished"] != 1 {
		t.Fatalf("status maps = %#v %#v %#v", status.ByDecisionKind, status.ByAuthorityProfile, status.ByLeaseStatus)
	}
}

func TestStatusCountsMalformedManifestDebt(t *testing.T) {
	root := t.TempDir()
	writePolicy(t, root, testPolicy())
	writeFile(t, root, "data/private/control-plane/manifests.jsonl", `{`+"\n")

	status, err := StatusForRoot(root)
	if err != nil {
		t.Fatal(err)
	}
	if status.InvalidManifestCount != 1 || status.ManifestDebtCount != 1 {
		t.Fatalf("status = %#v", status)
	}
}

func TestAppendManifestRejectsVerifierSelfApproval(t *testing.T) {
	root := t.TempDir()
	writePolicy(t, root, testPolicy())

	_, err := AppendManifest(root, ManifestRequest{
		DecisionKind:     "loop_once",
		PolicyVersion:    "control-plane:v1",
		OntologyVersion:  "concepts:v1",
		AuthorityProfile: "local_readonly",
		SelectedRoute:    "loop_once",
		ReviewerRole:     "same_role",
		VerifierRole:     "same_role",
		LeaseSeconds:     120,
		LeaseStatus:      "finished",
		EvidenceRefs:     []string{"generated/control_plane.generated.json"},
		OutputRef:        "data/private/checkpoints/one.json",
	})
	if err == nil {
		t.Fatal("expected verifier self-approval to fail")
	}
}

func TestStatusJSONDoesNotLeakRawControlPlaneDetails(t *testing.T) {
	root := t.TempDir()
	writePolicy(t, root, testPolicy())
	writeFile(t, root, "data/private/control-plane/manifests.jsonl", `{"id":"cpm_1","at":"2026-06-18T00:00:00Z","decision_kind":"loop_once","policy_version":"control-plane:v1","ontology_version":"concepts:v1","authority_profile":"local_readonly","selected_route":"loop_once","reviewer_role":"go_review_gate","verifier_role":"deterministic_quality_gate","lease_seconds":120,"lease_status":"finished","evidence_refs":["generated/control_plane.generated.json"],"output_ref":"data/private/checkpoints/one.json","raw_rationale":"private reasoning"}`+"\n")

	status, err := StatusForRoot(root)
	if err != nil {
		t.Fatal(err)
	}
	if status.InvalidManifestCount != 1 || status.ManifestDebtCount != 1 {
		t.Fatalf("expected raw rationale to count as debt, got %#v", status)
	}
	payload, err := json.Marshal(status)
	if err != nil {
		t.Fatal(err)
	}
	body := strings.ToLower(string(payload))
	for _, forbidden := range []string{
		"raw_rationale",
		"selection_rationale",
		"candidate_agents",
		"evidence_refs",
		"private reasoning",
		"data/private/checkpoints/one.json",
		"token",
		"secret",
		"linear.app/",
		string(filepath.Separator) + "users" + string(filepath.Separator),
	} {
		if strings.Contains(body, forbidden) {
			t.Fatalf("control-plane status leaked %q in %s", forbidden, body)
		}
	}
}

func TestAppendManifestRejectsSensitiveEvidenceRef(t *testing.T) {
	root := t.TempDir()
	writePolicy(t, root, testPolicy())
	sensitiveRef := "docs/" + "linear.app/" + "kim/private-note.md"

	_, err := AppendManifest(root, ManifestRequest{
		DecisionKind:     "loop_once",
		PolicyVersion:    "control-plane:v1",
		OntologyVersion:  "concepts:v1",
		AuthorityProfile: "local_readonly",
		SelectedRoute:    "loop_once",
		ReviewerRole:     "go_review_gate",
		VerifierRole:     "deterministic_quality_gate",
		LeaseSeconds:     120,
		LeaseStatus:      "finished",
		EvidenceRefs:     []string{sensitiveRef},
		OutputRef:        "data/private/checkpoints/one.json",
	})
	if err == nil {
		t.Fatal("expected sensitive evidence ref to fail")
	}
}

func TestReadPolicyRejectsRawPublicRationale(t *testing.T) {
	root := t.TempDir()
	policy := testPolicy()
	policy.RawRationalePublicAllowed = true
	writePolicy(t, root, policy)

	_, err := ReadPolicy(root)
	if err == nil {
		t.Fatal("expected raw rationale policy to fail")
	}
}

func testPolicy() Policy {
	return Policy{
		Context:                    "AgentOps",
		Version:                    "v1",
		GeneratedArtifact:          "generated/control_plane.generated.json",
		PrivateManifestLedger:      "data/private/control-plane/manifests.jsonl",
		ManifestRequired:           true,
		AppendOnly:                 true,
		PublicStatusRedacted:       true,
		RawRationalePublicAllowed:  false,
		VerifierSeparationRequired: true,
		MinLeaseSeconds:            1,
		MaxLeaseSeconds:            3600,
		AllowedDecisionKinds:       []string{"loop_once", "loop_worker_cycle", "linear_next_observation", "checkpoint_write"},
		AllowedAuthorityProfiles:   []string{"local_readonly", "external_write_gated", "review_only"},
		AllowedLeaseStatuses:       []string{"issued", "active", "expiring", "expired", "finished", "aborted", "escalated", "quarantined"},
		RequiredFields:             []string{"decision_kind", "policy_version", "ontology_version", "authority_profile", "selected_route", "reviewer_role", "verifier_role", "lease_seconds", "lease_status", "evidence_refs", "output_ref"},
		AllowedEvidencePrefixes:    []string{"data/private/", "generated/", "docs/", "cmd/", "internal/", "apps/flutter/", "lisp/", "crates/", "fixtures/", "harness/", ".github/"},
		PublicSummaryFields:        []string{"policy_path", "manifest_path", "exists", "count", "invalid_manifest_count", "manifest_debt_count", "verifier_separation_required", "verifier_violation_count", "min_lease_seconds", "max_lease_seconds", "by_decision_kind", "by_authority_profile", "by_lease_status", "last_observed_at", "checked_at"},
		Commands:                   []string{"mhj control-plane status", "mhj control-plane verify"},
	}
}

func writePolicy(t *testing.T, root string, policy Policy) {
	t.Helper()
	body, err := json.Marshal(policy)
	if err != nil {
		t.Fatal(err)
	}
	writeFile(t, root, PolicyRelativePath, string(body)+"\n")
}

func writeFile(t *testing.T, root string, rel string, body string) {
	t.Helper()
	path := filepath.Join(root, filepath.FromSlash(rel))
	if err := os.MkdirAll(filepath.Dir(path), 0o700); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(path, []byte(body), 0o600); err != nil {
		t.Fatal(err)
	}
}
