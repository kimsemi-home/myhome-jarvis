package daemon

import (
	"net/http"

	"github.com/kimsemi-home/myhome-jarvis/internal/evidence"
)

func (server *Server) handleEvidenceIntegrityStatus(
	writer http.ResponseWriter,
	request *http.Request,
) error {
	status, err := evidence.IntegrityForRoot(server.config.Root)
	if err != nil {
		return err
	}
	return writeJSON(writer, http.StatusOK, status)
}
