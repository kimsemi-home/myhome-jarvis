package learning

import (
	"bufio"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"time"
)

const PolicyRelativePath = "generated/learning.generated.json"

type Policy struct {
	Context                     string   `json:"context"`
	Version                     string   `json:"version"`
	PrivateLedger               string   `json:"private_ledger"`
	GeneratedArtifact           string   `json:"generated_artifact"`
	AppendOnly                  bool     `json:"append_only"`
	PrivateJournalRequired      bool     `json:"private_journal_required"`
	PublicStatusRedacted        bool     `json:"public_status_redacted"`
	RawObservationPublicAllowed bool     `json:"raw_observation_public_allowed"`
	RequiredFields              []string `json:"required_fields"`
	AllowedKinds                []string `json:"allowed_kinds"`
	Lifecycle                   []string `json:"lifecycle"`
	AllowedStatuses             []string `json:"allowed_statuses"`
	AllowedEvidencePrefixes     []string `json:"allowed_evidence_prefixes"`
	PublicSummaryFields         []string `json:"public_summary_fields"`
	ForbiddenPrivateMarkers     []string `json:"forbidden_private_markers"`
	Commands                    []string `json:"commands"`
}

type RecordRequest struct {
	Kind         string   `json:"kind"`
	Source       string   `json:"source"`
	Stage        string   `json:"stage,omitempty"`
	Status       string   `json:"status,omitempty"`
	Summary      string   `json:"summary"`
	EvidenceRefs []string `json:"evidence_refs"`
	Owner        string   `json:"owner"`
	NextAction   string   `json:"next_action"`
}

type Observation struct {
	ID           string   `json:"id"`
	At           string   `json:"at"`
	Kind         string   `json:"kind"`
	Source       string   `json:"source"`
	Stage        string   `json:"stage"`
	Status       string   `json:"status"`
	Summary      string   `json:"summary"`
	EvidenceRefs []string `json:"evidence_refs"`
	Owner        string   `json:"owner"`
	NextAction   string   `json:"next_action"`
}

type RecordResult struct {
	ID               string `json:"id"`
	Path             string `json:"path"`
	Kind             string `json:"kind"`
	Stage            string `json:"stage"`
	Status           string `json:"status"`
	EvidenceRefCount int    `json:"evidence_ref_count"`
	RecordedAt       string `json:"recorded_at"`
}

type Status struct {
	Path           string         `json:"path"`
	PolicyPath     string         `json:"policy_path"`
	Exists         bool           `json:"exists"`
	Count          int            `json:"count"`
	OpenCount      int            `json:"open_count"`
	ClosedCount    int            `json:"closed_count"`
	ByKind         map[string]int `json:"by_kind"`
	ByStage        map[string]int `json:"by_stage"`
	LastKind       string         `json:"last_kind,omitempty"`
	LastStage      string         `json:"last_stage,omitempty"`
	LastStatus     string         `json:"last_status,omitempty"`
	LastObservedAt string         `json:"last_observed_at,omitempty"`
	CheckedAt      string         `json:"checked_at"`
}

func Record(root string, payload []byte) (RecordResult, error) {
	policy, err := ReadPolicy(root)
	if err != nil {
		return RecordResult{}, err
	}
	var request RecordRequest
	decoder := json.NewDecoder(strings.NewReader(string(payload)))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&request); err != nil {
		return RecordResult{}, fmt.Errorf("invalid learning record payload: %w", err)
	}
	observation, err := normalizeObservation(policy, request)
	if err != nil {
		return RecordResult{}, err
	}
	if err := appendObservation(root, policy, observation); err != nil {
		return RecordResult{}, err
	}
	return RecordResult{
		ID:               observation.ID,
		Path:             policy.PrivateLedger,
		Kind:             observation.Kind,
		Stage:            observation.Stage,
		Status:           observation.Status,
		EvidenceRefCount: len(observation.EvidenceRefs),
		RecordedAt:       observation.At,
	}, nil
}

func StatusForRoot(root string) (Status, error) {
	policy, err := ReadPolicy(root)
	if err != nil {
		return Status{}, err
	}
	status := Status{
		Path:       policy.PrivateLedger,
		PolicyPath: PolicyRelativePath,
		ByKind:     map[string]int{},
		ByStage:    map[string]int{},
		CheckedAt:  time.Now().UTC().Format(time.RFC3339),
	}
	file, err := os.Open(filepath.Join(root, filepath.FromSlash(policy.PrivateLedger)))
	if errors.Is(err, os.ErrNotExist) {
		return status, nil
	}
	if err != nil {
		return status, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Buffer(make([]byte, 0, 64*1024), 1024*1024)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}
		var observation Observation
		if err := json.Unmarshal([]byte(line), &observation); err != nil {
			return status, err
		}
		status.Exists = true
		status.Count++
		status.ByKind[observation.Kind]++
		status.ByStage[observation.Stage]++
		if observation.Status == "closed" {
			status.ClosedCount++
		} else {
			status.OpenCount++
		}
		status.LastKind = observation.Kind
		status.LastStage = observation.Stage
		status.LastStatus = observation.Status
		status.LastObservedAt = observation.At
	}
	if err := scanner.Err(); err != nil {
		return status, err
	}
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
		return fmt.Errorf("learning policy context = %q", policy.Context)
	}
	if !strings.HasPrefix(policy.PrivateLedger, "data/private/") || !strings.HasSuffix(policy.PrivateLedger, ".jsonl") {
		return fmt.Errorf("learning ledger must stay in data/private JSONL")
	}
	if !policy.AppendOnly || !policy.PrivateJournalRequired || !policy.PublicStatusRedacted || policy.RawObservationPublicAllowed {
		return fmt.Errorf("learning policy must be private append-only with redacted public status")
	}
	for _, required := range []string{"kind", "source", "summary", "evidence_refs", "owner", "next_action"} {
		if !contains(normalizeList(policy.RequiredFields), required) {
			return fmt.Errorf("learning policy missing required field %q", required)
		}
	}
	if !contains(normalizeList(policy.AllowedKinds), "loop_gap") || !contains(normalizeList(policy.AllowedKinds), "evidence_debt") {
		return fmt.Errorf("learning policy must track loop_gap and evidence_debt")
	}
	if !contains(normalizeList(policy.Commands), "mhj learning status") || !contains(normalizeList(policy.Commands), "mhj learning record <json-payload>") {
		return fmt.Errorf("learning policy commands are incomplete")
	}
	return nil
}

