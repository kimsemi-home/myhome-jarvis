package incidents

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

const PolicyRelativePath = "generated/incidents.generated.json"

type Policy struct {
	Context                   string   `json:"context"`
	Version                   string   `json:"version"`
	GeneratedArtifact         string   `json:"generated_artifact"`
	PrivateIncidentLedger     string   `json:"private_incident_ledger"`
	AppendOnly                bool     `json:"append_only"`
	PublicStatusRedacted      bool     `json:"public_status_redacted"`
	RawIncidentPublicAllowed  bool     `json:"raw_incident_public_allowed"`
	QuarantineStaleAfterHours int      `json:"quarantine_stale_after_hours"`
	AllowedKinds              []string `json:"allowed_kinds"`
	Lifecycle                 []string `json:"lifecycle"`
	AllowedStatuses           []string `json:"allowed_statuses"`
	OwnerRoles                []string `json:"owner_roles"`
	QuarantineStates          []string `json:"quarantine_states"`
	RequiredFields            []string `json:"required_fields"`
	AllowedEvidencePrefixes   []string `json:"allowed_evidence_prefixes"`
	PublicSummaryFields       []string `json:"public_summary_fields"`
	ForbiddenPublicFields     []string `json:"forbidden_public_fields"`
	Commands                  []string `json:"commands"`
}

type Incident struct {
	ID              string   `json:"id"`
	At              string   `json:"at"`
	Kind            string   `json:"kind"`
	Stage           string   `json:"stage"`
	Status          string   `json:"status"`
	OwnerRole       string   `json:"owner_role"`
	QuarantineState string   `json:"quarantine_state"`
	EvidenceRefs    []string `json:"evidence_refs"`
}

type Status struct {
	PolicyPath                string         `json:"policy_path"`
	LedgerPath                string         `json:"ledger_path"`
	Exists                    bool           `json:"exists"`
	Count                     int            `json:"count"`
	OpenCount                 int            `json:"open_count"`
	ClosedCount               int            `json:"closed_count"`
	InvalidIncidentCount      int            `json:"invalid_incident_count"`
	IncidentDebtCount         int            `json:"incident_debt_count"`
	MissingOwnerCount         int            `json:"missing_owner_count"`
	MissingEvidenceRefCount   int            `json:"missing_evidence_ref_count"`
	StaleQuarantineCount      int            `json:"stale_quarantine_count"`
	QuarantineStaleAfterHours int            `json:"quarantine_stale_after_hours"`
	ByKind                    map[string]int `json:"by_kind"`
	ByStage                   map[string]int `json:"by_stage"`
	ByStatus                  map[string]int `json:"by_status"`
	ByOwnerRole               map[string]int `json:"by_owner_role"`
	ByQuarantineState         map[string]int `json:"by_quarantine_state"`
	LastObservedAt            string         `json:"last_observed_at,omitempty"`
	CheckedAt                 string         `json:"checked_at"`
}

var (
	errMissingOwner    = errors.New("incident owner role is required")
	errMissingEvidence = errors.New("incident evidence refs are required")
)

func StatusForRoot(root string) (Status, error) {
	policy, err := ReadPolicy(root)
	if err != nil {
		return Status{}, err
	}
	checkedAt := time.Now().UTC()
	status := Status{
		PolicyPath:                PolicyRelativePath,
		LedgerPath:                policy.PrivateIncidentLedger,
		QuarantineStaleAfterHours: policy.QuarantineStaleAfterHours,
		ByKind:                    map[string]int{},
		ByStage:                   map[string]int{},
		ByStatus:                  map[string]int{},
		ByOwnerRole:               map[string]int{},
		ByQuarantineState:         map[string]int{},
		CheckedAt:                 checkedAt.Format(time.RFC3339),
	}
	file, err := os.Open(filepath.Join(root, filepath.FromSlash(policy.PrivateIncidentLedger)))
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
		var incident Incident
		if err := json.Unmarshal([]byte(line), &incident); err != nil {
			status.InvalidIncidentCount++
			continue
		}
		normalized, err := normalizeIncident(policy, incident)
		if err != nil {
			switch {
			case errors.Is(err, errMissingOwner):
				status.MissingOwnerCount++
			case errors.Is(err, errMissingEvidence):
				status.MissingEvidenceRefCount++
			default:
				status.InvalidIncidentCount++
			}
			continue
		}
		status.Count++
		status.ByKind[normalized.Kind]++
		status.ByStage[normalized.Stage]++
		status.ByStatus[normalized.Status]++
		status.ByOwnerRole[normalized.OwnerRole]++
		status.ByQuarantineState[normalized.QuarantineState]++
		if normalized.Status == "closed" {
			status.ClosedCount++
		} else {
			status.OpenCount++
		}
		if isStaleQuarantine(policy, normalized, checkedAt) {
			status.StaleQuarantineCount++
		}
		status.LastObservedAt = laterRFC3339(status.LastObservedAt, normalized.At)
	}
	if err := scanner.Err(); err != nil {
		return Status{}, err
	}
	status.IncidentDebtCount = status.InvalidIncidentCount + status.MissingOwnerCount + status.MissingEvidenceRefCount + status.StaleQuarantineCount
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
		return fmt.Errorf("incident policy context = %q", policy.Context)
	}
	if !strings.HasPrefix(policy.PrivateIncidentLedger, "data/private/") || !strings.HasSuffix(policy.PrivateIncidentLedger, ".jsonl") {
		return fmt.Errorf("incident ledger must stay in data/private JSONL")
	}
	if !policy.AppendOnly || !policy.PublicStatusRedacted || policy.RawIncidentPublicAllowed {
		return fmt.Errorf("incident policy must be private append-only with redacted public status")
	}
	if policy.QuarantineStaleAfterHours <= 0 {
		return fmt.Errorf("incident quarantine stale threshold must be positive")
	}
	for _, kind := range []string{"quality_regression", "public_safety", "evidence_gap", "authority_violation", "quarantine", "feedback_loop_gap"} {
		if !contains(normalizeList(policy.AllowedKinds), kind) {
			return fmt.Errorf("incident kind %q is missing", kind)
		}
	}
	for _, stage := range []string{"observed", "evidence_recorded", "classified", "owner_assigned", "fix_planned", "verified", "knowledge_updated"} {
		if !contains(normalizeList(policy.Lifecycle), stage) {
			return fmt.Errorf("incident lifecycle stage %q is missing", stage)
		}
	}
	for _, status := range []string{"open", "mitigating", "verified", "closed", "quarantined"} {
		if !contains(normalizeList(policy.AllowedStatuses), status) {
			return fmt.Errorf("incident status %q is missing", status)
		}
	}
	for _, role := range []string{"producer", "independent_reviewer", "adversarial_reviewer", "deterministic_verifier", "governance_steward"} {
		if !contains(normalizeList(policy.OwnerRoles), role) {
			return fmt.Errorf("incident owner role %q is missing", role)
		}
	}
	for _, state := range []string{"none", "quarantined", "release_requested", "released"} {
		if !contains(normalizeList(policy.QuarantineStates), state) {
			return fmt.Errorf("incident quarantine state %q is missing", state)
		}
	}
	required := normalizeList(policy.RequiredFields)
	for _, field := range []string{"at", "kind", "stage", "status", "owner_role", "evidence_refs"} {
		if !contains(required, field) {
			return fmt.Errorf("incident required field %q is missing", field)
		}
	}
	summary := normalizeList(policy.PublicSummaryFields)
	for _, field := range []string{"count", "open_count", "incident_debt_count", "missing_owner_count", "missing_evidence_ref_count", "stale_quarantine_count", "by_stage", "by_owner_role", "checked_at"} {
		if !contains(summary, field) {
			return fmt.Errorf("incident public summary missing %q", field)
		}
	}
	if !contains(policy.Commands, "mhj incidents status") {
		return fmt.Errorf("incident status command is missing")
	}
	return nil
}

