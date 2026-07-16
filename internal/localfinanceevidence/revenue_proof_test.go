package localfinanceevidence

import (
	"bytes"
	"os"
	"path/filepath"
	"testing"
)

func TestTamperedRevenueProofFails(t *testing.T) {
	source := filepath.Join("..", "..", "fixtures", "local_finance", "proofs", "revenue-youtube-loopback.json")
	body, err := os.ReadFile(source)
	if err != nil {
		t.Fatal(err)
	}
	root := t.TempDir()
	for _, mutation := range []struct {
		name, from, to string
	}{
		{"reconciliation", `"net_minor": 6300`, `"net_minor": 6301`},
		{"OAuth origin pin", `"official_origin_pinned": true`, `"official_origin_pinned": false`},
	} {
		tampered := bytes.Replace(body, []byte(mutation.from), []byte(mutation.to), 1)
		if bytes.Equal(body, tampered) {
			t.Fatalf("Revenue proof fixture did not contain %s value", mutation.name)
		}
		if err := os.WriteFile(filepath.Join(root, "proof.json"), tampered, 0o600); err != nil {
			t.Fatal(err)
		}
		ref := ProofRef{Component: "revenue", Capability: "youtube-revenue-collection-rehearsal",
			ProofSchema: revenueProofSchema, Path: "proof.json", ArtifactSHA256: digest(string(tampered)),
			ReportHash: "c5ca6ec3de490614812d9f18976275779d0461663df889d07fffd17fe778d88a"}
		if err := validateProofFiles(root, "2026-07", []ProofRef{ref}); err == nil {
			t.Fatalf("expected tampered Revenue %s proof to fail", mutation.name)
		}
	}
}
