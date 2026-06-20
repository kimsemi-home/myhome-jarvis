package mergeevidence

import (
	"bytes"
	"encoding/json"
	"testing"
)

func TestStatusDoesNotExposePrivateEvidenceFields(t *testing.T) {
	root := t.TempDir()
	writePolicy(t, root, testPolicy())

	status, err := StatusForRoot(root)
	if err != nil {
		t.Fatal(err)
	}
	body, err := json.Marshal(status)
	if err != nil {
		t.Fatal(err)
	}
	for _, forbidden := range [][]byte{
		[]byte("private_linear_url"), []byte("local_absolute_path"),
		[]byte("raw_review_notes"), []byte("private_evidence"),
		[]byte("credential"), []byte("access_token"),
	} {
		if bytes.Contains(body, forbidden) {
			t.Fatalf("merge evidence status leaked %s in %s", forbidden, body)
		}
	}
}
