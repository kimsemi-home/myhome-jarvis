package repogovernance

type Document struct {
	SchemaVersion string    `json:"schema_version"`
	ID            string    `json:"id"`
	Title         string    `json:"title"`
	Output        string    `json:"output"`
	Summary       string    `json:"summary"`
	Sections      []Section `json:"sections"`
}

type Section struct {
	Heading      string   `json:"heading"`
	Paragraphs   []string `json:"paragraphs,omitempty"`
	Bullets      []string `json:"bullets,omitempty"`
	CodeLanguage string   `json:"code_language,omitempty"`
	Code         string   `json:"code,omitempty"`
}
