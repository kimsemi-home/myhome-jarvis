package evidence

import (
	"sort"
	"time"
)

const PolicyRelativePath = "generated/evidence.generated.json"

func StatusForRoot(root string) (Status, error) {
	policy, err := ReadPolicy(root)
	if err != nil {
		return Status{}, err
	}
	status := newStatus(policy, time.Now().UTC())
	artifactRefs := map[string]bool{}
	for _, source := range policy.PrivateSources {
		sourceStatus, err := inspectSource(root, policy, source, &status, artifactRefs)
		if err != nil {
			return Status{}, err
		}
		if sourceStatus.Present {
			status.PresentSourceCount++
		}
		status.Sources = append(status.Sources, sourceStatus)
	}
	status.ByNodeKind["evidence_artifact"] += len(artifactRefs)
	status.NodeCount += len(artifactRefs)
	sort.Slice(status.Sources, func(i, j int) bool {
		return status.Sources[i].Key < status.Sources[j].Key
	})
	return status, nil
}
