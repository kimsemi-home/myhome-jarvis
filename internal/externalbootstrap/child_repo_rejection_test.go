package externalbootstrap

import (
	"strings"
	"testing"
)

func TestChildRepoStatusDetectsContextAndHashDrift(t *testing.T) {
	packet := childPacketFixture(t)
	root := writeValidChildRepo(t, packet)
	context := childContextFixture(packet)
	context.OntologyVersion = "concept-registry/v0"
	writeChildJSON(t, root, ".mhj/context-pack.json", context)
	hashes := childHashFixture(packet)
	hashes.HashCacheInputs[0].SHA256 = strings.Repeat("0", 64)
	writeChildJSON(t, root, ".mhj/hash-cache-inputs.json", hashes)
	status, err := childRepoStatusFromPacket(packet, root, fixedChildTime())
	if err != nil {
		t.Fatal(err)
	}
	if status.Valid || status.EvidenceState != "drifted" ||
		status.DriftCount == 0 || status.InvalidHashCacheCount == 0 {
		t.Fatalf("drifted child repo status = %#v", status)
	}
}

func TestChildRepoStatusDetectsPrivateMaterial(t *testing.T) {
	packet := childPacketFixture(t)
	root := writeValidChildRepo(t, packet)
	writeChildFile(t, root, "docs/leak.md", strings.Join([]string{"kim", "jooyoon"}, ""))
	status, err := childRepoStatusFromPacket(packet, root, fixedChildTime())
	if err != nil {
		t.Fatal(err)
	}
	if status.Valid || status.PublicSafetyOK || status.PublicSafetyFindingCount == 0 {
		t.Fatalf("unsafe child repo status = %#v", status)
	}
}
