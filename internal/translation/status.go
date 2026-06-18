package translation

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

const PolicyRelativePath = "generated/translation.generated.json"

type Policy struct {
	Context                 string   `json:"context"`
	Version                 string   `json:"version"`
	GeneratedArtifact       string   `json:"generated_artifact"`
	PrivateLossLedger       string   `json:"private_loss_ledger"`
	PrivateManifestRoot     string   `json:"private_manifest_root"`
	ManifestRequired        bool     `json:"manifest_required"`
	PublicStatusRedacted    bool     `json:"public_status_redacted"`
	RawLossPublicAllowed    bool     `json:"raw_loss_public_allowed"`
	AllowedContexts         []string `json:"allowed_contexts"`
	RequiredManifestFields  []string `json:"required_manifest_fields"`
	LossLevels              []string `json:"loss_levels"`
	AllowedLossCategories   []string `json:"allowed_loss_categories"`
	ForbiddenLossCategories []string `json:"forbidden_loss_categories"`
	AllowedEvidencePrefixes []string `json:"allowed_evidence_prefixes"`
	PublicSummaryFields     []string `json:"public_summary_fields"`
	ForbiddenPublicFields   []string `json:"forbidden_public_fields"`
	Commands                []string `json:"commands"`
}

type Status struct {
	PolicyPath           string         `json:"policy_path"`
	LedgerPath           string         `json:"ledger_path"`
	ManifestRoot         string         `json:"manifest_root"`
	LedgerExists         bool           `json:"ledger_exists"`
	ManifestRootExists   bool           `json:"manifest_root_exists"`
	ManifestCount        int            `json:"manifest_count"`
	InvalidManifestCount int            `json:"invalid_manifest_count"`
	MissingManifestCount int            `json:"missing_manifest_count"`
	LossCount            int            `json:"loss_count"`
	OpenLossCount        int            `json:"open_loss_count"`
	ClosedLossCount      int            `json:"closed_loss_count"`
	InvalidLossCount     int            `json:"invalid_loss_count"`
	OpenDebtCount        int            `json:"open_debt_count"`
	ForbiddenLossCount   int            `json:"forbidden_loss_count"`
	ByLevel              map[string]int `json:"by_level"`
	BySourceContext      map[string]int `json:"by_source_context"`
	ByTargetContext      map[string]int `json:"by_target_context"`
	LastObservedAt       string         `json:"last_observed_at,omitempty"`
	CheckedAt            string         `json:"checked_at"`
}

type lossRecord struct {
	At            string   `json:"at"`
	SourceContext string   `json:"source_context"`
	TargetContext string   `json:"target_context"`
	Level         string   `json:"level"`
	Category      string   `json:"category"`
	Status        string   `json:"status"`
	ManifestPath  string   `json:"manifest_path"`
	EvidenceRefs  []string `json:"evidence_refs"`
}

type manifest struct {
	SourceContext  string      `json:"source_context"`
	TargetContext  string      `json:"target_context"`
	SourceVersion  string      `json:"source_version"`
	TargetVersion  string      `json:"target_version"`
	PreservedRules []string    `json:"preserved_rules"`
	KnownLosses    []knownLoss `json:"known_losses"`
	Owner          string      `json:"owner"`
	EvidenceRefs   []string    `json:"evidence_refs"`
}

type knownLoss struct {
	Level    string `json:"level"`
	Category string `json:"category"`
}

func StatusForRoot(root string) (Status, error) {
	policy, err := ReadPolicy(root)
	if err != nil {
		return Status{}, err
	}
	status := Status{
		PolicyPath:      PolicyRelativePath,
		LedgerPath:      policy.PrivateLossLedger,
		ManifestRoot:    policy.PrivateManifestRoot,
		ByLevel:         map[string]int{},
		BySourceContext: map[string]int{},
		ByTargetContext: map[string]int{},
		CheckedAt:       time.Now().UTC().Format(time.RFC3339),
	}
	if err := inspectLedger(root, policy, &status); err != nil {
		return Status{}, err
	}
	if err := inspectManifests(root, policy, &status); err != nil {
		return Status{}, err
	}
	status.OpenDebtCount = status.OpenLossCount + status.InvalidLossCount + status.InvalidManifestCount + status.MissingManifestCount
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

func inspectLedger(root string, policy Policy, status *Status) error {
	file, err := os.Open(filepath.Join(root, filepath.FromSlash(policy.PrivateLossLedger)))
	if errors.Is(err, os.ErrNotExist) {
		return nil
	}
	if err != nil {
		return err
	}
	defer file.Close()

	status.LedgerExists = true
	scanner := bufio.NewScanner(file)
	scanner.Buffer(make([]byte, 0, 64*1024), 1024*1024)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}
		var record lossRecord
		if err := json.Unmarshal([]byte(line), &record); err != nil {
			status.InvalidLossCount++
			continue
		}
		status.LossCount++
		updateLastObserved(status, record.At)
		if err := recordLoss(policy, status, record.SourceContext, record.TargetContext, record.Level, record.Category, record.Status); err != nil {
			status.InvalidLossCount++
		}
		if policy.ManifestRequired {
			if err := validateReferencedManifest(root, policy, record.ManifestPath); err != nil {
				status.MissingManifestCount++
			}
		}
	}
	return scanner.Err()
}

