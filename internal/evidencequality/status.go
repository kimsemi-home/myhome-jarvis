package evidencequality

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"time"
)

const PolicyRelativePath = "generated/evidence_quality.generated.json"

type Policy struct {
	Context                  string   `json:"context"`
	Version                  string   `json:"version"`
	GeneratedArtifact        string   `json:"generated_artifact"`
	PrivateSnapshotLedger    string   `json:"private_snapshot_ledger"`
	AppendOnly               bool     `json:"append_only"`
	PublicStatusRedacted     bool     `json:"public_status_redacted"`
	RawSnapshotPublicAllowed bool     `json:"raw_snapshot_public_allowed"`
	StaleAfterHours          int      `json:"stale_after_hours"`
	QualityLevels            []string `json:"quality_levels"`
	MappingConfidenceLevels  []string `json:"mapping_confidence_levels"`
	AllowedPurposes          []string `json:"allowed_purposes"`
	ReassessmentReasons      []string `json:"reassessment_reasons"`
	RequiredFields           []string `json:"required_fields"`
	AllowedEvidencePrefixes  []string `json:"allowed_evidence_prefixes"`
	PublicSummaryFields      []string `json:"public_summary_fields"`
	ForbiddenPublicFields    []string `json:"forbidden_public_fields"`
	Commands                 []string `json:"commands"`
}

type Snapshot struct {
	ID                  string   `json:"id"`
	At                  string   `json:"at"`
	EvidenceRef         string   `json:"evidence_ref"`
	Purpose             string   `json:"purpose"`
	QualityLevel        string   `json:"quality_level"`
	SchemaVersion       string   `json:"schema_version"`
	OntologyVersion     string   `json:"ontology_version"`
	MappingConfidence   string   `json:"mapping_confidence"`
	AssessedBy          string   `json:"assessed_by"`
	ReassessmentReasons []string `json:"reassessment_reasons"`
}

type Status struct {
	PolicyPath            string         `json:"policy_path"`
	LedgerPath            string         `json:"ledger_path"`
	Exists                bool           `json:"exists"`
	SnapshotCount         int            `json:"snapshot_count"`
	InvalidSnapshotCount  int            `json:"invalid_snapshot_count"`
	ReassessmentDebtCount int            `json:"reassessment_debt_count"`
	MissingEvidenceCount  int            `json:"missing_evidence_count"`
	StaleSnapshotCount    int            `json:"stale_snapshot_count"`
	LowQualityCount       int            `json:"low_quality_count"`
	BlockedQualityCount   int            `json:"blocked_quality_count"`
	MappingDriftCount     int            `json:"mapping_drift_count"`
	StaleAfterHours       int            `json:"stale_after_hours"`
	ByQualityLevel        map[string]int `json:"by_quality_level"`
	ByMappingConfidence   map[string]int `json:"by_mapping_confidence"`
	ByPurpose             map[string]int `json:"by_purpose"`
	ByReassessmentReason  map[string]int `json:"by_reassessment_reason"`
	LastObservedAt        string         `json:"last_observed_at,omitempty"`
	CheckedAt             string         `json:"checked_at"`
}

var errMissingEvidenceRef = errors.New("evidence quality snapshot evidence ref is required")

