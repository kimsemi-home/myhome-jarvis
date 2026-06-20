package mediareadiness

import (
	"bytes"
	"encoding/json"
	"testing"
)

func TestStatusDoesNotExposePayloadsOrAccountData(t *testing.T) {
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
		[]byte(`"payload"`), []byte(`"query"`), []byte("sample media"),
		[]byte("cookie"), []byte("credential"), []byte("account_id"),
	} {
		if bytes.Contains(body, forbidden) {
			t.Fatalf("media readiness status leaked %s in %s", forbidden, body)
		}
	}
}
