package evidence

import (
	"fmt"
	"path/filepath"
	"strings"
)

func validatePrivateSources(policy Policy) error {
	if len(policy.PrivateSources) == 0 {
		return fmt.Errorf("evidence graph policy requires private sources")
	}
	nodeKinds := normalizeList(policy.NodeKinds)
	for _, source := range policy.PrivateSources {
		if err := validatePrivateSource(source, nodeKinds); err != nil {
			return err
		}
	}
	return nil
}

func validatePrivateSource(source PrivateSource, nodeKinds []string) error {
	if strings.TrimSpace(source.Key) == "" || strings.TrimSpace(source.Path) == "" {
		return fmt.Errorf("evidence graph source key and path are required")
	}
	if !strings.HasPrefix(source.Path, "data/private/") ||
		filepath.IsAbs(filepath.FromSlash(source.Path)) ||
		strings.Contains(source.Path, "..") {
		return fmt.Errorf("evidence graph source must stay repo-relative under data/private")
	}
	if !contains(nodeKinds, source.NodeKind) {
		return fmt.Errorf("evidence graph source %q has unknown node kind", source.Key)
	}
	if source.Format != "jsonl" && source.Format != "directory" {
		return fmt.Errorf("evidence graph source %q has unsupported format", source.Key)
	}
	return nil
}
