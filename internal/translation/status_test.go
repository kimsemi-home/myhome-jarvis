package translation

import (
	"encoding/json"
	"os"
	"path/filepath"
	"strings"
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

func TestStatusJSONDoesNotLeakRawTranslationDetails(t *testing.T) {
	root := t.TempDir()
	writePolicy(t, root, testPolicy())
	writeFile(t, root, "data/private/translation/losses.jsonl", `{"at":"2026-06-18T00:00:00Z","source_context":"AgentCluster","target_context":"KnowledgeIndex","level":"l2_degraded","category":"mapping_gap","status":"open","manifest_path":"data/private/translation/manifests/missing.json","summary":"raw semantic notes","evidence_refs":["data/private/quality/runs.jsonl"]}`+"\n")

	status, err := StatusForRoot(root)
	if err != nil {
		t.Fatal(err)
	}
	payload, err := json.Marshal(status)
	if err != nil {
		t.Fatal(err)
	}
	body := strings.ToLower(string(payload))
	for _, forbidden := range []string{
		"summary",
		"semantic_notes",
		"raw_mapping",
		"known_losses",
		"evidence_refs",
		"data/private/quality/runs.jsonl",
		"raw semantic notes",
		"token",
		"secret",
		string(filepath.Separator) + "users" + string(filepath.Separator),
	} {
		if strings.Contains(body, forbidden) {
			t.Fatalf("translation status leaked %q in %s", forbidden, body)
		}
	}
}

func TestReadPolicyRejectsRawPublicLosses(t *testing.T) {
	root := t.TempDir()
	policy := testPolicy()
	policy.RawLossPublicAllowed = true
	writePolicy(t, root, policy)

	_, err := ReadPolicy(root)
	if err == nil {
		t.Fatal("expected raw public loss policy to fail")
	}
}

func testPolicy() Policy {
	return Policy{
		Context:                 "AgentCluster",
		Version:                 "v1",
		GeneratedArtifact:       "generated/translation.generated.json",
		PrivateLossLedger:       "data/private/translation/losses.jsonl",
		PrivateManifestRoot:     "data/private/translation/manifests",
		ManifestRequired:        true,
		PublicStatusRedacted:    true,
		RawLossPublicAllowed:    false,
		AllowedContexts:         []string{"HomeControl", "HouseholdFinance", "CommerceIntelligence", "ConnectorReadiness", "AgentCluster", "StorageLake", "SecurityPolicy", "AgentOps", "KnowledgeIndex"},
		RequiredManifestFields:  []string{"source_context", "target_context", "source_version", "target_version", "preserved_rules", "known_losses", "owner", "evidence_refs"},
		LossLevels:              []string{"l0_none", "l1_note", "l2_degraded", "l3_review_required", "l4_forbidden"},
		AllowedLossCategories:   []string{"none", "mapping_gap", "version_drift", "field_drop", "precision_loss", "review_needed", "authority", "security_boundary", "user_consent", "deletion_semantics", "audit_record", "legal_obligation", "financial_commitment"},
		ForbiddenLossCategories: []string{"authority", "security_boundary", "user_consent", "deletion_semantics", "audit_record", "legal_obligation", "financial_commitment"},
		AllowedEvidencePrefixes: []string{"data/private/", "generated/", "docs/", "cmd/", "internal/", "apps/flutter/", "lisp/", "crates/", "fixtures/", "harness/", ".github/"},
		PublicSummaryFields:     []string{"policy_path", "ledger_path", "manifest_root", "ledger_exists", "manifest_root_exists", "manifest_count", "invalid_manifest_count", "missing_manifest_count", "loss_count", "open_loss_count", "closed_loss_count", "invalid_loss_count", "open_debt_count", "forbidden_loss_count", "by_level", "by_source_context", "by_target_context", "last_observed_at", "checked_at"},
		Commands:                []string{"mhj translation status"},
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
