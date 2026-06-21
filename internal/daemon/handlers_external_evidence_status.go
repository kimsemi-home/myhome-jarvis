package daemon

import (
	"net/http"

	"github.com/kimsemi-home/myhome-jarvis/internal/externalevidence"
)

func (server *Server) handleExternalEvidenceStatus(
	writer http.ResponseWriter,
	request *http.Request,
) error {
	status, err := externalevidence.StatusForRoot(server.config.Root)
	if err != nil {
		return err
	}
	return writeJSON(writer, http.StatusOK, status)
}
