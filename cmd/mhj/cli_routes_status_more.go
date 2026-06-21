package main

func routeMoreStatuses(root string, args []string) (bool, error) {
	switch args[0] {
	case "finance-consent":
		return true, financeConsentStatus(root)
	case "codex-cost":
		return true, codexCostStatus(root)
	case "codex-sustainability":
		return true, codexSustainabilityStatus(root)
	case "media-readiness":
		return true, mediaReadinessStatus(root)
	case "merge-evidence":
		return true, mergeEvidenceStatus(root)
	case "storage-archive":
		return true, storageArchiveStatus(root)
	case "monetization":
		return true, monetizationStatus(root)
	case "review":
		return true, reviewStatus(root)
	case "authority":
		return true, authorityStatus(root)
	case "authority-review":
		return true, authorityReviewStatus(root)
	case "pdca":
		return true, pdcaStatus(root)
	case "repo":
		return true, repoStatus(root)
	case "repo-factory":
		return true, repoFactoryStatus(root)
	case "planner":
		return true, plannerStatus(root)
	case "quality":
		return true, qualityStatus(root)
	default:
		return false, nil
	}
}
