package storagearchive

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"

	"github.com/kimsemi-home/myhome-jarvis/internal/domain"
)

type configEvidenceRef struct {
	Field  string
	Inputs []string
	SHA256 string
}

type configEvidencePayload struct {
	Inputs              []string                   `json:"inputs"`
	LogArchive          domain.LogArchivePolicy    `json:"log_archive"`
	EvidenceNoiseBudget domain.EvidenceNoiseBudget `json:"evidence_noise_budget"`
}

func configEvidenceRefForPolicy(policy domain.StoragePolicy) configEvidenceRef {
	inputs := append([]string{}, policy.LogArchive.ConfigHashInputs...)
	payload := configEvidencePayload{
		Inputs:              inputs,
		LogArchive:          policy.LogArchive,
		EvidenceNoiseBudget: policy.EvidenceNoiseBudget,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return configEvidenceRef{Field: policy.EvidenceNoiseBudget.ConfigEvidenceField}
	}
	sum := sha256.Sum256(body)
	return configEvidenceRef{
		Field:  policy.EvidenceNoiseBudget.ConfigEvidenceField,
		Inputs: inputs,
		SHA256: hex.EncodeToString(sum[:]),
	}
}
