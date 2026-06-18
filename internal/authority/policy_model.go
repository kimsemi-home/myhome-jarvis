package authority

const PolicyRelativePath = "generated/authority.generated.json"

type Policy struct {
	Context                     string           `json:"context"`
	Version                     string           `json:"version"`
	GeneratedArtifact           string           `json:"generated_artifact"`
	PublicStatusRedacted        bool             `json:"public_status_redacted"`
	SelfAuthorityAllowed        bool             `json:"self_authority_allowed"`
	ReasoningTierGrantsApproval bool             `json:"reasoning_tier_grants_approval"`
	PublicRepoHighRiskBlocked   bool             `json:"public_repo_high_risk_blocked"`
	RequiredInputs              []string         `json:"required_inputs"`
	ReasoningTiers              []ReasoningTier  `json:"reasoning_tiers"`
	RolePermissions             []RolePermission `json:"role_permissions"`
	DomainAttributes            []string         `json:"domain_attributes"`
	Decisions                   []Decision       `json:"decisions"`
	Outcomes                    []string         `json:"outcomes"`
	AuthorityDebtClasses        []string         `json:"authority_debt_classes"`
	PublicSummaryFields         []string         `json:"public_summary_fields"`
	ForbiddenPublicFields       []string         `json:"forbidden_public_fields"`
	Commands                    []string         `json:"commands"`
}

type ReasoningTier struct {
	Key     string   `json:"key"`
	Label   string   `json:"label"`
	May     []string `json:"may"`
	MustNot []string `json:"must_not"`
}

type RolePermission struct {
	Role    string   `json:"role"`
	May     []string `json:"may"`
	MustNot []string `json:"must_not"`
}

type Decision struct {
	Key                 string `json:"key"`
	Risk                string `json:"risk"`
	PublicRepoAllowed   bool   `json:"public_repo_allowed"`
	RequiresHumanReview bool   `json:"requires_human_review"`
	AllowedWhenBlocked  bool   `json:"allowed_when_blocked"`
}
