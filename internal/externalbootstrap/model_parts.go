package externalbootstrap

type SkeletonFile struct {
	Role           string `json:"role"`
	Path           string `json:"path"`
	SourceArtifact string `json:"source_artifact"`
	Purpose        string `json:"purpose"`
	State          string `json:"state"`
}

type HashCacheInput struct {
	Key        string `json:"key"`
	Source     string `json:"source"`
	SHA256     string `json:"sha256"`
	PublicSafe bool   `json:"public_safe"`
}
