package daemon

import (
	"net/http"

	"github.com/kimsemi-home/myhome-jarvis/internal/linear"
)

func (server *Server) handleLinearStatus(writer http.ResponseWriter, request *http.Request) error {
	status := linear.StatusForRequest(request.Context(), server.config.Root, http.DefaultClient)
	return writeJSON(writer, http.StatusOK, linear.SummarizeStatus(status))
}

func (server *Server) handleLinearSync(writer http.ResponseWriter, request *http.Request) error {
	result := linear.PullIssues(request.Context(), server.config.Root, http.DefaultClient)
	if !result.Synced {
		if err := linear.AppendOfflineEvent(server.config.Root, "linear_sync", result.Message); err != nil {
			return err
		}
	}
	return writeJSON(writer, http.StatusOK, linear.SummarizeOperation(result))
}
