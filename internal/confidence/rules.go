package confidence

func normalizeRule(rule CapRule) CapRule {
	rule.Key = normalizeToken(rule.Key)
	rule.When = normalizeToken(rule.When)
	rule.Cap = normalizeToken(rule.Cap)
	return rule
}

func ruleTriggered(condition string, status Status) bool {
	switch condition {
	case "public_safety_not_ok":
		return !status.PublicSafetyOK
	case "latest_quality_failed":
		return status.QualityRecorded && !status.QualityOK
	case "evidence_edge_count_zero":
		return status.EvidenceLinkCount == 0
	case "dangling_evidence_ref_count_positive":
		return status.DanglingEvidenceRefCount > 0
	case "open_learning_count_positive":
		return status.OpenLearningCount > 0
	case "latest_quality_missing":
		return !status.QualityRecorded
	case "evidence_links_and_verification_clear":
		return evidenceBacked(status)
	default:
		return false
	}
}

func evidenceBacked(status Status) bool {
	return status.PublicSafetyOK &&
		status.QualityRecorded &&
		status.QualityOK &&
		status.EvidenceLinkCount > 0 &&
		status.DanglingEvidenceRefCount == 0 &&
		status.OpenLearningCount == 0
}
