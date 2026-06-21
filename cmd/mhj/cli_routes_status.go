package main

func routeStatuses(root string, args []string) (bool, error) {
	if ok, err := routeAuthorityReviewStatus(root, args); ok {
		return true, err
	}
	if len(args) != 2 || args[1] != "status" {
		return false, nil
	}
	switch args[0] {
	case "assistant":
		return true, assistantStatus(root)
	case "work-item":
		return true, workItemStatus(root)
	case "audit":
		return true, auditStatus(root)
	case "code-shape":
		return true, codeShapeStatus(root)
	case "context-pack":
		return true, contextPackStatus(root)
	case "ci-cache":
		return true, ciCacheStatus(root)
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
	case "external-evidence":
		return true, externalEvidenceStatus(root)
	default:
		return routeMoreStatuses(root, args)
	}
}
