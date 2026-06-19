package evidencequality

import "fmt"

func validateSnapshotEnums(policy Policy, snapshot Snapshot) error {
	if !contains(normalizeList(policy.AllowedPurposes), snapshot.Purpose) {
		return fmt.Errorf("evidence quality purpose %q is not allowed", snapshot.Purpose)
	}
	if !contains(normalizeList(policy.QualityLevels), snapshot.QualityLevel) {
		return fmt.Errorf("evidence quality level %q is not allowed", snapshot.QualityLevel)
	}
	if !contains(normalizeList(policy.MappingConfidenceLevels), snapshot.MappingConfidence) {
		return fmt.Errorf("evidence quality mapping confidence %q is not allowed", snapshot.MappingConfidence)
	}
	if snapshot.SchemaVersion == "" ||
		snapshot.OntologyVersion == "" ||
		snapshot.AssessedBy == "" {
		return fmt.Errorf("evidence quality schema, ontology, and assessor fields are required")
	}
	for _, reason := range snapshot.ReassessmentReasons {
		if !contains(normalizeList(policy.ReassessmentReasons), reason) {
			return fmt.Errorf("evidence quality reassessment reason %q is not allowed", reason)
		}
	}
	return nil
}
