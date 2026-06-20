package codexcost

func costSharePercent(costUnits int64, totalUnits int64) int {
	if costUnits <= 0 || totalUnits <= 0 {
		return 0
	}
	return int((costUnits * 100) / totalUnits)
}

func allocateValueProxy(valueProxy int64, costUnits int64, totalUnits int64) int64 {
	if valueProxy <= 0 || costUnits <= 0 || totalUnits <= 0 {
		return 0
	}
	return (valueProxy * costUnits) / totalUnits
}

func roiCostPerAcceptedChange(
	totalUnits int64,
	acceptedChanges int64,
	fallback int64,
) int64 {
	if totalUnits > 0 && acceptedChanges > 0 {
		return totalUnits / acceptedChanges
	}
	return fallback
}

func roiCostPerChange(costUnits int64, costPerChange int64) int64 {
	if costUnits <= 0 || costPerChange <= 0 {
		return 0
	}
	return costPerChange
}

func roiScopeStatus(directUnits int64, attributedUnits int64) string {
	if attributedUnits > 0 {
		return "attributed"
	}
	if directUnits > 0 {
		return "tracked"
	}
	return "no_usage_yet"
}
