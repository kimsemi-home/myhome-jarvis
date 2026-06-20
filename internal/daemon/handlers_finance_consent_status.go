package daemon

import (
	"net/http"

	"github.com/kimsemi-home/myhome-jarvis/internal/financeconsent"
)

func (server *Server) handleFinanceConsentStatus(writer http.ResponseWriter, request *http.Request) error {
	status, err := financeconsent.StatusForRoot(server.config.Root)
	if err != nil {
		return err
	}
	return writeJSON(writer, http.StatusOK, status)
}
