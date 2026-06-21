package externalbootstrap

import (
	"crypto/sha256"
	"encoding/hex"
	"os"
	"path/filepath"

	"github.com/kimsemi-home/myhome-jarvis/internal/repofactory"
)

func hashCacheInputs(
	root string,
	factory repofactory.DecisionPacket,
) ([]HashCacheInput, error) {
	inputs := []HashCacheInput{}
	for _, spec := range hashCacheSpecs(factory) {
		digest, err := hashInput(root, spec.source, spec.inline)
		if err != nil {
			return nil, err
		}
		inputs = append(inputs, HashCacheInput{
			Key: spec.key, Source: spec.source,
			SHA256: digest, PublicSafe: true,
		})
	}
	return inputs, nil
}

type hashCacheSpec struct {
	key    string
	source string
	inline bool
}

func hashCacheSpecs(factory repofactory.DecisionPacket) []hashCacheSpec {
	context := factory.ContextPackEvidence
	return []hashCacheSpec{
		{key: "generated_artifacts", source: "generated/external_evidence.generated.json"},
		{key: "source_descriptors", source: "generated/external_evidence.generated.json"},
		{key: "workflow_dependencies", source: ".github/workflows/quality.yml"},
		{key: "context_pack_version", source: context.ContextPackVersion, inline: true},
		{key: "ontology_version", source: context.OntologyVersion, inline: true},
	}
}

func hashInput(root string, source string, inline bool) (string, error) {
	body := []byte(source)
	if !inline {
		var err error
		body, err = os.ReadFile(filepath.Join(root, filepath.FromSlash(source)))
		if err != nil {
			return "", err
		}
	}
	sum := sha256.Sum256(body)
	return hex.EncodeToString(sum[:]), nil
}
