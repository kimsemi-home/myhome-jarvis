package confidence

func minLevel(left string, right string) string {
	if levelRank(right) < levelRank(left) {
		return right
	}
	return left
}

func levelRank(level string) int {
	switch normalizeToken(level) {
	case "blocked":
		return 0
	case "low":
		return 1
	case "medium":
		return 2
	case "high":
		return 3
	default:
		return 3
	}
}
