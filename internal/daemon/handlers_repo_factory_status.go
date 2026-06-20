package daemon

import (
	"net/http"

	"github.com/kimsemi-home/myhome-jarvis/internal/repofactory"
)

func (server *Server) handleRepoFactoryStatus(writer http.ResponseWriter, request *http.Request) error {
	status, err := repofactory.StatusForRoot(server.config.Root)
	if err != nil {
		return err
	}
	return writeJSON(writer, http.StatusOK, status)
}
