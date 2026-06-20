package daemon

import (
	"net/http"

	"github.com/kimsemi-home/myhome-jarvis/internal/contextpack"
)

func (server *Server) handleContextPackStatus(writer http.ResponseWriter, request *http.Request) error {
	status, err := contextpack.StatusForRoot(server.config.Root)
	if err != nil {
		return err
	}
	return writeJSON(writer, http.StatusOK, status)
}
