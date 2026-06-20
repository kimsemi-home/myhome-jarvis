package codexcost

import (
	"bytes"
	"testing"

	"github.com/kimsemi-home/myhome-jarvis/internal/storagearchive"
)

func TestScalingPacketDoesNotExposePrivateUsageFields(t *testing.T) {
	root := t.TempDir()
	writePolicy(t, root, testPolicy())
	writeSustainabilityPolicy(t, root)
	writeSustainableLedger(t, root)
	writeStoragePolicy(t, root)
	writeFile(t, root, "data/private/codex-cost/usage.jsonl",
		`{"at":"2026-06-19T00:00:00Z","scope":"repo","unit_kind":"codex_tokens","amount":1,"status":"recorded","evidence_refs":["docs/codex-cost-governor.md"],"raw_prompt":"private prompt","raw_transcript":"private transcript","private_notes":"private note"}`+"\n")
	if _, err := storagearchive.RunForRoot(root); err != nil {
		t.Fatal(err)
	}
	packet, err := ScalingPacketForRoot(root)
	if err != nil {
		t.Fatal(err)
	}
	for _, forbidden := range scalingPacketForbiddenFields() {
		if bytes.Contains(mustJSON(t, packet), forbidden) {
			t.Fatalf("scaling packet leaked %s in %#v", forbidden, packet)
		}
	}
}

func scalingPacketForbiddenFields() [][]byte {
	return [][]byte{
		[]byte("private prompt"), []byte("private transcript"),
		[]byte("private note"), []byte("ledger_path"),
		[]byte("evidence_refs"), []byte("raw_prompt"),
		[]byte("token"), []byte("credential"),
	}
}
