package controlplane

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
