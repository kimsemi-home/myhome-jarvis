package review

import "errors"

const PolicyRelativePath = "generated/review.generated.json"

var errMissingEvidenceRef = errors.New("review evidence ref is required")

type Policy struct {
	Context                 string         `json:"context"`
	Version                 string         `json:"version"`
	GeneratedArtifact       string         `json:"generated_artifact"`
	PrivateReviewQueue      string         `json:"private_review_queue"`
	AppendOnly              bool           `json:"append_only"`
	PublicStatusRedacted    bool           `json:"public_status_redacted"`
	RawReviewPublicAllowed  bool           `json:"raw_review_public_allowed"`
	MaxOpenReviews          int            `json:"max_open_reviews"`
	MaxHighRiskOpenReviews  int            `json:"max_high_risk_open_reviews"`
	MinBackupReviewers      int            `json:"min_backup_reviewers"`
	AllowedRisks            []string       `json:"allowed_risks"`
	QueueClasses            []string       `json:"queue_classes"`
	PriorityOrder           []string       `json:"priority_order"`
	AllowedStatuses         []string       `json:"allowed_statuses"`
	RequesterRoles          []string       `json:"requester_roles"`
	ReviewerRoles           []string       `json:"reviewer_roles"`
	OverloadPolicy          []OverloadRule `json:"overload_policy"`
	RequiredFields          []string       `json:"required_fields"`
	AllowedEvidencePrefixes []string       `json:"allowed_evidence_prefixes"`
	PublicSummaryFields     []string       `json:"public_summary_fields"`
	ForbiddenPublicFields   []string       `json:"forbidden_public_fields"`
	Commands                []string       `json:"commands"`
}

type OverloadRule struct {
	Key                   string `json:"key"`
	AllowedWhenOverloaded bool   `json:"allowed_when_overloaded"`
}
