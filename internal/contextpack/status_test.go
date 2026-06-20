package contextpack

import "testing"

func TestStatusSummarizesPublicSafeContextPack(t *testing.T) {
	root := t.TempDir()
	writePolicy(t, root, testPolicy())

	status, err := StatusForRoot(root)
	if err != nil {
		t.Fatal(err)
	}
	if !status.PublicSafe || status.SplitCriteriaCount != 5 ||
		status.ExportedArtifactCount != 6 {
		t.Fatalf("status = %#v", status)
	}
	if status.UpstreamCompatibilityVersion != "myhome-jarvis/context-pack/v1" ||
		status.OntologyVersion != "concept-registry/v1" {
		t.Fatalf("versions = %#v", status)
	}
}
