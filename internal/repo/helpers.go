package repo

func shortSHA(head string) string {
	if len(head) <= 12 {
		return head
	}
	return head[:12]
}

func trackedChanges(changes []Change) []Change {
	if changes == nil {
		return []Change{}
	}
	return changes
}

func stringList(values []string) []string {
	if values == nil {
		return []string{}
	}
	return values
}
