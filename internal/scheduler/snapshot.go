package scheduler

func Status(root string, policy Policy) (Snapshot, error) {
	state, err := Recover(root, policy)
	if err != nil {
		return Snapshot{}, err
	}
	return snapshot(policy, state), nil
}

func snapshot(policy Policy, state State) Snapshot {
	return Snapshot{
		Name:                     policy.Name,
		IntervalSeconds:          int64(policy.Interval.Seconds()),
		HeartbeatIntervalSeconds: int64(policy.HeartbeatInterval.Seconds()),
		MinBackoffSeconds:        int64(policy.MinBackoff.Seconds()),
		MaxBackoffSeconds:        int64(policy.MaxBackoff.Seconds()),
		CheckpointEvery:          policy.CheckpointEvery,
		State:                    state,
	}
}
