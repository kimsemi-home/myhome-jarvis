package daemon

import (
	"net/http"

	"github.com/kimsemi-home/myhome-jarvis/internal/authority"
)

func (server *Server) handleAuthorityReviewStatus(
	writer http.ResponseWriter,
	request *http.Request,
) error {
	status, err := authority.ReviewPlanForRoot(server.config.Root)
	if err != nil {
		return err
	}
	return writeJSON(writer, http.StatusOK, status)
}
