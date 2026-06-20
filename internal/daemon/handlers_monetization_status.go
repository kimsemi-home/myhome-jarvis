package daemon

import (
	"net/http"

	"github.com/kimsemi-home/myhome-jarvis/internal/monetization"
)

func (server *Server) handleMonetizationStatus(writer http.ResponseWriter, request *http.Request) error {
	status, err := monetization.StatusForRoot(server.config.Root)
	if err != nil {
		return err
	}
	return writeJSON(writer, http.StatusOK, status)
}