func StatusForRoot(root string) (Status, error) {
	policy, err := ReadPolicy(root)
	if err != nil {
		return Status{}, err
	}
	checkedAt := time.Now().UTC()
	status := Status{
		PolicyPath:           PolicyRelativePath,
		LedgerPath:           policy.PrivateSnapshotLedger,
		StaleAfterHours:      policy.StaleAfterHours,
		ByQualityLevel:       map[string]int{},
		ByMappingConfidence:  map[string]int{},
		ByPurpose:            map[string]int{},
		ByReassessmentReason: map[string]int{},
		CheckedAt:            checkedAt.Format(time.RFC3339),
	}
	file, err := os.Open(filepath.Join(root, filepath.FromSlash(policy.PrivateSnapshotLedger)))
	if errors.Is(err, os.ErrNotExist) {
		return status, nil
	}
	if err != nil {
		return Status{}, err
	}
	defer file.Close()

	status.Exists = true
	scanner := bufio.NewScanner(file)
	scanner.Buffer(make([]byte, 0, 64*1024), 1024*1024)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}
		var snapshot Snapshot
		if err := json.Unmarshal([]byte(line), &snapshot); err != nil {
			status.InvalidSnapshotCount++
			continue
		}
		normalized, err := normalizeSnapshot(policy, snapshot)
		if err != nil {
			if errors.Is(err, errMissingEvidenceRef) {
				status.MissingEvidenceCount++
			} else {
				status.InvalidSnapshotCount++
			}
			continue
		}
		status.SnapshotCount++
		status.ByQualityLevel[normalized.QualityLevel]++
		status.ByMappingConfidence[normalized.MappingConfidence]++
		status.ByPurpose[normalized.Purpose]++
		for _, reason := range normalized.ReassessmentReasons {
			status.ByReassessmentReason[reason]++
		}
		if isStale(policy, normalized, checkedAt) {
			status.StaleSnapshotCount++
		}
		if normalized.QualityLevel == "low" {
			status.LowQualityCount++
		}
		if normalized.QualityLevel == "blocked" {
			status.BlockedQualityCount++
		}
		if normalized.MappingConfidence == "low" || normalized.MappingConfidence == "unknown" {
			status.MappingDriftCount++
		}
		status.LastObservedAt = laterRFC3339(status.LastObservedAt, normalized.At)
	}
	if err := scanner.Err(); err != nil {
		return Status{}, err
	}
	status.ReassessmentDebtCount = status.InvalidSnapshotCount +
		status.MissingEvidenceCount +
		status.StaleSnapshotCount +
		status.LowQualityCount +
		status.BlockedQualityCount +
		status.MappingDriftCount
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

func validatePolicy(policy Policy) error {
	if policy.Context != "AgentCluster" {
		return fmt.Errorf("evidence quality policy context = %q", policy.Context)
	}
	if !strings.HasPrefix(policy.PrivateSnapshotLedger, "data/private/") || !strings.HasSuffix(policy.PrivateSnapshotLedger, ".jsonl") {
		return fmt.Errorf("evidence quality ledger must stay in data/private JSONL")
	}
	if !policy.AppendOnly || !policy.PublicStatusRedacted || policy.RawSnapshotPublicAllowed {
		return fmt.Errorf("evidence quality policy must be private append-only with redacted public status")
	}
	if policy.StaleAfterHours <= 0 {
		return fmt.Errorf("evidence quality stale threshold must be positive")
	}
	levels := normalizeList(policy.QualityLevels)
	for _, level := range []string{"high", "medium", "low", "blocked"} {
		if !contains(levels, level) {
			return fmt.Errorf("evidence quality level %q is missing", level)
		}
	}
	mappingLevels := normalizeList(policy.MappingConfidenceLevels)
	for _, level := range []string{"high", "medium", "low", "unknown"} {
		if !contains(mappingLevels, level) {
			return fmt.Errorf("evidence quality mapping confidence %q is missing", level)
		}
	}
	purposes := normalizeList(policy.AllowedPurposes)
	for _, purpose := range []string{"root_cause", "confidence_assessment", "incident_review", "release_gate", "conformance", "revalidation"} {
		if !contains(purposes, purpose) {
			return fmt.Errorf("evidence quality purpose %q is missing", purpose)
		}
	}
	reasons := normalizeList(policy.ReassessmentReasons)
	for _, reason := range []string{"age", "schema_version_change", "ontology_version_change", "counter_evidence", "security_incident", "quarantine", "translation_loss"} {
		if !contains(reasons, reason) {
			return fmt.Errorf("evidence quality reassessment reason %q is missing", reason)
		}
	}
	required := normalizeList(policy.RequiredFields)
	for _, field := range []string{"at", "evidence_ref", "purpose", "quality_level", "schema_version", "ontology_version", "mapping_confidence", "assessed_by", "reassessment_reasons"} {
		if !contains(required, field) {
			return fmt.Errorf("evidence quality required field %q is missing", field)
		}
	}
	summary := normalizeList(policy.PublicSummaryFields)
	for _, field := range []string{"snapshot_count", "invalid_snapshot_count", "reassessment_debt_count", "missing_evidence_count", "stale_snapshot_count", "low_quality_count", "blocked_quality_count", "mapping_drift_count", "by_quality_level", "by_mapping_confidence", "checked_at"} {
		if !contains(summary, field) {
			return fmt.Errorf("evidence quality public summary missing %q", field)
		}
	}
	if !contains(policy.Commands, "mhj evidence-quality status") {
		return fmt.Errorf("evidence quality status command is missing")
	}
	return nil
}

