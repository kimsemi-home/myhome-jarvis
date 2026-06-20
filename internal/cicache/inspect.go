package cicache

import "strings"

func inspectUnit(unit graphUnit, artifacts []string, workflow string) UnitStatus {
	status := UnitStatus{
		ID:                          unit.ID,
		Name:                        unit.Name,
		Kind:                        unit.Kind,
		CacheKey:                    unit.Cache,
		HashInputCount:              len(unit.HashInputs),
		GeneratedCoverageCount:      generatedCoverageCount(artifacts, unit.HashInputs),
		CacheRestoreConfigured:      restoreConfigured(unit, workflow),
		CacheHitSkipsVerification:   cacheHitSkips(unit, workflow),
		CacheMissRunsVerification:   cacheMissRuns(unit, workflow),
		PushOnlyCacheSaveConfigured: pushOnlySave(unit, workflow),
	}
	status.PublicSafetyNonSkippable = unit.ID == "public-safety" && unit.Cache == ""
	status.GeneratedArtifactCoverageOK = unit.ID != "ssot" ||
		status.GeneratedCoverageCount == generatedArtifactCount(artifacts)
	status.Valid = unit.Cache == "" || cachedUnitValid(status)
	return status
}

func cachedUnitValid(status UnitStatus) bool {
	return status.HashInputCount > 0 && status.CacheRestoreConfigured &&
		status.CacheHitSkipsVerification && status.CacheMissRunsVerification &&
		status.PushOnlyCacheSaveConfigured && status.GeneratedArtifactCoverageOK
}

func restoreConfigured(unit graphUnit, workflow string) bool {
	if unit.Cache == "" {
		return false
	}
	return strings.Contains(workflow, "path: .github/unit-cache/"+unit.Cache) &&
		strings.Contains(workflow, "key: "+unit.Cache+"-${{ runner.os }}")
}

func cacheHitSkips(unit graphUnit, workflow string) bool {
	if unit.Cache == "" {
		return false
	}
	return strings.Contains(workflow, "Report "+unit.Name+" cache hit") &&
		strings.Contains(workflow, "steps.unit-cache.outputs.cache-hit == 'true'")
}

func cacheMissRuns(unit graphUnit, workflow string) bool {
	if unit.Cache == "" {
		return false
	}
	return strings.Contains(workflow, "Run "+unit.Name+" verification") &&
		strings.Contains(workflow, "steps.unit-cache.outputs.cache-hit != 'true'")
}

func pushOnlySave(unit graphUnit, workflow string) bool {
	if unit.Cache == "" {
		return false
	}
	return strings.Contains(workflow, "Save "+unit.Name+" unit cache") &&
		strings.Contains(workflow, "github.event_name == 'push'")
}
