package commandcenter

func visionAuditRowByKey(
	audit VisionAudit,
	key string,
) VisionRequirementAudit {
	for _, row := range audit.Requirements {
		if row.CapabilityKey == key {
			return row
		}
	}
	return VisionRequirementAudit{}
}

func containsString(values []string, wanted string) bool {
	for _, value := range values {
		if value == wanted {
			return true
		}
	}
	return false
}
