package translation

import "testing"

func TestStatusCountsForbiddenLoss(t *testing.T) {
	root := t.TempDir()
	writePolicy(t, root, testPolicy())
	writeFile(t, root, "data/private/translation/manifests/forbidden.json", `{
  "source_context": "AgentCluster",
  "target_context": "SecurityPolicy",
  "source_version": "v1",
  "target_version": "v1",
  "preserved_rules": ["public_safety"],
  "known_losses": [{"level": "l4_forbidden", "category": "security_boundary"}],
  "owner": "security",
  "evidence_refs": ["generated/security.generated.json"]
}`)

	status, err := StatusForRoot(root)
	if err != nil {
		t.Fatal(err)
	}
	if status.ForbiddenLossCount != 1 || status.OpenDebtCount != 1 {
		t.Fatalf("status = %#v", status)
	}
}
