package repogovernance

type Manifest struct {
	SchemaVersion string  `json:"schema_version"`
	Groups        []Group `json:"groups"`
}

type Group struct {
	ID                 string   `json:"id"`
	DocumentSources    []string `json:"document_sources"`
	GeneratedDocuments []string `json:"generated_documents"`
	Code               []string `json:"code"`
	Tests              []string `json:"tests"`
	ChangePolicy       string   `json:"change_policy"`
}
