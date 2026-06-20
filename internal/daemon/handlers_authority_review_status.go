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

func (server *Server) handleAuthorityReviewRequest(
	writer http.ResponseWriter,
	request *http.Request,
) error {
	packet, err := authority.ReviewRequestPacketForRoot(server.config.Root)
	if err != nil {
		return err
	}
	return writeJSON(writer, http.StatusOK, packet)
}

func (server *Server) handleAuthorityReviewEvidence(
	writer http.ResponseWriter,
	request *http.Request,
) error {
	status, err := authority.ReviewRequestEvidenceForRoot(server.config.Root)
	if err != nil {
		return err
	}
	return writeJSON(writer, http.StatusOK, status)
}

func (server *Server) handleAuthorityReviewQueue(
	writer http.ResponseWriter,
	request *http.Request,
) error {
	status, err := authority.ReviewQueueStatusForRoot(server.config.Root)
	if err != nil {
		return err
	}
	return writeJSON(writer, http.StatusOK, status)
}
