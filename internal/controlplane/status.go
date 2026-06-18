package controlplane

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

const PolicyRelativePath = "generated/control_plane.generated.json"

type Policy struct {
	Context                    string   `json:"context"`
	Version                    string   `json:"version"`
	GeneratedArtifact          string   `json:"generated_artifact"`
	PrivateManifestLedger      string   `json:"private_manifest_ledger"`
	ManifestRequired           bool     `json:"manifest_required"`
	AppendOnly                 bool     `json:"append_only"`
	PublicStatusRedacted       bool     `json:"public_status_redacted"`
	RawRationalePublicAllowed  bool     `json:"raw_rationale_public_allowed"`
	VerifierSeparationRequired bool     `json:"verifier_separation_required"`
	MinLeaseSeconds            int      `json:"min_lease_seconds"`
	MaxLeaseSeconds            int      `json:"max_lease_seconds"`
	AllowedDecisionKinds       []string `json:"allowed_decision_kinds"`
	AllowedAuthorityProfiles   []string `json:"allowed_authority_profiles"`
	AllowedLeaseStatuses       []string `json:"allowed_lease_statuses"`
	RequiredFields             []string `json:"required_fields"`
	AllowedEvidencePrefixes    []string `json:"allowed_evidence_prefixes"`
	PublicSummaryFields        []string `json:"public_summary_fields"`
	ForbiddenPublicFields      []string `json:"forbidden_public_fields"`
	Commands                   []string `json:"commands"`
}

type ManifestRequest struct {
	DecisionKind     string   `json:"decision_kind"`
	PolicyVersion    string   `json:"policy_version"`
	OntologyVersion  string   `json:"ontology_version"`
	AuthorityProfile string   `json:"authority_profile"`
	SelectedRoute    string   `json:"selected_route"`
	ReviewerRole     string   `json:"reviewer_role"`
	VerifierRole     string   `json:"verifier_role"`
	LeaseSeconds     int      `json:"lease_seconds"`
	LeaseStatus      string   `json:"lease_status"`
	EvidenceRefs     []string `json:"evidence_refs"`
	OutputRef        string   `json:"output_ref"`
}

type Manifest struct {
	ID               string   `json:"id"`
	At               string   `json:"at"`
	DecisionKind     string   `json:"decision_kind"`
	PolicyVersion    string   `json:"policy_version"`
	OntologyVersion  string   `json:"ontology_version"`
	AuthorityProfile string   `json:"authority_profile"`
	SelectedRoute    string   `json:"selected_route"`
	ReviewerRole     string   `json:"reviewer_role"`
	VerifierRole     string   `json:"verifier_role"`
	LeaseSeconds     int      `json:"lease_seconds"`
	LeaseStatus      string   `json:"lease_status"`
	EvidenceRefs     []string `json:"evidence_refs"`
	OutputRef        string   `json:"output_ref"`
}

type RecordResult struct {
	ID           string `json:"id"`
	ManifestPath string `json:"manifest_path"`
	DecisionKind string `json:"decision_kind"`
	LeaseStatus  string `json:"lease_status"`
	RecordedAt   string `json:"recorded_at"`
}

type Status struct {
	PolicyPath                 string         `json:"policy_path"`
	ManifestPath               string         `json:"manifest_path"`
	Exists                     bool           `json:"exists"`
	Count                      int            `json:"count"`
	InvalidManifestCount       int            `json:"invalid_manifest_count"`
	ManifestDebtCount          int            `json:"manifest_debt_count"`
	VerifierSeparationRequired bool           `json:"verifier_separation_required"`
	VerifierViolationCount     int            `json:"verifier_violation_count"`
	MinLeaseSeconds            int            `json:"min_lease_seconds"`
	MaxLeaseSeconds            int            `json:"max_lease_seconds"`
	ByDecisionKind             map[string]int `json:"by_decision_kind"`
	ByAuthorityProfile         map[string]int `json:"by_authority_profile"`
	ByLeaseStatus              map[string]int `json:"by_lease_status"`
	LastObservedAt             string         `json:"last_observed_at,omitempty"`
	CheckedAt                  string         `json:"checked_at"`
}

func AppendManifest(root string, request ManifestRequest) (RecordResult, error) {
	policy, err := ReadPolicy(root)
	if err != nil {
		return RecordResult{}, err
	}
	manifest, err := normalizeManifest(policy, request)
	if err != nil {
		return RecordResult{}, err
	}
	path := filepath.Join(root, filepath.FromSlash(policy.PrivateManifestLedger))
	if err := os.MkdirAll(filepath.Dir(path), 0o700); err != nil {
		return RecordResult{}, err
	}
	file, err := os.OpenFile(path, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0o600)
	if err != nil {
		return RecordResult{}, err
	}
	defer file.Close()
	data, err := json.Marshal(manifest)
	if err != nil {
		return RecordResult{}, err
	}
	if _, err := file.Write(append(data, '\n')); err != nil {
		return RecordResult{}, err
	}
	return RecordResult{
		ID:           manifest.ID,
		ManifestPath: policy.PrivateManifestLedger,
		DecisionKind: manifest.DecisionKind,
		LeaseStatus:  manifest.LeaseStatus,
		RecordedAt:   manifest.At,
	}, nil
}

