package contextpack

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"
)

func testPolicy() Policy {
	return Policy{Context: "CrossRepoContextPack", Version: "v1",
		GeneratedArtifact: PolicyRelativePath, PackID: "myhome-jarvis/context-pack",
		UpstreamCompatibilityVersion: "myhome-jarvis/context-pack/v1",
		OntologyVersion:              "concept-registry/v1", DeclarationPath: ".mhj/context-pack.json",
		PublicStatusRedacted: true, MissionSource: "generated/assistant_vision.generated.json",
		BoundedContextSource: "generated/concepts.generated.json",
		SplitCriteria:        testSplitCriteria(), ExportedArtifacts: testArtifacts(),
		AuthorityContract: AuthorityContract{Path: "generated/authority.generated.json",
			Version: "authority/v1", PublicSafetyGateRequired: true},
		SecurityContract: SecurityContract{Path: "generated/security.generated.json",
			Version: "security/v1"}, VerificationProfile: VerificationProfile{
			Name: "quality", Graph: "generated/verification_graph.generated.json",
			RequiredUnits: []string{"public-safety", "ssot", "go", "rust", "flutter"}},
		RequiredDeclarationFields: requiredDeclarationFields,
		Commands:                  requiredCommands}
}

func writePolicy(t *testing.T, root string, policy Policy) {
	t.Helper()
	body, err := json.Marshal(policy)
	if err != nil {
		t.Fatal(err)
	}
	writeFile(t, root, PolicyRelativePath, string(body)+"\n")
}

func writeFile(t *testing.T, root string, rel string, body string) {
	t.Helper()
	path := filepath.Join(root, filepath.FromSlash(rel))
	if err := os.MkdirAll(filepath.Dir(path), 0o700); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(path, []byte(body), 0o600); err != nil {
		t.Fatal(err)
	}
}
