package externalbootstrap

import (
	"encoding/json"
	"strings"
	"testing"
)

func TestPacketRedactsPrivateApprovalLedger(t *testing.T) {
	packet, err := PacketForRoot(repoRoot(t))
	if err != nil {
		t.Fatal(err)
	}
	body, err := json.Marshal(packet)
	if err != nil {
		t.Fatal(err)
	}
	text := string(body)
	for _, forbidden := range []string{
		"data/private", "approvals.jsonl", "reviewer_boundary",
		"/" + "Users" + "/", "raw_private",
	} {
		if strings.Contains(text, forbidden) {
			t.Fatalf("bootstrap packet leaked %q: %s", forbidden, text)
		}
	}
}

func TestPacketIncludesHashCacheEvidence(t *testing.T) {
	packet, err := PacketForRoot(repoRoot(t))
	if err != nil {
		t.Fatal(err)
	}
	for _, key := range []string{
		"generated_artifacts", "source_descriptors",
		"workflow_dependencies", "context_pack_version",
		"ontology_version",
	} {
		if !hasHashCacheKey(packet, key) {
			t.Fatalf("missing cache key %q in %#v", key, packet.HashCacheInputs)
		}
	}
}

func hasHashCacheKey(packet Packet, key string) bool {
	for _, input := range packet.HashCacheInputs {
		if input.Key == key && input.SHA256 != "" && input.PublicSafe {
			return true
		}
	}
	return false
}
