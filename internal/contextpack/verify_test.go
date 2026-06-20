package contextpack

import "testing"

func TestVerifyAcceptsMatchingDeclaration(t *testing.T) {
	root := t.TempDir()
	writePolicy(t, root, testPolicy())
	writeFile(t, root, ".mhj/context-pack.json", declarationJSON(validDeclaration()))

	result, err := VerifyDeclarationForRoot(root, "")
	if err != nil {
		t.Fatal(err)
	}
	if !result.Valid || result.DriftCount != 0 {
		t.Fatalf("result = %#v", result)
	}
}

func TestVerifyRejectsStaleDeclaration(t *testing.T) {
	root := t.TempDir()
	writePolicy(t, root, testPolicy())
	declaration := validDeclaration()
	declaration.OntologyVersion = "concept-registry/v0"
	declaration.SSOTArtifactVersions = declaration.SSOTArtifactVersions[:5]
	writeFile(t, root, ".mhj/context-pack.json", declarationJSON(declaration))

	result, err := VerifyDeclarationForRoot(root, "")
	if err != nil {
		t.Fatal(err)
	}
	if result.Valid || result.StaleVersionCount == 0 ||
		result.MissingArtifactCount == 0 {
		t.Fatalf("result = %#v", result)
	}
}
