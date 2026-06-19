package daemon

import (
	"net/http"

	"github.com/kimsemi-home/myhome-jarvis/internal/authority"
	"github.com/kimsemi-home/myhome-jarvis/internal/controlplane"
	"github.com/kimsemi-home/myhome-jarvis/internal/evidencequality"
	"github.com/kimsemi-home/myhome-jarvis/internal/incidents"
	"github.com/kimsemi-home/myhome-jarvis/internal/review"
	"github.com/kimsemi-home/myhome-jarvis/internal/translation"
)

func (server *Server) handleTranslationStatus(writer http.ResponseWriter, request *http.Request) error {
	status, err := translation.StatusForRoot(server.config.Root)
	if err != nil {
		return err
	}
	return writeJSON(writer, http.StatusOK, status)
}

func (server *Server) handleControlPlaneStatus(writer http.ResponseWriter, request *http.Request) error {
	status, err := controlplane.StatusForRoot(server.config.Root)
	if err != nil {
		return err
	}
	return writeJSON(writer, http.StatusOK, status)
}

func (server *Server) handleIncidentsStatus(writer http.ResponseWriter, request *http.Request) error {
	status, err := incidents.StatusForRoot(server.config.Root)
	if err != nil {
		return err
	}
	return writeJSON(writer, http.StatusOK, status)
}

func (server *Server) handleEvidenceQualityStatus(writer http.ResponseWriter, request *http.Request) error {
	status, err := evidencequality.StatusForRoot(server.config.Root)
	if err != nil {
		return err
	}
	return writeJSON(writer, http.StatusOK, status)
}

func (server *Server) handleReviewStatus(writer http.ResponseWriter, request *http.Request) error {
	status, err := review.StatusForRoot(server.config.Root)
	if err != nil {
		return err
	}
	return writeJSON(writer, http.StatusOK, status)
}

func (server *Server) handleAuthorityStatus(writer http.ResponseWriter, request *http.Request) error {
	status, err := authority.StatusForRoot(server.config.Root)
	if err != nil {
		return err
	}
	return writeJSON(writer, http.StatusOK, status)
}
