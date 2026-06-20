package daemon

import (
	"net/http"

	"github.com/kimsemi-home/myhome-jarvis/internal/codexsustainability"
)

func (server *Server) handleCodexSustainabilityStatus(writer http.ResponseWriter, request *http.Request) error {
	status, err := codexsustainability.StatusForRoot(server.config.Root)
	if err != nil {
		return err
	}
	return writeJSON(writer, http.StatusOK, status)
}
