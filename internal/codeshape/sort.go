package codeshape

import "sort"

const maxPublicFindings = 20

func sortFindings(status *Status) {
	sort.Slice(status.TopDebt, func(left int, right int) bool {
		return findingLess(status.TopDebt[left], status.TopDebt[right])
	})
	sort.Slice(status.Regressions, func(left int, right int) bool {
		return findingLess(status.Regressions[left], status.Regressions[right])
	})
	if len(status.TopDebt) > maxPublicFindings {
		status.TopDebt = status.TopDebt[:maxPublicFindings]
	}
	if len(status.Regressions) > maxPublicFindings {
		status.Regressions = status.Regressions[:maxPublicFindings]
	}
}

func findingLess(left FileFinding, right FileFinding) bool {
	if left.Lines != right.Lines {
		return left.Lines > right.Lines
	}
	return left.Path < right.Path
}
