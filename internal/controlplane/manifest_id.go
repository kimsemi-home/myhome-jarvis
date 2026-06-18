package controlplane

import (
	"crypto/sha256"
	"encoding/hex"
	"strings"
)

func manifestID(manifest Manifest) string {
	sum := sha256.Sum256([]byte(strings.Join([]string{
		manifest.At,
		manifest.DecisionKind,
		manifest.PolicyVersion,
		manifest.OntologyVersion,
		manifest.SelectedRoute,
		manifest.OutputRef,
	}, "\x00")))
	return "cpm_" + hex.EncodeToString(sum[:])[:16]
}
