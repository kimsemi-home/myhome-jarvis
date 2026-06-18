package security

import "sort"

func sortHistoryFindings(findings []HistoryFinding) {
	sort.Slice(findings, func(i, j int) bool {
		left := findings[i]
		right := findings[j]
		if left.Commit != right.Commit {
			return left.Commit < right.Commit
		}
		if left.Path != right.Path {
			return left.Path < right.Path
		}
		if left.Line != right.Line {
			return left.Line < right.Line
		}
		return left.Code < right.Code
	})
}
