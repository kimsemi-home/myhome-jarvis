package planner

import "path/filepath"

const generatedRelativePath = "generated/planner.generated.json"

func StatusForRoot(root string) (Status, error) {
	policy, err := loadPolicy(root)
	if err != nil {
		return Status{}, err
	}
	status := newStatus(policy)
	if err := attachWriteEvidence(root, policy, &status); err != nil {
		return Status{}, err
	}
	if err := attachKnowledgeEvidence(root, policy, &status); err != nil {
		return Status{}, err
	}
	summarizeTasks(policy.TaskGraph, &status)
	return status, nil
}

func loadPolicy(root string) (Policy, error) {
	policy, err := ReadPolicy(filepath.Join(root, filepath.FromSlash(generatedRelativePath)))
	if err != nil {
		return Policy{}, err
	}
	if err := validatePolicy(policy); err != nil {
		return Policy{}, err
	}
	return policy, nil
}