func StatusForRoot(root string) (Status, error) {
	policy, err := ReadPolicy(root)
	if err != nil {
		return Status{}, err
	}
	status := Status{
		PolicyPath:                 PolicyRelativePath,
		ManifestPath:               policy.PrivateManifestLedger,
		VerifierSeparationRequired: policy.VerifierSeparationRequired,
		MinLeaseSeconds:            policy.MinLeaseSeconds,
		MaxLeaseSeconds:            policy.MaxLeaseSeconds,
		ByDecisionKind:             map[string]int{},
		ByAuthorityProfile:         map[string]int{},
		ByLeaseStatus:              map[string]int{},
		CheckedAt:                  time.Now().UTC().Format(time.RFC3339),
	}
	file, err := os.Open(filepath.Join(root, filepath.FromSlash(policy.PrivateManifestLedger)))
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
		if containsForbiddenManifestMarker(policy, line) {
			status.InvalidManifestCount++
			continue
		}
		var manifest Manifest
		if err := json.Unmarshal([]byte(line), &manifest); err != nil {
			status.InvalidManifestCount++
			continue
		}
		if err := validateManifest(policy, manifest); err != nil {
			status.InvalidManifestCount++
			continue
		}
		status.Count++
		status.ByDecisionKind[manifest.DecisionKind]++
		status.ByAuthorityProfile[manifest.AuthorityProfile]++
		status.ByLeaseStatus[manifest.LeaseStatus]++
		if policy.VerifierSeparationRequired && manifest.ReviewerRole == manifest.VerifierRole {
			status.VerifierViolationCount++
		}
		status.LastObservedAt = laterRFC3339(status.LastObservedAt, manifest.At)
	}
	if err := scanner.Err(); err != nil {
		return Status{}, err
	}
	status.ManifestDebtCount = status.InvalidManifestCount + status.VerifierViolationCount
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
	if policy.Context != "AgentOps" {
		return fmt.Errorf("control-plane policy context = %q", policy.Context)
	}
	if !strings.HasPrefix(policy.PrivateManifestLedger, "data/private/") || !strings.HasSuffix(policy.PrivateManifestLedger, ".jsonl") {
		return fmt.Errorf("control-plane manifest ledger must stay in data/private JSONL")
	}
	if !policy.ManifestRequired || !policy.AppendOnly || !policy.PublicStatusRedacted || policy.RawRationalePublicAllowed {
		return fmt.Errorf("control-plane policy must require private append-only redacted manifests")
	}
	if !policy.VerifierSeparationRequired {
		return fmt.Errorf("control-plane verifier separation must be required")
	}
	if policy.MinLeaseSeconds <= 0 || policy.MaxLeaseSeconds <= policy.MinLeaseSeconds {
		return fmt.Errorf("control-plane lease bounds are invalid")
	}
	for _, kind := range []string{"loop_once", "loop_worker_cycle", "checkpoint_write"} {
		if !contains(normalizeList(policy.AllowedDecisionKinds), kind) {
			return fmt.Errorf("control-plane decision kind %q is missing", kind)
		}
	}
	for _, profile := range []string{"local_readonly", "external_write_gated"} {
		if !contains(normalizeList(policy.AllowedAuthorityProfiles), profile) {
			return fmt.Errorf("control-plane authority profile %q is missing", profile)
		}
	}
	for _, status := range []string{"issued", "active", "finished", "aborted", "quarantined"} {
		if !contains(normalizeList(policy.AllowedLeaseStatuses), status) {
			return fmt.Errorf("control-plane lease status %q is missing", status)
		}
	}
	required := normalizeList(policy.RequiredFields)
	for _, field := range []string{"decision_kind", "policy_version", "ontology_version", "authority_profile", "selected_route", "reviewer_role", "verifier_role", "lease_seconds", "lease_status", "evidence_refs", "output_ref"} {
		if !contains(required, field) {
			return fmt.Errorf("control-plane required field %q is missing", field)
		}
	}
	summary := normalizeList(policy.PublicSummaryFields)
	for _, field := range []string{"count", "invalid_manifest_count", "manifest_debt_count", "verifier_violation_count", "by_decision_kind", "by_authority_profile", "by_lease_status", "checked_at"} {
		if !contains(summary, field) {
			return fmt.Errorf("control-plane public summary missing %q", field)
		}
	}
	if !contains(policy.Commands, "mhj control-plane status") {
		return fmt.Errorf("control-plane status command is missing")
	}
	return nil
}

