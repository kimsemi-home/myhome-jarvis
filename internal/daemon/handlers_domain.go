package daemon

import (
	"net/http"
)

func (server *Server) handleDomainSummary(writer http.ResponseWriter, request *http.Request) error {
	summary, err := server.buildSummary(request.Context())
	if err != nil {
		return err
	}
	return writeJSON(writer, http.StatusOK, summary)
}

func (server *Server) handleHouseholdSummary(writer http.ResponseWriter, request *http.Request) error {
	summary, err := server.buildSummary(request.Context())
	if err != nil {
		return err
	}
	return writeJSON(writer, http.StatusOK, summary.Household)
}

func (server *Server) handleRecommendationsSummary(writer http.ResponseWriter, request *http.Request) error {
	summary, err := server.buildSummary(request.Context())
	if err != nil {
		return err
	}
	return writeJSON(writer, http.StatusOK, summary.Recommendations)
}
