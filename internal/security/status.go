package security

import "time"

func StatusForRoot(root string) (Status, error) {
	current, err := Check(root)
	if err != nil {
		return Status{}, err
	}
	history, cache, err := cachedHistoryAggregate(root)
	if err != nil {
		return Status{}, err
	}
	return Status{
		OK:                  current.OK && history.OK,
		CurrentOK:           current.OK,
		CurrentFindingCount: len(current.Findings),
		HistoryOK:           history.OK,
		HistoryFindingCount: history.FindingCount,
		Cache:               cache,
		CheckedAt:           time.Now().UTC().Format(time.RFC3339),
	}, nil
}
