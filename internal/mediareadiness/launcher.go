package mediareadiness

import (
	"os/exec"
	"runtime"
)

func localLauncherAvailable() (bool, string) {
	if runtime.GOOS != "darwin" {
		return false, "requires_darwin"
	}
	if _, err := exec.LookPath("open"); err != nil {
		return false, "open_command_missing"
	}
	return true, "open_command_available"
}
