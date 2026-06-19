package main

type verificationGraphFile struct {
	SchemaVersion      string                `json:"schema_version"`
	GeneratedArtifacts []string              `json:"generated_artifacts"`
	Backends           []verificationBackend `json:"backends"`
	Units              []verificationUnit    `json:"units"`
}

type verificationBackend struct {
	ID   string `json:"id"`
	Path string `json:"path"`
}

type verificationUnit struct {
	ID       string   `json:"id"`
	Kind     string   `json:"kind"`
	Commands []string `json:"commands"`
}

type verificationConformanceFile struct {
	GraphArtifact                string                `json:"graph_artifact"`
	SchemaArtifact               string                `json:"schema_artifact"`
	TestsArtifact                string                `json:"tests_artifact"`
	ReleaseArtifact              string                `json:"release_artifact"`
	ControlPlaneVerifierArtifact string                `json:"control_plane_verifier_artifact"`
	BackendArtifacts             []verificationBackend `json:"backend_artifacts"`
}