func inspectManifests(root string, policy Policy, status *Status) error {
	manifestRoot := filepath.Join(root, filepath.FromSlash(policy.PrivateManifestRoot))
	info, err := os.Stat(manifestRoot)
	if errors.Is(err, os.ErrNotExist) {
		return nil
	}
	if err != nil {
		return err
	}
	status.ManifestRootExists = true
	if !info.IsDir() {
		status.InvalidManifestCount++
		return nil
	}
	return filepath.WalkDir(manifestRoot, func(path string, entry fs.DirEntry, walkErr error) error {
		if walkErr != nil {
			return walkErr
		}
		if entry.IsDir() {
			return nil
		}
		if filepath.Ext(path) != ".json" {
			return nil
		}
		status.ManifestCount++
		body, err := os.ReadFile(path)
		if err != nil {
			return err
		}
		var current manifest
		if err := json.Unmarshal(body, &current); err != nil {
			status.InvalidManifestCount++
			return nil
		}
		if err := validateManifest(policy, current); err != nil {
			status.InvalidManifestCount++
			return nil
		}
		recordContext(status, current.SourceContext, current.TargetContext)
		for _, loss := range current.KnownLosses {
			status.LossCount++
			lossStatus := "closed"
			if normalizeToken(loss.Level) != "l0_none" {
				lossStatus = "open"
			}
			if err := recordLoss(policy, status, current.SourceContext, current.TargetContext, loss.Level, loss.Category, lossStatus); err != nil {
				status.InvalidManifestCount++
				return nil
			}
		}
		return nil
	})
}

func validatePolicy(policy Policy) error {
	if policy.Context != "AgentCluster" {
		return fmt.Errorf("translation policy context = %q", policy.Context)
	}
	if !strings.HasPrefix(policy.PrivateLossLedger, "data/private/") || !strings.HasSuffix(policy.PrivateLossLedger, ".jsonl") {
		return fmt.Errorf("translation loss ledger must stay in data/private JSONL")
	}
	if !strings.HasPrefix(policy.PrivateManifestRoot, "data/private/") {
		return fmt.Errorf("translation manifest root must stay in data/private")
	}
	if !policy.ManifestRequired || !policy.PublicStatusRedacted || policy.RawLossPublicAllowed {
		return fmt.Errorf("translation policy must require private manifests and redacted public status")
	}
	for _, context := range []string{"AgentCluster", "KnowledgeIndex", "AgentOps", "SecurityPolicy"} {
		if !contains(policy.AllowedContexts, context) {
			return fmt.Errorf("translation policy missing context %q", context)
		}
	}
	fields := normalizeList(policy.RequiredManifestFields)
	for _, field := range []string{"source_context", "target_context", "source_version", "target_version", "preserved_rules", "known_losses", "owner", "evidence_refs"} {
		if !contains(fields, field) {
			return fmt.Errorf("translation manifest required field %q is missing", field)
		}
	}
	levels := normalizeList(policy.LossLevels)
	for _, level := range []string{"l0_none", "l1_note", "l2_degraded", "l3_review_required", "l4_forbidden"} {
		if !contains(levels, level) {
			return fmt.Errorf("translation loss level %q is missing", level)
		}
	}
	categories := normalizeList(policy.AllowedLossCategories)
	for _, category := range []string{"mapping_gap", "authority", "security_boundary", "financial_commitment"} {
		if !contains(categories, category) {
			return fmt.Errorf("translation loss category %q is missing", category)
		}
	}
	summary := normalizeList(policy.PublicSummaryFields)
	for _, field := range []string{"open_debt_count", "forbidden_loss_count", "invalid_manifest_count", "missing_manifest_count", "checked_at"} {
		if !contains(summary, field) {
			return fmt.Errorf("translation public summary missing %q", field)
		}
	}
	if !contains(policy.Commands, "mhj translation status") {
		return fmt.Errorf("translation status command is missing")
	}
	return nil
}

