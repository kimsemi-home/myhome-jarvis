package agentcluster

import (
	"os"
	"path/filepath"
	"testing"
)

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
