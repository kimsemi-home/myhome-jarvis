package codexcost

func roiRowsByScope(rows []ROIRow) map[string]ROIRow {
	byScope := map[string]ROIRow{}
	for _, row := range rows {
		byScope[row.Scope] = row
	}
	return byScope
}
