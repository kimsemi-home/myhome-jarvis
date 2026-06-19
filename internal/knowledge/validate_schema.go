package knowledge

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func validateKnowledgeSchema(registry Registry) []string {
	var failures []string
	schema := registry.KnowledgeIndexSchema
	if schema.Kind != "local-lexical" {
		failures = append(failures, "knowledge index must be local-lexical")
	}
	if schema.CloudRAGAllowed || schema.ExternalVectorDBAllowed {
		failures = append(failures, "knowledge index must not allow cloud RAG or external vector DB")
	}
	if len(schema.IndexRoots) == 0 {
		failures = append(failures, "knowledge index roots are required")
	}
	for _, rootPath := range schema.IndexRoots {
		if strings.HasPrefix(rootPath, "/") || strings.Contains(rootPath, "..") {
			failures = append(failures, fmt.Sprintf("knowledge index root %q must be repo-relative", rootPath))
		}
	}
	return failures
}

func validatePlanningRules(registry Registry) []string {
	if registry.PlanningRules.KnowledgeIndexRequiredBeforePlanning {
		return nil
	}
	return []string{"planning rules must require KnowledgeIndex before planning"}
}

func requirePublicTarget(root string, rel string) error {
	rel = clean(rel)
	if rel == "" {
		return fmt.Errorf("path is required")
	}
	if filepath.IsAbs(filepath.FromSlash(rel)) || strings.Contains(rel, "..") {
		return fmt.Errorf("path must be repo-relative")
	}
	if strings.HasPrefix(rel, "data/private/") {
		return fmt.Errorf("generated target must not be private")
	}
	if _, err := os.Stat(filepath.Join(root, filepath.FromSlash(rel))); err != nil {
		return err
	}
	return nil
}
