package daemon

import (
	"net/http"

	"github.com/kimsemi-home/myhome-jarvis/internal/agentcluster"
	"github.com/kimsemi-home/myhome-jarvis/internal/confidence"
	"github.com/kimsemi-home/myhome-jarvis/internal/connectors"
	"github.com/kimsemi-home/myhome-jarvis/internal/evidence"
	"github.com/kimsemi-home/myhome-jarvis/internal/learning"
)

func (server *Server) handleConnectorsStatus(writer http.ResponseWriter, request *http.Request) error {
	status, err := connectors.StatusForRoot(server.config.Root)
	if err != nil {
		return err
	}
	return writeJSON(writer, http.StatusOK, status)
}

func (server *Server) handleAgentClusterStatus(writer http.ResponseWriter, request *http.Request) error {
	status, err := agentcluster.StatusForRoot(server.config.Root)
	if err != nil {
		return err
	}
	return writeJSON(writer, http.StatusOK, status)
}

func (server *Server) handleLearningStatus(writer http.ResponseWriter, request *http.Request) error {
	status, err := learning.StatusForRoot(server.config.Root)
	if err != nil {
		return err
	}
	return writeJSON(writer, http.StatusOK, status)
}

func (server *Server) handleEvidenceStatus(writer http.ResponseWriter, request *http.Request) error {
	status, err := evidence.StatusForRoot(server.config.Root)
	if err != nil {
		return err
	}
	return writeJSON(writer, http.StatusOK, status)
}

func (server *Server) handleConfidenceStatus(writer http.ResponseWriter, request *http.Request) error {
	status, err := confidence.StatusForRoot(server.config.Root)
	if err != nil {
		return err
	}
	return writeJSON(writer, http.StatusOK, status)
}
