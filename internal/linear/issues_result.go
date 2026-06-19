package linear

func baseOperationResult(root string) OperationResult {
	return OperationResult{
		Mode:      "offline",
		Synced:    false,
		QueuePath: privateRelativePath(filepathJoinSlash(root, "data", "private", "linear-offline-queue.jsonl")),
	}
}
