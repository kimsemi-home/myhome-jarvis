package repofactory

import (
	"encoding/json"
	"strings"
	"testing"
)

func TestDecisionPacketKeepsRepoCreationBlockedForReview(t *testing.T) {
	root := t.TempDir()
	writePolicy(t, root, testPolicy())

	packet, err := DecisionPacketForRoot(root)
	if err != nil {
		t.Fatal(err)
	}
	if packet.CreationAllowed || !packet.RepoCreationBlockedUntilReview {
		t.Fatalf("creation gate = %#v", packet)
	}
	if packet.TemplateReadyCount != 6 || packet.GateReadyCount != 3 {
		t.Fatalf("readiness counts = %#v", packet)
	}
	if packet.BlockingGateCount != 2 ||
		!contains(packet.MissingEvidenceKeys, "authority_review") ||
		!contains(packet.MissingEvidenceKeys, "public_safety_evidence") {
		t.Fatalf("missing evidence = %#v", packet)
	}
	if packet.NextSafeAction != "await_human_authority_review" {
		t.Fatalf("next action = %q", packet.NextSafeAction)
	}
}

func TestDecisionPacketRedactsUnsafeTemplatePath(t *testing.T) {
	policy := testPolicy()
	policy.TemplateFiles[0].Path = "/" + "Users" + "/private/repo"

	packet := decisionPacketFromPolicy(policy)
	body, err := json.Marshal(packet)
	if err != nil {
		t.Fatal(err)
	}
	if packet.PublicSafe || packet.CreationDecision != "repair_repo_factory_policy" {
		t.Fatalf("packet = %#v", packet)
	}
	if strings.Contains(string(body), "/"+"Users"+"/") {
		t.Fatalf("decision packet leaked private path: %s", string(body))
	}
	if !strings.Contains(string(body), "redacted_forbidden_template_path") {
		t.Fatalf("decision packet did not mark redaction: %s", string(body))
	}
}
