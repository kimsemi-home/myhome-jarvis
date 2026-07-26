package daemon

import (
	"net/http"

	"github.com/kimsemi-home/myhome-jarvis/internal/financeconnector"
)

func (server *Server) handleFinanceParity(writer http.ResponseWriter, request *http.Request) error {
	report, err := financeconnector.VerifyFixture(server.config.Root)
	if err != nil {
		return err
	}
	return writeJSON(writer, http.StatusOK, report)
}
