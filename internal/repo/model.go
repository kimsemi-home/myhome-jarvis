package repo

type Status struct {
	Branch              string   `json:"branch"`
	HeadSHA             string   `json:"head_sha"`
	ShortSHA            string   `json:"short_sha"`
	Upstream            string   `json:"upstream,omitempty"`
	Origin              string   `json:"origin,omitempty"`
	WorktreeClean       bool     `json:"worktree_clean"`
	TrackedChanges      []Change `json:"tracked_changes"`
	UntrackedFiles      []string `json:"untracked_files"`
	IgnoredPrivatePaths []string `json:"ignored_private_paths"`
	CheckedAt           string   `json:"checked_at"`
}

type Change struct {
	Code string `json:"code"`
	Path string `json:"path"`
}
