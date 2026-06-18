package evidence

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"time"
)

const PolicyRelativePath = "generated/evidence.generated.json"

type Policy struct {
	Context                  string          `json:"context"`
	Version                  string          `json:"version"`
	GeneratedArtifact        string          `json:"generated_artifact"`
	PrivateRoot              string          `json:"private_root"`
	PrivateGraphRequired     bool            `json:"private_graph_required"`
	PublicStatusRedacted     bool            `json:"public_status_redacted"`
	RawEvidencePublicAllowed bool            `json:"raw_evidence_public_allowed"`
	NodeKinds                []string        `json:"node_kinds"`
	EdgeKinds                []string        `json:"edge_kinds"`
	PrivateSources           []PrivateSource `json:"private_sources"`
	AllowedEvidencePrefixes  []string        `json:"allowed_evidence_prefixes"`
	PublicSummaryFields      []string        `json:"public_summary_fields"`
	ForbiddenPublicFields    []string        `json:"forbidden_public_fields"`
	Commands                 []string        `json:"commands"`
}

type PrivateSource struct {
	Key      string `json:"key"`
	Path     string `json:"path"`
	NodeKind string `json:"node_kind"`
	Format   string `json:"format"`
}

type SourceStatus struct {
	Key      string `json:"key"`
	NodeKind string `json:"node_kind"`
	Format   string `json:"format"`
	Present  bool   `json:"present"`
	Count    int    `json:"count"`
}

type Status struct {
	PolicyPath               string         `json:"policy_path"`
	PrivateRoot              string         `json:"private_root"`
	SourceCount              int            `json:"source_count"`
	PresentSourceCount       int            `json:"present_source_count"`
	NodeCount                int            `json:"node_count"`
	EdgeCount                int            `json:"edge_count"`
	DanglingEvidenceRefCount int            `json:"dangling_evidence_ref_count"`
	OpenLearningCount        int            `json:"open_learning_count"`
	ByNodeKind               map[string]int `json:"by_node_kind"`
	ByEdgeKind               map[string]int `json:"by_edge_kind"`
	Sources                  []SourceStatus `json:"sources"`
	LastObservedAt           string         `json:"last_observed_at,omitempty"`
	CheckedAt                string         `json:"checked_at"`
}

type learningObservation struct {
	ID           string   `json:"id"`
	At           string   `json:"at"`
	Status       string   `json:"status"`
	EvidenceRefs []string `json:"evidence_refs"`
}

func StatusForRoot(root string) (Status, error) {
	policy, err := ReadPolicy(root)
	if err != nil {
		return Status{}, err
	}
	status := Status{
		PolicyPath:  PolicyRelativePath,
		PrivateRoot: policy.PrivateRoot,
		SourceCount: len(policy.PrivateSources),
		ByNodeKind:  map[string]int{},
		ByEdgeKind:  map[string]int{},
		CheckedAt:   time.Now().UTC().Format(time.RFC3339),
	}
	artifactRefs := map[string]bool{}
	for _, source := range policy.PrivateSources {
		sourceStatus, err := inspectSource(root, policy, source, &status, artifactRefs)
		if err != nil {
			return Status{}, err
		}
		if sourceStatus.Present {
			status.PresentSourceCount++
		}
		status.Sources = append(status.Sources, sourceStatus)
	}
	status.ByNodeKind["evidence_artifact"] += len(artifactRefs)
	status.NodeCount += len(artifactRefs)
	sort.Slice(status.Sources, func(i, j int) bool {
		return status.Sources[i].Key < status.Sources[j].Key
	})
	return status, nil
}

func ReadPolicy(root string) (Policy, error) {
	body, err := os.ReadFile(filepath.Join(root, filepath.FromSlash(PolicyRelativePath)))
	if err != nil {
		return Policy{}, err
	}
	var policy Policy
	if err := json.Unmarshal(body, &policy); err != nil {
		return Policy{}, err
	}
	if err := validatePolicy(policy); err != nil {
		return Policy{}, err
	}
	return policy, nil
}