func normalizeSnapshot(policy Policy, snapshot Snapshot) (Snapshot, error) {
	normalized := Snapshot{
		ID:                  publicText(snapshot.ID),
		At:                  publicText(snapshot.At),
		EvidenceRef:         filepath.ToSlash(strings.TrimSpace(snapshot.EvidenceRef)),
		Purpose:             normalizeToken(snapshot.Purpose),
		QualityLevel:        normalizeToken(snapshot.QualityLevel),
		SchemaVersion:       publicText(snapshot.SchemaVersion),
		OntologyVersion:     publicText(snapshot.OntologyVersion),
		MappingConfidence:   normalizeToken(snapshot.MappingConfidence),
		AssessedBy:          normalizeToken(snapshot.AssessedBy),
		ReassessmentReasons: normalizeList(snapshot.ReassessmentReasons),
	}
	if normalized.At == "" {
		return Snapshot{}, fmt.Errorf("evidence quality snapshot at timestamp is required")
	}
	if _, err := time.Parse(time.RFC3339, normalized.At); err != nil {
		return Snapshot{}, fmt.Errorf("evidence quality snapshot at timestamp is invalid")
	}
	if normalized.EvidenceRef == "" {
		return Snapshot{}, errMissingEvidenceRef
	}
	if err := validateRef(policy, normalized.EvidenceRef); err != nil {
		return Snapshot{}, err
	}
	if err := rejectSensitiveText(normalized.EvidenceRef); err != nil {
		return Snapshot{}, err
	}
	if !contains(normalizeList(policy.AllowedPurposes), normalized.Purpose) {
		return Snapshot{}, fmt.Errorf("evidence quality purpose %q is not allowed", normalized.Purpose)
	}
	if !contains(normalizeList(policy.QualityLevels), normalized.QualityLevel) {
		return Snapshot{}, fmt.Errorf("evidence quality level %q is not allowed", normalized.QualityLevel)
	}
	if !contains(normalizeList(policy.MappingConfidenceLevels), normalized.MappingConfidence) {
		return Snapshot{}, fmt.Errorf("evidence quality mapping confidence %q is not allowed", normalized.MappingConfidence)
	}
	if normalized.SchemaVersion == "" || normalized.OntologyVersion == "" || normalized.AssessedBy == "" {
		return Snapshot{}, fmt.Errorf("evidence quality schema, ontology, and assessor fields are required")
	}
	for _, reason := range normalized.ReassessmentReasons {
		if !contains(normalizeList(policy.ReassessmentReasons), reason) {
			return Snapshot{}, fmt.Errorf("evidence quality reassessment reason %q is not allowed", reason)
		}
	}
	return normalized, nil
}

func validateRef(policy Policy, ref string) error {
	ref = filepath.ToSlash(strings.TrimSpace(ref))
	if ref == "" {
		return errMissingEvidenceRef
	}
	if filepath.IsAbs(filepath.FromSlash(ref)) || strings.Contains(ref, "..") {
		return fmt.Errorf("evidence quality ref must be repo-relative")
	}
	for _, prefix := range policy.AllowedEvidencePrefixes {
		if strings.HasPrefix(ref, prefix) {
			return nil
		}
	}
	return fmt.Errorf("evidence quality ref is outside allowed prefixes")
}

func isStale(policy Policy, snapshot Snapshot, checkedAt time.Time) bool {
	at, err := time.Parse(time.RFC3339, snapshot.At)
	if err != nil {
		return false
	}
	return checkedAt.Sub(at) > time.Duration(policy.StaleAfterHours)*time.Hour
}

func rejectSensitiveText(value string) error {
	lower := strings.ToLower(value)
	for _, marker := range []string{
		"kim" + "jooyoon",
		"kim-joo" + "-yoon",
		"/us" + "ers/" + "al" + "ice",
		"al" + "ice/" + "git" + "hub",
		"bearer ",
		"begin private key",
		"raw_prompt",
		"raw_transcript",
		"account_id",
		"card_number",
		"api_secret",
		"credential=",
		"linear." + "app/",
		string(filepath.Separator) + "users" + string(filepath.Separator),
		"\\" + "users" + "\\",
	} {
		if strings.Contains(lower, marker) {
			return fmt.Errorf("evidence quality ref contains forbidden private marker")
		}
	}
	return nil
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

func publicText(value string) string {
	return strings.Join(strings.Fields(strings.TrimSpace(value)), " ")
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

func contains(values []string, wanted string) bool {
	for _, value := range values {
		if value == wanted {
			return true
		}
	}
	return false
}
