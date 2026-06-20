package daemon

import (
	"net/http"

	"github.com/kimsemi-home/myhome-jarvis/internal/mediareadiness"
)

func (server *Server) handleMediaReadinessStatus(
	writer http.ResponseWriter,
	request *http.Request,
) error {
	status, err := mediareadiness.StatusForRoot(server.config.Root)
	if err != nil {
		return err
	}
	return writeJSON(writer, http.StatusOK, status)
}