func inspectSource(root string, policy Policy, source PrivateSource, status *Status, artifactRefs map[string]bool) (SourceStatus, error) {
	sourceStatus := SourceStatus{
		Key:      source.Key,
		NodeKind: source.NodeKind,
		Format:   source.Format,
	}
	switch source.Format {
	case "jsonl":
		if source.Key == "learning" {
			return inspectLearningSource(root, policy, source, status, artifactRefs)
		}
		count, present, last, err := countJSONL(root, source.Path)
		if err != nil {
			return sourceStatus, err
		}
		sourceStatus.Present = present
		sourceStatus.Count = count
		status.NodeCount += count
		status.ByNodeKind[source.NodeKind] += count
		updateLastObserved(status, last)
		return sourceStatus, nil
	case "directory":
		count, present, err := countDirectoryFiles(root, source.Path)
		if err != nil {
			return sourceStatus, err
		}
		sourceStatus.Present = present
		sourceStatus.Count = count
		status.NodeCount += count
		status.ByNodeKind[source.NodeKind] += count
		return sourceStatus, nil
	default:
		return sourceStatus, fmt.Errorf("evidence source %q has unsupported format", source.Key)
	}
}

func inspectLearningSource(root string, policy Policy, source PrivateSource, status *Status, artifactRefs map[string]bool) (SourceStatus, error) {
	sourceStatus := SourceStatus{
		Key:      source.Key,
		NodeKind: source.NodeKind,
		Format:   source.Format,
	}
	file, err := os.Open(filepath.Join(root, filepath.FromSlash(source.Path)))
	if errors.Is(err, os.ErrNotExist) {
		return sourceStatus, nil
	}
	if err != nil {
		return sourceStatus, err
	}
	defer file.Close()

	sourceStatus.Present = true
	scanner := bufio.NewScanner(file)
	scanner.Buffer(make([]byte, 0, 64*1024), 1024*1024)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}
		var observation learningObservation
		if err := json.Unmarshal([]byte(line), &observation); err != nil {
			return sourceStatus, err
		}
		sourceStatus.Count++
		status.NodeCount++
		status.ByNodeKind[source.NodeKind]++
		if observation.Status != "closed" {
			status.OpenLearningCount++
		}
		updateLastObserved(status, observation.At)
		for _, ref := range observation.EvidenceRefs {
			normalized, err := normalizeEvidenceRef(policy, ref)
			if err != nil {
				return sourceStatus, err
			}
			artifactRefs[normalized] = true
			status.EdgeCount++
			status.ByEdgeKind["supports"]++
			if _, err := os.Stat(filepath.Join(root, filepath.FromSlash(normalized))); errors.Is(err, os.ErrNotExist) {
				status.DanglingEvidenceRefCount++
			} else if err != nil {
				return sourceStatus, err
			}
		}
	}
	if err := scanner.Err(); err != nil {
		return sourceStatus, err
	}
	return sourceStatus, nil
}

func countJSONL(root string, rel string) (int, bool, string, error) {
	file, err := os.Open(filepath.Join(root, filepath.FromSlash(rel)))
	if errors.Is(err, os.ErrNotExist) {
		return 0, false, "", nil
	}
	if err != nil {
		return 0, false, "", err
	}
	defer file.Close()

	count := 0
	last := ""
	scanner := bufio.NewScanner(file)
	scanner.Buffer(make([]byte, 0, 64*1024), 1024*1024)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}
		count++
		var row struct {
			At string `json:"at"`
		}
		if err := json.Unmarshal([]byte(line), &row); err == nil {
			last = laterRFC3339(last, row.At)
		}
	}
	if err := scanner.Err(); err != nil {
		return 0, false, "", err
	}
	return count, true, last, nil
}

func countDirectoryFiles(root string, rel string) (int, bool, error) {
	path := filepath.Join(root, filepath.FromSlash(rel))
	info, err := os.Stat(path)
	if errors.Is(err, os.ErrNotExist) {
		return 0, false, nil
	}
	if err != nil {
		return 0, false, err
	}
	if !info.IsDir() {
		return 0, true, fmt.Errorf("evidence source is not a directory")
	}
	count := 0
	err = filepath.WalkDir(path, func(current string, entry fs.DirEntry, walkErr error) error {
		if walkErr != nil {
			return walkErr
		}
		if entry.Type().IsRegular() {
			count++
		}
		return nil
	})
	return count, true, err
}

