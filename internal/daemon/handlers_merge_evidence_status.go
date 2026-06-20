package daemon

import (
	"net/http"

	"github.com/kimsemi-home/myhome-jarvis/internal/mergeevidence"
)

func (server *Server) handleMergeEvidenceStatus(
	writer http.ResponseWriter,
	request *http.Request,
) error {
	status, err := mergeevidence.StatusForRoot(server.config.Root)
	if err != nil {
		return err
	}
	return writeJSON(writer, http.StatusOK, status)
}
