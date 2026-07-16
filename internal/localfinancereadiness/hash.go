package localfinancereadiness

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
)

func planHash(value Plan) string {
	value.PlanHash = ""
	body, _ := json.Marshal(value)
	return digest(body)
}

func aggregateHash(value Manifest) string {
	value.AggregateHash = ""
	body, _ := json.Marshal(value)
	return digest(body)
}

func digest(body []byte) string {
	hash := sha256.Sum256(body)
	return hex.EncodeToString(hash[:])
}
