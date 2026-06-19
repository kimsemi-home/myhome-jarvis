package linear

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func assertReplayEvidenceRedacted(t *testing.T, root string) {
	t.Helper()
	replayData, err := os.ReadFile(filepath.Join(root, OfflineReplayRelativePath))
	if err != nil {
		t.Fatal(err)
	}
	body := string(replayData)
	if strings.Contains(body, "Started work") || strings.Contains(body, "not replay-safe") {
		t.Fatalf("replay evidence leaked raw payload: %s", body)
	}
}