func validateManifest(policy Policy, current manifest) error {
	if !contains(policy.AllowedContexts, strings.TrimSpace(current.SourceContext)) {
		return fmt.Errorf("source context is not allowed")
	}
	if !contains(policy.AllowedContexts, strings.TrimSpace(current.TargetContext)) {
		return fmt.Errorf("target context is not allowed")
	}
	if strings.TrimSpace(current.SourceVersion) == "" || strings.TrimSpace(current.TargetVersion) == "" {
		return fmt.Errorf("translation manifest versions are required")
	}
	if len(current.PreservedRules) == 0 || len(current.KnownLosses) == 0 || strings.TrimSpace(current.Owner) == "" || len(current.EvidenceRefs) == 0 {
		return fmt.Errorf("translation manifest required fields are incomplete")
	}
	for _, ref := range current.EvidenceRefs {
		if err := validateEvidenceRef(policy, ref); err != nil {
			return err
		}
	}
	return nil
}

func validateReferencedManifest(root string, policy Policy, manifestPath string) error {
	manifestPath = filepath.ToSlash(strings.TrimSpace(manifestPath))
	if manifestPath == "" {
		return fmt.Errorf("translation manifest path is required")
	}
	if filepath.IsAbs(filepath.FromSlash(manifestPath)) || strings.Contains(manifestPath, "..") {
		return fmt.Errorf("translation manifest path must be repo-relative")
	}
	if !strings.HasPrefix(manifestPath, strings.TrimRight(policy.PrivateManifestRoot, "/")+"/") {
		return fmt.Errorf("translation manifest path must stay under private manifest root")
	}
	if _, err := os.Stat(filepath.Join(root, filepath.FromSlash(manifestPath))); err != nil {
		return err
	}
	return nil
}

func recordLoss(policy Policy, status *Status, sourceContext string, targetContext string, level string, category string, lossStatus string) error {
	sourceContext = strings.TrimSpace(sourceContext)
	targetContext = strings.TrimSpace(targetContext)
	level = normalizeToken(level)
	category = normalizeToken(category)
	lossStatus = normalizeToken(lossStatus)
	if lossStatus == "" {
		lossStatus = "open"
	}
	if !contains(policy.AllowedContexts, sourceContext) || !contains(policy.AllowedContexts, targetContext) {
		return fmt.Errorf("translation loss context is not allowed")
	}
	if !contains(normalizeList(policy.LossLevels), level) || !contains(normalizeList(policy.AllowedLossCategories), category) {
		return fmt.Errorf("translation loss level or category is not allowed")
	}
	recordContext(status, sourceContext, targetContext)
	status.ByLevel[level]++
	if lossStatus == "closed" {
		status.ClosedLossCount++
	} else {
		status.OpenLossCount++
	}
	if level == "l4_forbidden" || contains(normalizeList(policy.ForbiddenLossCategories), category) {
		status.ForbiddenLossCount++
	}
	return nil
}

func recordContext(status *Status, sourceContext string, targetContext string) {
	sourceContext = strings.TrimSpace(sourceContext)
	targetContext = strings.TrimSpace(targetContext)
	if sourceContext != "" {
		status.BySourceContext[sourceContext]++
	}
	if targetContext != "" {
		status.ByTargetContext[targetContext]++
	}
}

func validateEvidenceRef(policy Policy, ref string) error {
	ref = filepath.ToSlash(strings.TrimSpace(ref))
	if ref == "" {
		return fmt.Errorf("translation evidence ref is required")
	}
	if filepath.IsAbs(filepath.FromSlash(ref)) || strings.Contains(ref, "..") {
		return fmt.Errorf("translation evidence ref must be repo-relative")
	}
	for _, prefix := range policy.AllowedEvidencePrefixes {
		if strings.HasPrefix(ref, prefix) {
			return nil
		}
	}
	return fmt.Errorf("translation evidence ref is outside allowed prefixes")
}

func updateLastObserved(status *Status, candidate string) {
	status.LastObservedAt = laterRFC3339(status.LastObservedAt, candidate)
}

func laterRFC3339(left string, right string) string {
	if strings.TrimSpace(right) == "" {
		return left
	}
	if strings.TrimSpace(left) == "" {
		return right
	}
	leftTime, leftErr := time.Parse(time.RFC3339, left)
	rightTime, rightErr := time.Parse(time.RFC3339, right)
	if leftErr != nil || rightErr != nil {
		return right
	}
	if rightTime.After(leftTime) {
		return right
	}
	return left
}

func normalizeList(values []string) []string {
	seen := map[string]bool{}
	normalized := make([]string, 0, len(values))
	for _, value := range values {
		item := normalizeToken(value)
		if item == "" || seen[item] {
			continue
		}
		seen[item] = true
		normalized = append(normalized, item)
	}
	sort.Strings(normalized)
	return normalized
}

func normalizeToken(value string) string {
	return strings.TrimSpace(strings.ToLower(value))
}

func contains(values []string, wanted string) bool {
	for _, value := range values {
		if value == wanted {
			return true
		}
	}
	return false
}
