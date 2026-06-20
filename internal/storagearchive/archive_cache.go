package storagearchive

import (
	"bufio"
	"encoding/json"
	"os"
	"path/filepath"
)

type archiveCacheKey struct {
	SourceKey            string
	InputSHA256          string
	ConfigEvidenceSHA256 string
}

type archiveCacheEntry struct {
	ArchivePath string
}

type archiveCache map[archiveCacheKey]archiveCacheEntry

func readArchiveCache(root string, manifestPath string) (archiveCache, error) {
	cache := archiveCache{}
	path := filepath.Join(root, filepath.FromSlash(manifestPath))
	file, err := os.Open(path)
	if os.IsNotExist(err) {
		return cache, nil
	}
	if err != nil {
		return archiveCache{}, err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		applyArchiveCacheLine(cache, scanner.Bytes())
	}
	if err := scanner.Err(); err != nil {
		return archiveCache{}, err
	}
	return cache, nil
}

func applyArchiveCacheLine(cache archiveCache, line []byte) {
	var entry manifestEntry
	if err := json.Unmarshal(line, &entry); err != nil {
		return
	}
	if entry.State != "archived" ||
		entry.ArchivePath == "" ||
		entry.InputSHA256 == "" ||
		entry.ConfigEvidenceSHA256 == "" {
		return
	}
	key := archiveCacheKey{
		SourceKey:            entry.SourceKey,
		InputSHA256:          entry.InputSHA256,
		ConfigEvidenceSHA256: entry.ConfigEvidenceSHA256,
	}
	cache[key] = archiveCacheEntry{ArchivePath: entry.ArchivePath}
}

func (cache archiveCache) hit(
	sourceKey string,
	inputSHA256 string,
	configEvidenceSHA256 string,
) (archiveCacheEntry, bool) {
	key := archiveCacheKey{sourceKey, inputSHA256, configEvidenceSHA256}
	entry, ok := cache[key]
	return entry, ok
}
