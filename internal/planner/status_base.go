package planner

import "time"

func newStatus(policy Policy) Status {
	return Status{
		LoopMode:               policy.LoopMode,
		TaskCount:              len(policy.TaskGraph),
		LinearTemplateCount:    len(policy.LinearTemplates),
		QualityRequired:        policy.QualityRequired,
		LinearOfflineFallback:  policy.LinearOfflineFallback,
		KnowledgeIndexRequired: policy.KnowledgeIndexRequiredBeforePlanning,
		CheckpointRoot:         policy.CheckpointRoot,
		CheckedAt:              time.Now().UTC().Format(time.RFC3339),
	}
}
