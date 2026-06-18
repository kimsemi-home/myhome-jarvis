package planner

import (
	"errors"
	"strings"
)

func validatePolicy(policy Policy) error {
	if strings.TrimSpace(policy.LoopMode) == "" {
		return errors.New("planner loop mode is required")
	}
	if len(policy.TaskGraph) == 0 {
		return errors.New("planner task graph is required")
	}
	if err := validateCheckpointRoot(policy.CheckpointRoot); err != nil {
		return err
	}
	if err := validateExternalWriteGate(policy.ExternalWriteGate); err != nil {
		return err
	}
	if err := validateTaskGraph(policy); err != nil {
		return err
	}
	if policy.KnowledgeIndexRequiredBeforePlanning &&
		strings.TrimSpace(policy.KnowledgeIndexDefaultQuery) == "" {
		return errors.New("planner knowledge index default query is required")
	}
	return validateDependencies(policy.TaskGraph)
}
