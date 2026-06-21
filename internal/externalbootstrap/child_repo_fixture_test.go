package externalbootstrap

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"
	"time"
)

func childPacketFixture(t *testing.T) Packet {
	t.Helper()
	now := time.Date(2026, 6, 21, 6, 0, 0, 0, time.UTC)
	packet, err := packetFromEvidence(
		repoRoot(t),
		splitFixture(now),
		approvalFixture(now),
		factoryFixture(),
		now,
	)
	if err != nil {
		t.Fatal(err)
	}
	return packet
}

func writeValidChildRepo(t *testing.T, packet Packet) string {
	t.Helper()
	root := t.TempDir()
	for _, rel := range childRequiredFiles(packet) {
		writeChildFile(t, root, rel, "public-safe skeleton\n")
	}
	writeChildJSON(t, root, ".mhj/context-pack.json", childContextFixture(packet))
	writeChildJSON(t, root, ".mhj/hash-cache-inputs.json", childHashFixture(packet))
	return root
}

func writeChildFile(t *testing.T, root string, rel string, body string) {
	t.Helper()
	path := filepath.Join(root, filepath.FromSlash(rel))
	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(path, []byte(body), 0o644); err != nil {
		t.Fatal(err)
	}
}

func writeChildJSON(t *testing.T, root string, rel string, value any) {
	t.Helper()
	body, err := json.Marshal(value)
	if err != nil {
		t.Fatal(err)
	}
	writeChildFile(t, root, rel, string(body)+"\n")
}
