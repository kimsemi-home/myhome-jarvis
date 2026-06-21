package externalbootstrap

import "sort"

func requiredHashCacheKeys(packet Packet) []string {
	return sortedHashCacheKeys(packet.HashCacheInputs)
}

func sortedHashCacheKeys(inputs []HashCacheInput) []string {
	keys := make([]string, 0, len(inputs))
	seen := map[string]bool{}
	for _, input := range inputs {
		if input.Key == "" || seen[input.Key] {
			continue
		}
		seen[input.Key] = true
		keys = append(keys, input.Key)
	}
	sort.Strings(keys)
	return keys
}

func hashCacheByKey(inputs []HashCacheInput) map[string]HashCacheInput {
	result := map[string]HashCacheInput{}
	for _, input := range inputs {
		result[input.Key] = input
	}
	return result
}