func normalizeManifest(policy Policy, request ManifestRequest) (Manifest, error) {
	manifest := Manifest{
		At:               time.Now().UTC().Format(time.RFC3339),
		DecisionKind:     normalizeToken(request.DecisionKind),
		PolicyVersion:    publicText(request.PolicyVersion),
		OntologyVersion:  publicText(request.OntologyVersion),
		AuthorityProfile: normalizeToken(request.AuthorityProfile),
		SelectedRoute:    normalizeToken(request.SelectedRoute),
		ReviewerRole:     normalizeToken(request.ReviewerRole),
		VerifierRole:     normalizeToken(request.VerifierRole),
		LeaseSeconds:     request.LeaseSeconds,
		LeaseStatus:      normalizeToken(request.LeaseStatus),
		EvidenceRefs:     normalizeRefs(request.EvidenceRefs),
		OutputRef:        filepath.ToSlash(strings.TrimSpace(request.OutputRef)),
	}
	if err := validateManifest(policy, manifest); err != nil {
		return Manifest{}, err
	}
	manifest.ID = manifestID(manifest)
	return manifest, nil
}

func validateManifest(policy Policy, manifest Manifest) error {
	if !contains(normalizeList(policy.AllowedDecisionKinds), manifest.DecisionKind) {
		return fmt.Errorf("control-plane decision kind %q is not allowed", manifest.DecisionKind)
	}
	if manifest.PolicyVersion == "" || manifest.OntologyVersion == "" || manifest.SelectedRoute == "" {
		return fmt.Errorf("control-plane manifest requires policy_version, ontology_version, and selected_route")
	}
	if !contains(normalizeList(policy.AllowedAuthorityProfiles), manifest.AuthorityProfile) {
		return fmt.Errorf("control-plane authority profile %q is not allowed", manifest.AuthorityProfile)
	}
	if manifest.ReviewerRole == "" || manifest.VerifierRole == "" {
		return fmt.Errorf("control-plane manifest requires reviewer_role and verifier_role")
	}
	if policy.VerifierSeparationRequired && manifest.ReviewerRole == manifest.VerifierRole {
		return fmt.Errorf("control-plane reviewer and verifier roles must be separated")
	}
	if manifest.LeaseSeconds < policy.MinLeaseSeconds || manifest.LeaseSeconds > policy.MaxLeaseSeconds {
		return fmt.Errorf("control-plane lease seconds out of bounds")
	}
	if !contains(normalizeList(policy.AllowedLeaseStatuses), manifest.LeaseStatus) {
		return fmt.Errorf("control-plane lease status %q is not allowed", manifest.LeaseStatus)
	}
	if len(manifest.EvidenceRefs) == 0 || manifest.OutputRef == "" {
		return fmt.Errorf("control-plane manifest requires evidence refs and output ref")
	}
	for _, ref := range append(append([]string{}, manifest.EvidenceRefs...), manifest.OutputRef) {
		if err := validateRef(policy, ref); err != nil {
			return err
		}
	}
	values := append([]string{
		manifest.PolicyVersion,
		manifest.OntologyVersion,
		manifest.SelectedRoute,
		manifest.ReviewerRole,
		manifest.VerifierRole,
		manifest.OutputRef,
	}, manifest.EvidenceRefs...)
	for _, value := range values {
		if err := rejectSensitiveText(value); err != nil {
			return err
		}
		if err := rejectForbiddenPolicyText(policy, value); err != nil {
			return err
		}
	}
	return nil
}

func validateRef(policy Policy, ref string) error {
	ref = filepath.ToSlash(strings.TrimSpace(ref))
	if ref == "" {
		return fmt.Errorf("control-plane ref is required")
	}
	if filepath.IsAbs(filepath.FromSlash(ref)) || strings.Contains(ref, "..") {
		return fmt.Errorf("control-plane ref must be repo-relative")
	}
	for _, prefix := range policy.AllowedEvidencePrefixes {
		if strings.HasPrefix(ref, prefix) {
			return nil
		}
	}
	return fmt.Errorf("control-plane ref %q is outside allowed prefixes", ref)
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
		"raw_rationale",
		"selection_rationale",
		"candidate_agents",
		"raw_prompt",
		"raw_transcript",
		"private_evidence",
		"account_id",
		"card_number",
		"api_secret",
		"credential=",
		"linear.app/",
		string(filepath.Separator) + "users" + string(filepath.Separator),
		"\\" + "users" + "\\",
	} {
		if strings.Contains(lower, marker) {
			return fmt.Errorf("control-plane manifest contains forbidden private marker")
		}
	}
	return nil
}

func rejectForbiddenPolicyText(policy Policy, value string) error {
	lower := strings.ToLower(value)
	for _, field := range policy.ForbiddenPublicFields {
		marker := strings.ToLower(strings.TrimSpace(field))
		if marker != "" && strings.Contains(lower, marker) {
			return fmt.Errorf("control-plane manifest contains forbidden public marker")
		}
	}
	return nil
}

func containsForbiddenManifestMarker(policy Policy, line string) bool {
	if rejectSensitiveText(line) != nil {
		return true
	}
	return rejectForbiddenPolicyText(policy, line) != nil
}

func manifestID(manifest Manifest) string {
	sum := sha256.Sum256([]byte(strings.Join([]string{
		manifest.At,
		manifest.DecisionKind,
		manifest.PolicyVersion,
		manifest.OntologyVersion,
		manifest.SelectedRoute,
		manifest.OutputRef,
	}, "\x00")))
	return "cpm_" + hex.EncodeToString(sum[:])[:16]
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
