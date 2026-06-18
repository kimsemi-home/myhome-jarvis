package translation

import (
	"testing"
)

func TestStatusCountsManifestLossDebt(t *testing.T) {
	root := t.TempDir()
	writePolicy(t, root, testPolicy())
	writeFile(t, root, "data/private/translation/manifests/agent-to-knowledge.json", `{
  "source_context": "AgentCluster",
  "target_context": "KnowledgeIndex",
  "source_version": "v1",
  "target_version": "v1",
  "preserved_rules": ["authority_gate"],
  "known_losses": [{"level": "l1_note", "category": "mapping_gap"}],
  "owner": "knowledge",
  "evidence_refs": ["generated/concepts.generated.json"]
}`)

	status, err := StatusForRoot(root)
	if err != nil {
		t.Fatal(err)
	}
	if status.ManifestCount != 1 || status.LossCount != 1 || status.OpenDebtCount != 1 {
		t.Fatalf("status = %#v", status)
	}
	if status.BySourceContext["AgentCluster"] == 0 || status.ByTargetContext["KnowledgeIndex"] == 0 {
		t.Fatalf("expected context counts, got %#v %#v", status.BySourceContext, status.ByTargetContext)
	}
	if status.ForbiddenLossCount != 0 {
		t.Fatalf("forbidden loss count = %d", status.ForbiddenLossCount)
	}
}

func TestStatusTracksMissingAndMalformedManifestsAsDebt(t *testing.T) {
	root := t.TempDir()
	writePolicy(t, root, testPolicy())
	writeFile(t, root, "data/private/translation/losses.jsonl", `{"at":"2026-06-18T00:00:00Z","source_context":"AgentCluster","target_context":"KnowledgeIndex","level":"l2_degraded","category":"version_drift","status":"open","manifest_path":"data/private/translation/manifests/missing.json","evidence_refs":["generated/concepts.generated.json"]}`+"\n")
	writeFile(t, root, "data/private/translation/manifests/bad.json", `{`)

	status, err := StatusForRoot(root)
	if err != nil {
		t.Fatal(err)
	}
	if status.LedgerExists != true || status.ManifestRootExists != true {
		t.Fatalf("expected ledger and manifest root, got %#v", status)
	}
	if status.MissingManifestCount != 1 || status.InvalidManifestCount != 1 || status.OpenDebtCount != 3 {
		t.Fatalf("status = %#v", status)
	}
	if status.OpenLossCount != 1 || status.ByLevel["l2_degraded"] != 1 {
		t.Fatalf("expected open degraded loss, got %#v", status)
	}
}
