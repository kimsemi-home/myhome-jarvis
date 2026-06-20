package evidence

func learningSource(policy Policy) (PrivateSource, bool) {
	for _, source := range policy.PrivateSources {
		if source.Key == "learning" {
			return source, true
		}
	}
	return PrivateSource{}, false
}
