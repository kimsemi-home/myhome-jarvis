package daemon

import (
	"net/http"

	"github.com/kimsemi-home/myhome-jarvis/internal/domain"
)

func (server *Server) handleDomainSummary(writer http.ResponseWriter, request *http.Request) error {
	summary, err := domain.BuildSummary(server.config.Root)
	if err != nil {
		return err
	}
	return writeJSON(writer, http.StatusOK, summary)
}

func (server *Server) handleHouseholdSummary(writer http.ResponseWriter, request *http.Request) error {
	summary, err := domain.BuildSummary(server.config.Root)
	if err != nil {
		return err
	}
	return writeJSON(writer, http.StatusOK, summary.Household)
}

func (server *Server) handleRecommendationsSummary(writer http.ResponseWriter, request *http.Request) error {
	summary, err := domain.BuildSummary(server.config.Root)
	if err != nil {
		return err
	}
	return writeJSON(writer, http.StatusOK, summary.Recommendations)
}