func normalizeObservation(policy Policy, request RecordRequest) (Observation, error) {
	observation := Observation{
		At:           time.Now().UTC().Format(time.RFC3339),
		Kind:         normalizeToken(request.Kind),
		Source:       normalizeToken(request.Source),
		Stage:        normalizeToken(request.Stage),
		Status:       normalizeToken(request.Status),
		Summary:      publicText(request.Summary),
		EvidenceRefs: normalizeRefs(request.EvidenceRefs),
		Owner:        normalizeToken(request.Owner),
		NextAction:   publicText(request.NextAction),
	}
	if observation.Stage == "" {
		observation.Stage = "evidence_recorded"
	}
	if observation.Status == "" {
		observation.Status = "open"
	}
	if !contains(normalizeList(policy.AllowedKinds), observation.Kind) {
		return Observation{}, fmt.Errorf("learning kind %q is not allowed", observation.Kind)
	}
	if !contains(normalizeList(policy.Lifecycle), observation.Stage) {
		return Observation{}, fmt.Errorf("learning lifecycle stage %q is not allowed", observation.Stage)
	}
	if !contains(normalizeList(policy.AllowedStatuses), observation.Status) {
		return Observation{}, fmt.Errorf("learning status %q is not allowed", observation.Status)
	}
	if observation.Source == "" || observation.Summary == "" || observation.Owner == "" || observation.NextAction == "" {
		return Observation{}, fmt.Errorf("learning record requires source, summary, owner, and next_action")
	}
	if len(observation.EvidenceRefs) == 0 {
		return Observation{}, fmt.Errorf("learning record requires at least one evidence ref")
	}
	for _, ref := range observation.EvidenceRefs {
		if err := validateEvidenceRef(policy, ref); err != nil {
			return Observation{}, err
		}
	}
	for _, value := range []string{observation.Source, observation.Summary, observation.Owner, observation.NextAction} {
		if err := rejectSensitiveText(value); err != nil {
			return Observation{}, err
		}
	}
	observation.ID = observationID(observation)
	return observation, nil
}

func appendObservation(root string, policy Policy, observation Observation) error {
	if strings.TrimSpace(root) == "" {
		return errors.New("root is required")
	}
	path := filepath.Join(root, filepath.FromSlash(policy.PrivateLedger))
	if err := os.MkdirAll(filepath.Dir(path), 0o700); err != nil {
		return err
	}
	file, err := os.OpenFile(path, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0o600)
	if err != nil {
		return err
	}
	defer file.Close()
	data, err := json.Marshal(observation)
	if err != nil {
		return err
	}
	data = append(data, '\n')
	_, err = file.Write(data)
	return err
}

func validateEvidenceRef(policy Policy, ref string) error {
	if ref == "" {
		return fmt.Errorf("learning evidence ref is required")
	}
	if filepath.IsAbs(filepath.FromSlash(ref)) || strings.Contains(ref, "..") {
		return fmt.Errorf("learning evidence ref %q must be repo-relative", ref)
	}
	for _, prefix := range policy.AllowedEvidencePrefixes {
		if strings.HasPrefix(ref, prefix) {
			return nil
		}
	}
	return fmt.Errorf("learning evidence ref %q is outside allowed prefixes", ref)
}

func rejectSensitiveText(value string) error {
	lower := strings.ToLower(value)
	for _, marker := range []string{
		"bearer ",
		"begin private key",
		"raw_prompt",
		"raw_transcript",
		"account_id",
		"card_number",
		"api_secret",
		"credential=",
		string(filepath.Separator) + "users" + string(filepath.Separator),
		"\\" + "users" + "\\",
	} {
		if strings.Contains(lower, marker) {
			return fmt.Errorf("learning record contains forbidden private marker")
		}
	}
	return nil
}

func observationID(observation Observation) string {
	sum := sha256.Sum256([]byte(strings.Join([]string{
		observation.At,
		observation.Kind,
		observation.Source,
		observation.Stage,
		observation.Owner,
		observation.Summary,
	}, "\x00")))
	return "learn_" + hex.EncodeToString(sum[:])[:16]
}

func normalizeRefs(values []string) []string {
	seen := map[string]bool{}
	refs := make([]string, 0, len(values))
	for _, value := range values {
		ref := filepath.ToSlash(strings.TrimSpace(value))
		if ref == "" || seen[ref] {
			continue
		}
		seen[ref] = true
		refs = append(refs, ref)
	}
	sort.Strings(refs)
	return refs
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

func contains(values []string, wanted string) bool {
	for _, value := range values {
		if value == wanted {
			return true
		}
	}
	return false
}
