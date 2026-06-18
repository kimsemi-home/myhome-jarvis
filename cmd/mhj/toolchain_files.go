package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

func readTrimmedFile(root string, rel string) (string, error) {
	body, err := os.ReadFile(filepath.Join(root, filepath.FromSlash(rel)))
	if err != nil {
		return "", err
	}
	value := strings.TrimSpace(string(body))
	if value == "" {
		return "", fmt.Errorf("%s is empty", rel)
	}
	return value, nil
}

func parseFirstMatchFile(root string, rel string, pattern string) (string, error) {
	body, err := os.ReadFile(filepath.Join(root, filepath.FromSlash(rel)))
	if err != nil {
		return "", err
	}
	re := regexp.MustCompile(pattern)
	matches := re.FindStringSubmatch(string(body))
	if len(matches) < 2 {
		return "", fmt.Errorf("%s does not match expected toolchain pattern", rel)
	}
	return strings.TrimSpace(matches[1]), nil
}

func parseWorkflowEnv(root string, name string) (string, error) {
	pattern := fmt.Sprintf(`(?m)^\s+%s:\s*"([^"]+)"\s*$`, regexp.QuoteMeta(name))
	return parseFirstMatchFile(root, ".github/workflows/quality.yml", pattern)
}

func generatedProjectGoVersion(root string) (string, error) {
	body, err := os.ReadFile(filepath.Join(root, "generated", "commands.generated.json"))
	if err != nil {
		return "", err
	}
	var catalog struct {
		Project struct {
			GoVersion string `json:"go_version"`
		} `json:"project"`
	}
	if err := json.Unmarshal(body, &catalog); err != nil {
		return "", err
	}
	if strings.TrimSpace(catalog.Project.GoVersion) == "" {
		return "", errors.New("generated project go_version is empty")
	}
	return strings.TrimSpace(catalog.Project.GoVersion), nil
}
