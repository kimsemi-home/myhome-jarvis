package auth

import "os"

func Status(root string) LocalTokenStatus {
	path := LocalTokenPath(root)
	info, err := os.Stat(path)
	if err != nil {
		return LocalTokenStatus{
			Configured: false,
			Path:       localTokenRelativePath,
			Message:    "local LAN token is not configured",
		}
	}
	return LocalTokenStatus{
		Configured: true,
		Path:       localTokenRelativePath,
		Mode:       info.Mode().Perm().String(),
		Message:    "local LAN token is configured",
	}
}
