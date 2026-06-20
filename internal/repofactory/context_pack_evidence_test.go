package repofactory

import (
	"encoding/json"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestDecisionPacketIncludesContextPackEvidence(t *testing.T) {
	packet, err := DecisionPacketForRoot(repoRoot(t))
	if err != nil {
		t.Fatal(err)
	}
	evidence := packet.ContextPackEvidence
	if !evidence.Valid || evidence.EvidenceState != "ready" {
		t.Fatalf("context pack evidence = %#v", evidence)
	}
	if evidence.DeclarationPath != ".mhj/context-pack.json" ||
		evidence.OntologyVersion != "concept-registry/v1" ||
		evidence.AuthorityContractVersion != "authority/v1" ||
		evidence.SecurityContractVersion != "security/v1" ||
		evidence.VerificationProfile != "quality" ||
		evidence.ExportedArtifactCount != 6 ||
		evidence.RawDetailsPublicAllowed {
		t.Fatalf("context pack summary = %#v", evidence)
	}
	if packet.CreationAllowed || packet.SelfApprovalAllowed {
		t.Fatalf("context pack evidence granted authority = %#v", packet)
	}
	body, err := json.Marshal(packet)
	if err != nil {
		t.Fatal(err)
	}
	if strings.Contains(string(body), "/"+"Users"+"/") ||
		strings.Contains(string(body), "linear"+".app") {
		t.Fatalf("decision packet leaked private value: %s", body)
	}
}

func repoRoot(t *testing.T) string {
	t.Helper()
	dir, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}
	for {
		if _, err := os.Stat(filepath.Join(dir, "go.mod")); err == nil {
			return dir
		}
		next := filepath.Dir(dir)
		if next == dir {
			t.Fatal("could not locate repo root")
		}
		dir = next
	}
}
