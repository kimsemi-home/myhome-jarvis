package supervisor

import (
	"path/filepath"
)

const (
	daemonStateName         = "daemon"
	daemonStateRelativePath = "data/private/supervisor/daemon-state.json"
)

func statePath(root string) string {
	return filepath.Join(root, filepath.FromSlash(daemonStateRelativePath))
}
