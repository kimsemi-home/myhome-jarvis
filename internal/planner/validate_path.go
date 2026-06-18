package planner

import (
	"errors"
	"fmt"
	"path/filepath"
	"strings"
)

func validateExternalWriteGate(gate ExternalWriteGate) error {
	if !gate.StandingBoundary {
		return errors.New("planner external-write gate must keep a standing boundary")
	}
	if !gate.ApprovalRequired {
		return errors.New("planner external-write gate must require approval")
	}
	if !gate.MutationSuccessRequired {
		return errors.New("planner external-write gate must require mutation success")
	}
	if strings.TrimSpace(gate.BoundaryTaskID) == "" {
		return errors.New("planner external-write gate boundary task id is required")
	}
	return validatePrivateRepoPath("planner external-write gate evidence path", gate.EvidencePath)
}

func validateCheckpointRoot(value string) error {
	return validatePrivateRepoPath("planner checkpoint root", value)
}

func validatePrivateRepoPath(name string, value string) error {
	value = strings.TrimSpace(value)
	if value == "" {
		return fmt.Errorf("%s is required", name)
	}
	native := filepath.FromSlash(value)
	if filepath.IsAbs(native) {
		return fmt.Errorf("%s must be repo-relative", name)
	}
	clean := filepath.ToSlash(filepath.Clean(native))
	if clean == "." || clean == ".." || strings.HasPrefix(clean, "../") {
		return fmt.Errorf("%s must stay inside the repo", name)
	}
	if !strings.HasPrefix(clean, "data/private/") {
		return fmt.Errorf("%s must stay under data/private", name)
	}
	return nil
}
