package agentcluster

import (
	"encoding/json"
	"os"
	"path/filepath"
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

func TestStatusRejectsSelfApproval(t *testing.T) {
	root := t.TempDir()
	generated := filepath.Join(root, filepath.FromSlash(GeneratedPolicyPath))
	if err := os.MkdirAll(filepath.Dir(generated), 0o755); err != nil {
		t.Fatal(err)
	}
	body := `{"context":"AgentCluster","version":"v1","public_safe":true,"raw_transcript_storage_allowed":false,"private_data_in_evidence_allowed":false,"external_agent_execution_allowed":false,"self_approval_allowed":true,"confidence_self_report_allowed":false,"authority_gate_required":true,"evidence_flow":["observation","evidence","code","verification_evidence","knowledge_update"],"incident_lifecycle":["observed","evidence_recorded","classified","owner_assigned","verified","knowledge_updated"],"agent_roles":[{"key":"producer","label":"Producer","reasoning_tier":"R2","authority":"propose","must_produce":["evidence"],"must_not":["self_approve"]},{"key":"reviewer","label":"Reviewer","reasoning_tier":"R3","authority":"review","must_produce":["risk"],"must_not":["self_approve"]},{"key":"verifier","label":"Verifier","reasoning_tier":"R0","authority":"verify","must_produce":["test"],"must_not":["mutate"]},{"key":"steward","label":"Steward","reasoning_tier":"R4","authority":"gate","must_produce":["approval"],"must_not":["erase"]}],"sidecars":[{"key":"verification","label":"Verification","checks":["tests"]},{"key":"confidence","label":"Confidence","checks":["evidence"]},{"key":"security","label":"Security","checks":["boundary"]},{"key":"control","label":"Control","checks":["manifest"]}],"commands":["mhj agent-cluster status"]}`
	if err := os.WriteFile(generated, []byte(body), 0o644); err != nil {
		t.Fatal(err)
	}

	if _, err := StatusForRoot(root); err == nil {
		t.Fatal("expected self-approval policy to fail")
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
