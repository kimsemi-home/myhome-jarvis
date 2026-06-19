package evidencequality

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
