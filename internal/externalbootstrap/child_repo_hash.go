package externalbootstrap

import "regexp"

var childSHA256Pattern = regexp.MustCompile(`^[0-9a-f]{64}$`)

func checkChildRepoHashCache(root string, packet Packet, status *ChildRepoStatus) {
	document, ok := readChildJSON[childHashCacheDocument](root, ".mhj/hash-cache-inputs.json")
	if !ok {
		status.addFinding(".mhj/hash-cache-inputs.json", "missing_file",
			"hash cache input declaration is missing or invalid")
		return
	}
	checkChildString(status, ".mhj/hash-cache-inputs.json", document.Context,
		packet.Context, "hash cache context drift")
	checkChildString(status, ".mhj/hash-cache-inputs.json", document.CandidateRepo,
		packet.CandidateRepo, "hash cache candidate repo drift")
	expected := hashCacheByKey(packet.HashCacheInputs)
	observed := hashCacheByKey(document.HashCacheInputs)
	status.ObservedHashCacheKeys = sortedHashCacheKeys(document.HashCacheInputs)
	for _, key := range status.RequiredHashCacheKeys {
		item, ok := observed[key]
		if !ok {
			status.addFinding(".mhj/hash-cache-inputs.json",
				"invalid_hash_cache_missing_key", "required hash cache key is missing")
			continue
		}
		validateChildHashCacheInput(status, key, item, expected[key])
	}
	status.HashCacheValid = countChildFindings(*status, "invalid_hash_cache") == 0
}

func validateChildHashCacheInput(
	status *ChildRepoStatus,
	key string,
	observed HashCacheInput,
	expected HashCacheInput,
) {
	if !observed.PublicSafe || !childSHA256Pattern.MatchString(observed.SHA256) {
		status.addFinding(".mhj/hash-cache-inputs.json",
			"invalid_hash_cache_value", "hash cache input is not public-safe sha256")
	}
	if expected.Key != "" && observed.SHA256 != expected.SHA256 {
		status.addFinding(".mhj/hash-cache-inputs.json",
			"invalid_hash_cache_drift", "hash cache input differs from upstream packet")
	}
	if observed.Key != key {
		status.addFinding(".mhj/hash-cache-inputs.json",
			"invalid_hash_cache_key", "hash cache input key is inconsistent")
	}
}
