package mediareadiness

func availabilityLabel(available bool) string {
	if available {
		return "available"
	}
	return "degraded"
}
