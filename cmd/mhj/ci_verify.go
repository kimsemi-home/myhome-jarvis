package main

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/kimsemi-home/myhome-jarvis/internal/cicache"
)

func runCIVerify(root string) error {
	if err := validateCIWorkflowContract(root); err != nil {
		return err
	}
	status, err := cicache.StatusForRoot(root)
	if err != nil {
		return err
	}
	if !status.OK {
		return fmt.Errorf("ci cache contract is not ok")
	}
	return writeJSON(map[string]any{"ok": true})
}

func validateCIWorkflowContract(root string) error {
	body, err := os.ReadFile(filepath.Join(root, ".github", "workflows", "quality.yml"))
	if err != nil {
		return err
	}
	workflow := string(body)
	for _, token := range requiredCIWorkflowTokens() {
		if !strings.Contains(workflow, token) {
			return fmt.Errorf("quality workflow missing CI contract token %q", token)
		}
	}
	for _, token := range []string{"pull_request_target", "write-all"} {
		if strings.Contains(workflow, token) {
			return fmt.Errorf("quality workflow contains forbidden CI contract token %q", token)
		}
	}
	writePermissionPattern := regexp.MustCompile(`(?m)^\s*[A-Za-z0-9_-]+:\s*write\s*$`)
	if match := writePermissionPattern.FindString(workflow); match != "" {
		return fmt.Errorf("quality workflow contains forbidden write permission %q", strings.TrimSpace(match))
	}
	return validateGeneratedWorkflow(root, body)
}

func validateGeneratedWorkflow(root string, workflow []byte) error {
	generated, err := os.ReadFile(filepath.Join(root, "generated", "github_quality_workflow.generated.yml"))
	if err != nil {
		return err
	}
	if !bytes.Equal(workflow, generated) {
		return fmt.Errorf("quality workflow is out of date with generated/github_quality_workflow.generated.yml")
	}
	return nil
}
