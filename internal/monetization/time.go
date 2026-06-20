package monetization

func updateLastObserved(status *Status, candidate string) {
	if candidate > status.LastObservedAt {
		status.LastObservedAt = candidate
	}
}
