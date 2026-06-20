package contextpack

import (
	"bytes"
	"encoding/json"
	"testing"
)

func TestStatusDoesNotExposePrivateContextFields(t *testing.T) {
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
		[]byte("raw_prompt"), []byte("raw_transcript"),
		[]byte("credential"), []byte("cookie"), []byte("/" + "Users" + "/"),
		[]byte("private_evidence"), []byte("browser_session"),
		[]byte("unpublished_monetization"),
	} {
		if bytes.Contains(body, forbidden) {
			t.Fatalf("status leaked %s in %s", forbidden, body)
		}
	}
}
