package codeshape

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"
)

func testPolicy(legacy []LegacyDebtFile) Policy {
	return Policy{
		Context:              "AgentCluster",
		Version:              "v1",
		GeneratedArtifact:    PolicyRelativePath,
		MaxFileLines:         75,
		PublicStatusRedacted: true,
		SourceRoots:          []string{"src"},
		Extensions:           []string{".go"},
		ExcludedPrefixes:     []string{"src/private/"},
		LegacyDebtFiles:      legacy,
		PublicSummaryFields: []string{
			"budget_regression_count",
			"legacy_debt_count",
			"ok",
		},
		Commands: []string{"mhj code-shape status"},
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

func numberedLines(count int) string {
	body := ""
	for index := 0; index < count; index++ {
		body += "line\n"
	}
	return body
}