func validatePolicy(policy Policy) error {
	if policy.Context != "AgentCluster" {
		return fmt.Errorf("evidence graph policy context = %q", policy.Context)
	}
	if policy.PrivateRoot != "data/private" {
		return fmt.Errorf("evidence graph private root must be data/private")
	}
	if !policy.PrivateGraphRequired || !policy.PublicStatusRedacted || policy.RawEvidencePublicAllowed {
		return fmt.Errorf("evidence graph policy must require private graph evidence and redacted public status")
	}
	nodeKinds := normalizeList(policy.NodeKinds)
	edgeKinds := normalizeList(policy.EdgeKinds)
	if !contains(nodeKinds, "learning_observation") || !contains(nodeKinds, "evidence_artifact") {
		return fmt.Errorf("evidence graph policy must include learning and artifact nodes")
	}
	if !contains(edgeKinds, "supports") {
		return fmt.Errorf("evidence graph policy must include supports edges")
	}
	if len(policy.PrivateSources) == 0 {
		return fmt.Errorf("evidence graph policy requires private sources")
	}
	for _, source := range policy.PrivateSources {
		if strings.TrimSpace(source.Key) == "" || strings.TrimSpace(source.Path) == "" {
			return fmt.Errorf("evidence graph source key and path are required")
		}
		if !strings.HasPrefix(source.Path, "data/private/") || filepath.IsAbs(filepath.FromSlash(source.Path)) || strings.Contains(source.Path, "..") {
			return fmt.Errorf("evidence graph source must stay repo-relative under data/private")
		}
		if !contains(nodeKinds, source.NodeKind) {
			return fmt.Errorf("evidence graph source %q has unknown node kind", source.Key)
		}
		if source.Format != "jsonl" && source.Format != "directory" {
			return fmt.Errorf("evidence graph source %q has unsupported format", source.Key)
		}
	}
	summaryFields := normalizeList(policy.PublicSummaryFields)
	for _, field := range []string{"node_count", "edge_count", "by_node_kind", "by_edge_kind", "checked_at"} {
		if !contains(summaryFields, field) {
			return fmt.Errorf("evidence graph public summary missing %q", field)
		}
	}
	if !contains(normalizeList(policy.AllowedEvidencePrefixes), "data/private/") {
		return fmt.Errorf("evidence graph evidence refs must allow data/private")
	}
	if !contains(policy.Commands, "mhj evidence status") {
		return fmt.Errorf("evidence graph status command is missing")
	}
	return nil
}

func normalizeEvidenceRef(policy Policy, ref string) (string, error) {
	normalized := filepath.ToSlash(strings.TrimSpace(ref))
	if normalized == "" || filepath.IsAbs(filepath.FromSlash(normalized)) || strings.Contains(normalized, "..") {
		return "", fmt.Errorf("evidence graph found invalid evidence ref")
	}
	for _, prefix := range policy.AllowedEvidencePrefixes {
		if strings.HasPrefix(normalized, prefix) {
			return normalized, nil
		}
	}
	return "", fmt.Errorf("evidence graph found evidence ref outside allowed prefixes")
}

func updateLastObserved(status *Status, candidate string) {
	status.LastObservedAt = laterRFC3339(status.LastObservedAt, candidate)
}

func laterRFC3339(current string, candidate string) string {
	candidate = strings.TrimSpace(candidate)
	if candidate == "" {
		return current
	}
	if current == "" {
		return candidate
	}
	currentTime, currentErr := time.Parse(time.RFC3339, current)
	candidateTime, candidateErr := time.Parse(time.RFC3339, candidate)
	if currentErr != nil || candidateErr != nil {
		if candidate > current {
			return candidate
		}
		return current
	}
	if candidateTime.After(currentTime) {
		return candidate
	}
	return current
}

func normalizeList(values []string) []string {
	seen := map[string]bool{}
	normalized := make([]string, 0, len(values))
	for _, value := range values {
		item := strings.TrimSpace(strings.ToLower(value))
		if item == "" || seen[item] {
			continue
		}
		seen[item] = true
		normalized = append(normalized, item)
	}
	sort.Strings(normalized)
	return normalized
}

func contains(values []string, wanted string) bool {
	for _, value := range values {
		if value == wanted {
			return true
		}
	}
	return false
}
