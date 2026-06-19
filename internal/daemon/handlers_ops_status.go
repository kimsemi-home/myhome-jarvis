package daemon

import (
	"net/http"

	"github.com/kimsemi-home/myhome-jarvis/internal/audit"
	"github.com/kimsemi-home/myhome-jarvis/internal/planner"
	"github.com/kimsemi-home/myhome-jarvis/internal/qualitylog"
	"github.com/kimsemi-home/myhome-jarvis/internal/supervisor"
)

func (server *Server) handleSupervisorStatus(writer http.ResponseWriter, request *http.Request) error {
	return writeJSON(writer, http.StatusOK, supervisor.Status(server.config.Root, nil))
}

func (server *Server) handleAuditStatus(writer http.ResponseWriter, request *http.Request) error {
	status, err := audit.CommandIntentStatusForRoot(server.config.Root)
	if err != nil {
		return err
	}
	return writeJSON(writer, http.StatusOK, status)
}

func (server *Server) handleQualityStatus(writer http.ResponseWriter, request *http.Request) error {
	status, err := qualitylog.StatusForRoot(server.config.Root)
	if err != nil {
		return err
	}
	return writeJSON(writer, http.StatusOK, status)
}

func (server *Server) handlePlannerStatus(writer http.ResponseWriter, request *http.Request) error {
	status, err := planner.StatusForRoot(server.config.Root)
	if err != nil {
		return err
	}
	return writeJSON(writer, http.StatusOK, status)
}
