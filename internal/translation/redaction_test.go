package translation

import (
	"encoding/json"
	"path/filepath"
	"strings"
	"testing"
)

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
	for _, forbidden := range forbiddenTranslationStatusMarkers() {
		if strings.Contains(body, forbidden) {
			t.Fatalf("translation status leaked %q in %s", forbidden, body)
		}
	}
}

func forbiddenTranslationStatusMarkers() []string {
	return []string{
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
	}
}
