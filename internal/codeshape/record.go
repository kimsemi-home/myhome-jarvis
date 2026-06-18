package codeshape

func recordFile(policy Policy, legacy map[string]int, status *Status, rel string, lines int) {
	status.FileCount++
	if lines > status.MaxObservedLines {
		status.MaxObservedLines = lines
		status.MaxObservedPath = rel
	}
	if lines <= policy.MaxFileLines {
		return
	}
	status.OverBudgetCount++
	legacyMax := legacy[rel]
	finding := FileFinding{
		Path:           rel,
		Lines:          lines,
		MaxLines:       policy.MaxFileLines,
		LegacyMaxLines: legacyMax,
	}
	if legacyMax > 0 && lines <= legacyMax {
		status.LegacyDebtCount++
		status.TopDebt = append(status.TopDebt, finding)
		return
	}
	status.BudgetRegressionCount++
	status.Regressions = append(status.Regressions, finding)
}

func legacyMap(entries []LegacyDebtFile) map[string]int {
	legacy := map[string]int{}
	for _, entry := range entries {
		legacy[entry.Path] = entry.MaxLines
	}
	return legacy
}