func normalizeIncident(policy Policy, incident Incident) (Incident, error) {
	normalized := Incident{
		ID:              publicText(incident.ID),
		At:              publicText(incident.At),
		Kind:            normalizeToken(incident.Kind),
		Stage:           normalizeToken(incident.Stage),
		Status:          normalizeToken(incident.Status),
		OwnerRole:       normalizeToken(incident.OwnerRole),
		QuarantineState: normalizeToken(incident.QuarantineState),
		EvidenceRefs:    normalizeRefs(incident.EvidenceRefs),
	}
	if normalized.QuarantineState == "" {
		normalized.QuarantineState = "none"
	}
	if normalized.At == "" {
		return Incident{}, fmt.Errorf("incident at timestamp is required")
	}
	if _, err := time.Parse(time.RFC3339, normalized.At); err != nil {
		return Incident{}, fmt.Errorf("incident at timestamp is invalid")
	}
	if !contains(normalizeList(policy.AllowedKinds), normalized.Kind) {
		return Incident{}, fmt.Errorf("incident kind %q is not allowed", normalized.Kind)
	}
	if !contains(normalizeList(policy.Lifecycle), normalized.Stage) {
		return Incident{}, fmt.Errorf("incident stage %q is not allowed", normalized.Stage)
	}
	if !contains(normalizeList(policy.AllowedStatuses), normalized.Status) {
		return Incident{}, fmt.Errorf("incident status %q is not allowed", normalized.Status)
	}
	if normalized.OwnerRole == "" {
		return Incident{}, errMissingOwner
	}
	if !contains(normalizeList(policy.OwnerRoles), normalized.OwnerRole) {
		return Incident{}, fmt.Errorf("incident owner role %q is not allowed", normalized.OwnerRole)
	}
	if !contains(normalizeList(policy.QuarantineStates), normalized.QuarantineState) {
		return Incident{}, fmt.Errorf("incident quarantine state %q is not allowed", normalized.QuarantineState)
	}
	if len(normalized.EvidenceRefs) == 0 {
		return Incident{}, errMissingEvidence
	}
	for _, ref := range normalized.EvidenceRefs {
		if err := validateRef(policy, ref); err != nil {
			return Incident{}, err
		}
		if err := rejectSensitiveText(ref); err != nil {
			return Incident{}, err
		}
	}
	return normalized, nil
}

func validateRef(policy Policy, ref string) error {
	ref = filepath.ToSlash(strings.TrimSpace(ref))
	if ref == "" {
		return errMissingEvidence
	}
	if filepath.IsAbs(filepath.FromSlash(ref)) || strings.Contains(ref, "..") {
		return fmt.Errorf("incident evidence ref must be repo-relative")
	}
	for _, prefix := range policy.AllowedEvidencePrefixes {
		if strings.HasPrefix(ref, prefix) {
			return nil
		}
	}
	return fmt.Errorf("incident evidence ref %q is outside allowed prefixes", ref)
}

func isStaleQuarantine(policy Policy, incident Incident, checkedAt time.Time) bool {
	if incident.QuarantineState != "quarantined" && incident.Status != "quarantined" {
		return false
	}
	at, err := time.Parse(time.RFC3339, incident.At)
	if err != nil {
		return false
	}
	return checkedAt.Sub(at) > time.Duration(policy.QuarantineStaleAfterHours)*time.Hour
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
			return fmt.Errorf("incident evidence ref contains forbidden private marker")
		}
	}
	return nil
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
