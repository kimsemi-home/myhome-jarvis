package daemon

import (
	"net/http"

	"github.com/kimsemi-home/myhome-jarvis/internal/financeconnector"
)

func (server *Server) handleFinanceParity(writer http.ResponseWriter, request *http.Request) error {
	loaded, err := financeconnector.LoadConfigured(request.Context(), server.config.Root)
	if err != nil {
		return err
	}
	return writeJSON(writer, http.StatusOK, loaded.Parity)
}
