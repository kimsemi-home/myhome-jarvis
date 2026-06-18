package agentcluster

import (
	"fmt"
	"sort"
)

var requiredLifecycleStages = []string{
	"observed",
	"evidence_recorded",
	"classified",
	"owner_assigned",
	"verified",
	"knowledge_updated",
}

func requireOrdered(values []string, required ...string) error {
	last := -1
	for _, item := range required {
		index := indexOf(values, item)
		if index < 0 {
			return fmt.Errorf("agent cluster evidence flow missing %q", item)
		}
		if index <= last {
			return fmt.Errorf("agent cluster evidence flow must keep %q after the previous stage", item)
		}
		last = index
	}
	return nil
}

func normalizeList(values []string) []string {
	seen := map[string]bool{}
	clean := make([]string, 0, len(values))
	for _, value := range values {
		item := normalizeToken(value)
		if item == "" || seen[item] {
			continue
		}
		seen[item] = true
		clean = append(clean, item)
	}
	sort.Strings(clean)
	return clean
}

func normalizeOrderedList(values []string) []string {
	seen := map[string]bool{}
	clean := make([]string, 0, len(values))
	for _, value := range values {
		item := normalizeToken(value)
		if item == "" || seen[item] {
			continue
		}
		seen[item] = true
		clean = append(clean, item)
	}
	return clean
}

func contains(values []string, wanted string) bool {
	return indexOf(values, wanted) >= 0
}

func indexOf(values []string, wanted string) int {
	for index, value := range values {
		if value == wanted {
			return index
		}
	}
	return -1
}
