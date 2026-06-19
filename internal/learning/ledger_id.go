package learning

import (
	"crypto/sha256"
	"encoding/hex"
	"strings"
)

func observationID(observation Observation) string {
	sum := sha256.Sum256([]byte(strings.Join([]string{
		observation.At,
		observation.Kind,
		observation.Source,
		observation.Stage,
		observation.Owner,
		observation.Summary,
	}, "\x00")))
	return "learn_" + hex.EncodeToString(sum[:])[:16]
}
