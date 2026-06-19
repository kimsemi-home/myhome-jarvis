package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func verifyGraphArtifacts(root string, graph verificationGraphFile) error {
	listed := stringSet(graph.GeneratedArtifacts)
	for _, rel := range requiredVerificationArtifacts() {
		if !listed[rel] {
			return fmt.Errorf("verification graph missing generated artifact %q", rel)
		}
		if err := requirePublicRepoFile(root, rel); err != nil {
			return err
		}
	}
	for _, backend := range graph.Backends {
		if backend.ID == "" || backend.Path == "" {
			return fmt.Errorf("verification backend must declare id and path")
		}
		if err := requirePublicRepoFile(root, backend.Path); err != nil {
			return err
		}
	}
	return nil
}

func requirePublicRepoFile(root string, rel string) error {
	if filepath.IsAbs(rel) || strings.Contains(rel, "..") || strings.HasPrefix(rel, "data/private/") {
		return fmt.Errorf("verification artifact path is not public-safe: %q", rel)
	}
	info, err := os.Stat(filepath.Join(root, filepath.FromSlash(rel)))
	if err != nil {
		return err
	}
	if info.IsDir() {
		return fmt.Errorf("verification artifact path is a directory: %q", rel)
	}
	return nil
}
