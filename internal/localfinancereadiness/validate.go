package localfinancereadiness

import (
	"errors"
	"path/filepath"
	"regexp"
	"strings"
)

var hashPattern = regexp.MustCompile(`^[a-f0-9]{64}$`)

func Validate(value Manifest) error {
	if value.SchemaVersion != ManifestSchema || value.ExecutionMode != "plan_only" ||
		value.CredentialsRead || value.ExternalNetworkEnabled || value.ExternalWritesEnabled ||
		value.InstallAllowed || value.Timezone != "Asia/Seoul" || len(value.Plans) != 3 || len(value.Stages) != 4 {
		return errors.New("local finance readiness manifest enables execution or is incomplete")
	}
	seen := map[string]bool{}
	for _, ref := range value.Plans {
		clean := filepath.Clean(ref.Path)
		if seen[ref.Component] || !expectedComponent(ref.Component) || filepath.IsAbs(ref.Path) ||
			clean == ".." || strings.HasPrefix(clean, ".."+string(filepath.Separator)) ||
			!hashPattern.MatchString(ref.ArtifactSHA256) || !hashPattern.MatchString(ref.PlanHash) {
			return errors.New("local finance readiness reference is invalid")
		}
		seen[ref.Component] = true
	}
	if len(seen) != 3 || validateStages(value.Stages) != nil || value.AggregateHash != aggregateHash(value) {
		return errors.New("local finance readiness DAG or aggregate hash is invalid")
	}
	return nil
}

func expectedComponent(value string) bool {
	return value == "ledger" || value == "portfolio" || value == "revenue"
}
