package externalbootstrap

import (
	"encoding/json"
	"strings"
	"testing"
	"time"
)

func TestChildRepoStatusVerifiesValidSkeleton(t *testing.T) {
	packet := childPacketFixture(t)
	root := writeValidChildRepo(t, packet)
	status, err := childRepoStatusFromPacket(packet, root, fixedChildTime())
	if err != nil {
		t.Fatal(err)
	}
	if !status.Valid || status.EvidenceState != "ready" ||
		status.NextSafeAction != "use_child_repo_as_cross_repo_evidence" {
		t.Fatalf("child repo status = %#v", status)
	}
	if !status.ContextPackValid || !status.HashCacheValid ||
		!status.PublicSafetyOK || !status.PrivateDataAbsent {
		t.Fatalf("child repo gates = %#v", status)
	}
	body, err := json.Marshal(status)
	if err != nil {
		t.Fatal(err)
	}
	if strings.Contains(string(body), root) || strings.Contains(string(body), "/"+"Users"+"/") {
		t.Fatalf("status leaked local path: %s", body)
	}
}

func TestChildRepoStatusRequiresCheckoutPath(t *testing.T) {
	status, err := childRepoStatusFromPacket(childPacketFixture(t), "", fixedChildTime())
	if err != nil {
		t.Fatal(err)
	}
	if status.Valid || status.CheckoutState != "missing" ||
		status.NextSafeAction != "provide_child_repo_checkout_path" {
		t.Fatalf("missing child repo status = %#v", status)
	}
}

func fixedChildTime() time.Time {
	return time.Date(2026, 6, 21, 7, 0, 0, 0, time.UTC)
}
