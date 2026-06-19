package linear

import (
	"crypto/sha256"
	"encoding/hex"
)

func offlineEntryID(line string) string {
	sum := sha256.Sum256([]byte(line))
	return hex.EncodeToString(sum[:])
}
