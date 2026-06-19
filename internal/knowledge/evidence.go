package knowledge

type Evidence struct {
	Query        string   `json:"query"`
	ConceptCount int      `json:"concept_count"`
	EventCount   int      `json:"event_count"`
	HarnessCount int      `json:"harness_count"`
	HitCount     int      `json:"hit_count"`
	LinearIssues []string `json:"linear_issues,omitempty"`
	MustRead     []string `json:"must_read,omitempty"`
	CheckedAt    string   `json:"checked_at"`
}

func SummarizeSearch(report SearchReport) Evidence {
	return Evidence{
		Query:        report.Query,
		ConceptCount: len(report.Concepts),
		EventCount:   len(report.Events),
		HarnessCount: len(report.HarnessCases),
		HitCount:     len(report.Hits),
		LinearIssues: append([]string(nil), report.LinearIssues...),
		MustRead:     append([]string(nil), report.MustRead...),
		CheckedAt:    report.CheckedAt,
	}
}
