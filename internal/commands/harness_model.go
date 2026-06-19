package commands

type HarnessCase struct {
	Name       string `json:"name"`
	Command    string `json:"command"`
	Payload    string `json:"payload"`
	ShouldPass bool   `json:"should_pass"`
	Contains   string `json:"contains,omitempty"`
}

type HarnessCaseResult struct {
	Name    string `json:"name"`
	Passed  bool   `json:"passed"`
	Message string `json:"message"`
}

type HarnessReport struct {
	Name    string              `json:"name"`
	Passed  bool                `json:"passed"`
	Results []HarnessCaseResult `json:"results"`
}
