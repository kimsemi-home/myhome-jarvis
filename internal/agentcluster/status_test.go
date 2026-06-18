package agentcluster

import (
	"encoding/json"
	"strings"
	"testing"
)

func TestStatusForRootReadsPublicSafeLearningLoopPolicy(t *testing.T) {
	root := repoRoot(t)
	status, err := StatusForRoot(root)
	if err != nil {
		t.Fatal(err)
	}

	if !status.PublicSafe || status.ExternalAgentExecutionAllowed || status.SelfApprovalAllowed || status.ConfidenceSelfReportAllowed {
		t.Fatalf("agent cluster safety flags = %#v", status)
	}
	if !status.AuthorityGateRequired {
		t.Fatal("expected authority gate")
	}
	if status.Context != "AgentCluster" || status.Version == "" {
		t.Fatalf("context/version = %q %q", status.Context, status.Version)
	}
	if status.GeneratedPath != GeneratedPolicyPath {
		t.Fatalf("generated path = %q", status.GeneratedPath)
	}
	if status.RoleCount < 4 || status.SidecarCount < 4 || status.EvidenceStageCount < 6 {
		t.Fatalf("counts = %#v", status)
	}
	if err := requireOrdered(status.EvidenceFlow, "observation", "evidence", "code", "verification_evidence", "knowledge_update"); err != nil {
		t.Fatal(err)
	}
	for _, required := range []string{"observed", "evidence_recorded", "owner_assigned", "verified", "knowledge_updated"} {
		if !contains(status.IncidentLifecycle, required) {
			t.Fatalf("lifecycle missing %q: %#v", required, status.IncidentLifecycle)
		}
	}
	if !contains(status.Commands, "mhj agent-cluster status") {
		t.Fatalf("commands = %#v", status.Commands)
	}
	if len(status.Signals) < 3 {
		t.Fatalf("signals = %#v", status.Signals)
	}

	payload, err := json.Marshal(status)
	if err != nil {
		t.Fatal(err)
	}
	body := strings.ToLower(string(payload))
	for _, forbidden := range []string{
		`"token":`,
		`"secret":`,
		`"credential":`,
		`"cookie":`,
		`"generated_path":"/`,
		`"generated_path":"\\`,
	} {
		if strings.Contains(body, forbidden) {
			t.Fatalf("status leaked forbidden marker %q in %s", forbidden, body)
		}
	}
}
