package controlplane

import (
	"path/filepath"
	"strings"
	"time"
)

func normalizeManifest(policy Policy, request ManifestRequest) (Manifest, error) {
	manifest := Manifest{
		At:               time.Now().UTC().Format(time.RFC3339),
		DecisionKind:     normalizeToken(request.DecisionKind),
		PolicyVersion:    publicText(request.PolicyVersion),
		OntologyVersion:  publicText(request.OntologyVersion),
		AuthorityProfile: normalizeToken(request.AuthorityProfile),
		SelectedRoute:    normalizeToken(request.SelectedRoute),
		ReviewerRole:     normalizeToken(request.ReviewerRole),
		VerifierRole:     normalizeToken(request.VerifierRole),
		LeaseSeconds:     request.LeaseSeconds,
		LeaseStatus:      normalizeToken(request.LeaseStatus),
		EvidenceRefs:     normalizeRefs(request.EvidenceRefs),
		OutputRef:        filepath.ToSlash(strings.TrimSpace(request.OutputRef)),
	}
	if err := validateManifest(policy, manifest); err != nil {
		return Manifest{}, err
	}
	manifest.ID = manifestID(manifest)
	return manifest, nil
}
