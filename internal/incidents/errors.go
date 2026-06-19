package incidents

import "errors"

var (
	errMissingOwner    = errors.New("incident owner role is required")
	errMissingEvidence = errors.New("incident evidence refs are required")
)
