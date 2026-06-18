package review

type Review struct {
	ID              string   `json:"id"`
	At              string   `json:"at"`
	ItemKey         string   `json:"item_key"`
	QueueClass      string   `json:"queue_class"`
	Risk            string   `json:"risk"`
	Status          string   `json:"status"`
	RequesterRole   string   `json:"requester_role"`
	ReviewerRole    string   `json:"reviewer_role"`
	BackupAvailable bool     `json:"backup_available"`
	EvidenceRefs    []string `json:"evidence_refs"`
}

type Status struct {
	PolicyPath             string         `json:"policy_path"`
	QueuePath              string         `json:"queue_path"`
	Exists                 bool           `json:"exists"`
	Count                  int            `json:"count"`
	OpenCount              int            `json:"open_count"`
	HighRiskOpenCount      int            `json:"high_risk_open_count"`
	InvalidReviewCount     int            `json:"invalid_review_count"`
	MissingEvidenceCount   int            `json:"missing_evidence_count"`
	MissingReviewerCount   int            `json:"missing_reviewer_count"`
	BackupAvailableCount   int            `json:"backup_available_count"`
	ReviewDebtCount        int            `json:"review_debt_count"`
	CapacityState          string         `json:"capacity_state"`
	ActiveRule             string         `json:"active_rule"`
	MaxOpenReviews         int            `json:"max_open_reviews"`
	MaxHighRiskOpenReviews int            `json:"max_high_risk_open_reviews"`
	ByRisk                 map[string]int `json:"by_risk"`
	ByStatus               map[string]int `json:"by_status"`
	ByReviewerRole         map[string]int `json:"by_reviewer_role"`
	ByQueueClass           map[string]int `json:"by_queue_class"`
	LastObservedAt         string         `json:"last_observed_at,omitempty"`
	CheckedAt              string         `json:"checked_at"`
}
