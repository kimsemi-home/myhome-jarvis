package repofactory

import (
	"encoding/json"
	"strings"
	"testing"

	"github.com/kimsemi-home/myhome-jarvis/internal/security"
)

func TestDecisionPacketMarksPublicSafetyEvidenceReady(t *testing.T) {
	evidence := publicSafetyEvidenceFromStatus(security.Status{
		OK:        true,
		CurrentOK: true,
		HistoryOK: true,
	})
	packet := decisionPacketFromPolicyEvidence(testPolicy(), evidence)
	if !packet.PublicSafetyEvidence.OK ||
		packet.PublicSafetyEvidence.EvidenceState != "ready" {
		t.Fatalf("public safety evidence = %#v", packet.PublicSafetyEvidence)
	}
	if packet.BlockingGateCount != 1 ||
		contains(packet.MissingEvidenceKeys, "public_safety_evidence") {
		t.Fatalf("missing evidence = %#v", packet)
	}
}

func TestDecisionPacketDoesNotExposeRawSecurityPayload(t *testing.T) {
	evidence := publicSafetyEvidenceFromStatus(security.Status{
		CurrentFindingCount: 1,
		HistoryFindingCount: 1,
	})
	packet := decisionPacketFromPolicyEvidence(testPolicy(), evidence)
	body, err := json.Marshal(packet)
	if err != nil {
		t.Fatal(err)
	}
	for _, forbidden := range []string{
		"/" + "Users" + "/", "linear" + ".app", "access" + "_token",
		"client" + "_secret", "private" + "_key", "raw finding",
	} {
		if strings.Contains(string(body), forbidden) {
			t.Fatalf("decision packet leaked %q in %s", forbidden, body)
		}
	}
}
