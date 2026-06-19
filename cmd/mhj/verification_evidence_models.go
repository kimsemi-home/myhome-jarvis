package main

type verificationEvidenceFile struct {
	SchemaVersion         string                       `json:"schema_version"`
	GraphArtifact         string                       `json:"graph_artifact"`
	Command               string                       `json:"command"`
	QualityJournal        string                       `json:"quality_journal"`
	QualityStatusCommand  string                       `json:"quality_status_command"`
	EvidenceStatusCommand string                       `json:"evidence_status_command"`
	Sources               []verificationEvidenceSource `json:"sources"`
}

type verificationEvidenceSource struct {
	ID       string `json:"id"`
	Kind     string `json:"kind"`
	Evidence string `json:"evidence"`
}

type evidencePolicyFile struct {
	PrivateSources []evidencePrivateSource `json:"private_sources"`
}

type evidencePrivateSource struct {
	Key      string `json:"key"`
	Path     string `json:"path"`
	NodeKind string `json:"node_kind"`
	Format   string `json:"format"`
}
