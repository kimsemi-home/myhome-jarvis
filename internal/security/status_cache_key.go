package security

import (
	"crypto/sha256"
	"encoding/hex"
	"hash"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)

func buildStatusCacheKey(root string) (statusCacheKey, error) {
	head, err := gitOutput(root, "rev-parse", "HEAD")
	if err != nil {
		return statusCacheKey{}, err
	}
	sum := sha256.New()
	writeHashPart(sum, "version", statusCacheVersion)
	writeHashPart(sum, "head", strings.TrimSpace(head))
	hashRelativeFile(root, "generated/security.generated.json", sum)
	if err := hashSecuritySource(root, sum); err != nil {
		return statusCacheKey{}, err
	}
	inputHash := hex.EncodeToString(sum.Sum(nil))
	return statusCacheKey{
		Head:      strings.TrimSpace(head),
		InputHash: inputHash,
		Key:       "security-history:" + strings.TrimSpace(head) + ":" + inputHash[:16],
	}, nil
}

func hashSecuritySource(root string, sum hash.Hash) error {
	dir := filepath.Join(root, "internal", "security")
	return filepath.WalkDir(dir, func(path string, entry fs.DirEntry, walkErr error) error {
		if os.IsNotExist(walkErr) {
			writeHashPart(sum, "missing-dir", "internal/security")
			return filepath.SkipDir
		}
		if walkErr != nil || entry.IsDir() || filepath.Ext(path) != ".go" {
			return walkErr
		}
		rel, err := filepath.Rel(root, path)
		if err != nil {
			return err
		}
		hashRelativeFile(root, filepath.ToSlash(rel), sum)
		return nil
	})
}

func hashRelativeFile(root string, rel string, sum hash.Hash) {
	content, err := os.ReadFile(filepath.Join(root, filepath.FromSlash(rel)))
	if err != nil {
		writeHashPart(sum, "missing-file", rel)
		return
	}
	writeHashPart(sum, rel, string(content))
}

func writeHashPart(sum hash.Hash, key string, value string) {
	_, _ = sum.Write([]byte(key))
	_, _ = sum.Write([]byte{0})
	_, _ = sum.Write([]byte(value))
	_, _ = sum.Write([]byte{0})
}
