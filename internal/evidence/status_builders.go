package evidence

import "time"

func newStatus(policy Policy, checkedAt time.Time) Status {
	return Status{
		PolicyPath:  PolicyRelativePath,
		PrivateRoot: policy.PrivateRoot,
		SourceCount: len(policy.PrivateSources),
		ByNodeKind:  map[string]int{},
		ByEdgeKind:  map[string]int{},
		CheckedAt:   checkedAt.Format(time.RFC3339),
	}
}

func newSourceStatus(source PrivateSource) SourceStatus {
	return SourceStatus{
		Key:      source.Key,
		NodeKind: source.NodeKind,
		Format:   source.Format,
	}
}
