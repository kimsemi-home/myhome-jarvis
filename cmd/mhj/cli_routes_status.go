package main

func routeStatuses(root string, args []string) (bool, error) {
	if len(args) == 2 && args[0] == "authority-review" && args[1] == "request" {
		return true, authorityReviewRequest(root)
	}
	if len(args) != 2 || args[1] != "status" {
		return false, nil
	}
	switch args[0] {
	case "assistant":
		return true, assistantStatus(root)
	case "audit":
		return true, auditStatus(root)
	case "code-shape":
		return true, codeShapeStatus(root)
	case "context-pack":
		return true, contextPackStatus(root)
	case "connectors":
		return true, connectorsStatus(root)
	case "agent-cluster":
		return true, agentClusterStatus(root)
	case "evidence":
		return true, evidenceStatus(root)
	case "evidence-integrity":
		return true, evidenceIntegrityStatus(root)
	case "confidence":
		return true, confidenceStatus(root)
	case "translation":
		return true, translationStatus(root)
	case "control-plane":
		return true, controlPlaneStatus(root)
	case "incidents":
		return true, incidentsStatus(root)
	case "evidence-quality":
		return true, evidenceQualityStatus(root)
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
