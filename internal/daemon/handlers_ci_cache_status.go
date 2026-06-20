package daemon

import (
	"net/http"

	"github.com/kimsemi-home/myhome-jarvis/internal/cicache"
)

func (server *Server) handleCICacheStatus(writer http.ResponseWriter, request *http.Request) error {
	status, err := cicache.StatusForRoot(server.config.Root)
	if err != nil {
		return err
	}
	return writeJSON(writer, http.StatusOK, status)
}
