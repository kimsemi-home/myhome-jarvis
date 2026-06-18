package main

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

func runCIVerify(root string) error {
	if err := validateCIWorkflowContract(root); err != nil {
		return err
	}
	return writeJSON(map[string]any{"ok": true})
}

func validateCIWorkflowContract(root string) error {
	body, err := os.ReadFile(filepath.Join(root, ".github", "workflows", "quality.yml"))
	if err != nil {
		return err
	}
	workflow := string(body)
	required := []string{
		"cancel-in-progress: true", "permissions:", "contents: read",
		"fetch-depth: 0", "go run ./cmd/mhj security check",
		"go run ./cmd/mhj security history", "go run ./cmd/mhj ci verify",
		"go run ./cmd/mhj code-shape status", "go run ./cmd/mhj toolchain verify",
		"'.go-version'", "'rust-toolchain.toml'", "'generated/*.json'",
		"'generated/commands.generated.json'", "'generated/connectors.generated.json'",
		"'generated/agent_cluster.generated.json'", "'generated/learning.generated.json'",
		"'generated/evidence.generated.json'", "'generated/confidence.generated.json'",
		"'generated/translation.generated.json'", "'generated/control_plane.generated.json'",
		"'generated/incidents.generated.json'", "'generated/evidence_quality.generated.json'",
		"'generated/review.generated.json'", "'generated/code_shape.generated.json'",
		"'generated/authority.generated.json'",
		"github.event_name == 'push' && github.repository == 'kimsemi-home/myhome-jarvis'",
	}
	for _, token := range required {
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
	return nil
}
