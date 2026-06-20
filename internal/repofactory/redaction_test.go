package repofactory

import (
	"bytes"
	"encoding/json"
	"testing"
)

func TestStatusDoesNotExposePrivateTemplatePayload(t *testing.T) {
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
		[]byte("absolute_home_path"),
		[]byte("old_private_owner"),
		[]byte("private_team_slug"),
		[]byte("client_secret"),
		[]byte("access_token"),
		[]byte("/" + "Users" + "/"),
	} {
		if bytes.Contains(body, forbidden) {
			t.Fatalf("repo factory status leaked %s in %s", forbidden, body)
		}
	}
}
