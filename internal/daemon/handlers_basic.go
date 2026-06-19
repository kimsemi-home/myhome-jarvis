package daemon

import (
	"net/http"
	"time"

	"github.com/kimsemi-home/myhome-jarvis/internal/auth"
	"github.com/kimsemi-home/myhome-jarvis/internal/commands"
)

func (server *Server) handleHealth(writer http.ResponseWriter, request *http.Request) error {
	return writeJSON(writer, http.StatusOK, map[string]any{
		"ok":       true,
		"mode":     "local",
		"dry_run":  true,
		"host":     server.config.Host,
		"started":  server.started.Format(time.RFC3339),
		"requests": server.requests.Load(),
	})
}

func (server *Server) handleVersion(writer http.ResponseWriter, request *http.Request) error {
	return writeJSON(writer, http.StatusOK, map[string]any{
		"name":    "myhome-jarvis",
		"version": server.config.Version,
	})
}

func (server *Server) handleAuthStatus(writer http.ResponseWriter, request *http.Request) error {
	return writeJSON(writer, http.StatusOK, auth.Status(server.config.Root))
}

func (server *Server) handleCommands(writer http.ResponseWriter, request *http.Request) error {
	return writeJSON(writer, http.StatusOK, commands.Specs())
}
