package confidence

import "time"

func Assess(policy Policy, inputs Inputs) Status {
	status := statusFromInputs(policy, inputs)
	for _, rule := range policy.CapRules {
		rule = normalizeRule(rule)
		rule.Triggered = ruleTriggered(rule.When, status)
		if rule.Triggered {
			status.LevelCap = minLevel(status.LevelCap, rule.Cap)
			if status.ActiveRule == "" {
				status.ActiveRule = rule.Key
			}
		}
	}
	if status.LevelCap == "" {
		status.LevelCap = "high"
	}
	status.Blocked = status.LevelCap == "blocked"
	return status
}

func statusFromInputs(policy Policy, inputs Inputs) Status {
	status := Status{
		PolicyPath:               PolicyRelativePath,
		AssessorKey:              policy.AssessorKey,
		LevelCap:                 "high",
		SelfReportAllowed:        policy.SelfReportAllowed,
		EvidenceLinkCount:        inputs.Evidence.EdgeCount,
		DanglingEvidenceRefCount: inputs.Evidence.DanglingEvidenceRefCount,
		OpenLearningCount:        inputs.Evidence.OpenLearningCount,
		QualityRecorded:          inputs.Quality.Exists && inputs.Quality.Last != nil,
		PublicSafetyOK:           inputs.PublicSafety.OK,
		CheckedAt:                time.Now().UTC().Format(time.RFC3339),
	}
	if inputs.Quality.Last != nil {
		status.QualityOK = inputs.Quality.Last.OK
	}
	return status
}
