package localfinanceevidence

import (
	"bytes"
	"os"
	"path/filepath"
	"testing"
)

func TestTamperedShortsChannelBindingFails(t *testing.T) {
	path := filepath.Join("..", "..", "fixtures", "local_finance", "manifest.json")
	manifest, err := Read(path)
	if err != nil {
		t.Fatal(err)
	}
	proofPath := filepath.Join(filepath.Dir(path), "proofs", "shorts-youtube-loopback.json")
	body, err := os.ReadFile(proofPath)
	if err != nil {
		t.Fatal(err)
	}
	ref := manifest.ExecutionProofs[len(manifest.ExecutionProofs)-1]
	body = bytes.Replace(body, []byte(`"binding_hash_matched": true`), []byte(`"binding_hash_matched": false`), 1)
	if err := validateProofBody(body, manifest.Month, ref); err == nil {
		t.Fatal("accepted a Shorts proof with a mismatched OAuth channel binding")
	}
}
