package externalbootstrap

import (
	"os"
	"strings"
	"time"
)

func ChildRepoStatusForRoot(root string, childRoot string) (ChildRepoStatus, error) {
	packet, err := PacketForRoot(root)
	if err != nil {
		return ChildRepoStatus{}, err
	}
	return childRepoStatusFromPacket(packet, childRoot, time.Now().UTC())
}

func childRepoStatusFromPacket(
	packet Packet,
	childRoot string,
	now time.Time,
) (ChildRepoStatus, error) {
	status := initialChildRepoStatus(packet, now)
	childRoot = strings.TrimSpace(childRoot)
	if childRoot == "" {
		status.addFinding(".", "missing_checkout", "child repo checkout path is required")
		return finalizeChildRepoStatus(status), nil
	}
	info, err := os.Stat(childRoot)
	if err != nil || !info.IsDir() {
		status.addFinding(".", "missing_checkout", "child repo checkout is not present")
		return finalizeChildRepoStatus(status), nil
	}
	status.CheckoutState = "present"
	checkChildRepoFiles(childRoot, packet, &status)
	checkChildRepoContext(childRoot, packet, &status)
	checkChildRepoHashCache(childRoot, packet, &status)
	checkChildRepoPrivateData(childRoot, &status)
	if err := scanChildRepoPublicSafety(childRoot, &status); err != nil {
		return ChildRepoStatus{}, err
	}
	return finalizeChildRepoStatus(status), nil
}

func initialChildRepoStatus(packet Packet, now time.Time) ChildRepoStatus {
	return ChildRepoStatus{
		Context:                 "ExternalEvidenceChildRepoStatus",
		Version:                 "v1",
		PublicSafe:              true,
		CandidateRepo:           packet.CandidateRepo,
		CheckoutState:           "missing",
		ContextHandoff:          packet.ContextHandoff,
		RequiredHashCacheKeys:   requiredHashCacheKeys(packet),
		RawDetailsPublicAllowed: false,
		CheckedAt:               now.Format(time.RFC3339),
	}
}
