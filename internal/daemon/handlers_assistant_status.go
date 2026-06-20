package daemon

import (
	"net/http"

	"github.com/kimsemi-home/myhome-jarvis/internal/commandcenter"
)

func (server *Server) handleAssistantStatus(writer http.ResponseWriter, request *http.Request) error {
	status, err := commandcenter.StatusForRoot(server.config.Root)
	if err != nil {
		return err
	}
	return writeJSON(writer, http.StatusOK, status)
}

func (server *Server) handleWorkItemStatus(writer http.ResponseWriter, request *http.Request) error {
	status, err := commandcenter.WorkItemForRoot(server.config.Root)
	if err != nil {
		return err
	}
	return writeJSON(writer, http.StatusOK, status)
}
