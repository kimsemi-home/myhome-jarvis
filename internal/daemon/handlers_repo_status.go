package daemon

import (
	"net/http"

	"github.com/kimsemi-home/myhome-jarvis/internal/codeshape"
	"github.com/kimsemi-home/myhome-jarvis/internal/repo"
	"github.com/kimsemi-home/myhome-jarvis/internal/scheduler"
	"github.com/kimsemi-home/myhome-jarvis/internal/security"
)

func (server *Server) handleRepoStatus(writer http.ResponseWriter, request *http.Request) error {
	status, err := repo.Inspect(server.config.Root)
	if err != nil {
		return err
	}
	return writeJSON(writer, http.StatusOK, status)
}

func (server *Server) handleSecurityStatus(writer http.ResponseWriter, request *http.Request) error {
	status, err := security.StatusForRoot(server.config.Root)
	if err != nil {
		return err
	}
	return writeJSON(writer, http.StatusOK, status)
}

func (server *Server) handleCodeShapeStatus(writer http.ResponseWriter, request *http.Request) error {
	status, err := codeshape.StatusForRoot(server.config.Root)
	if err != nil {
		return err
	}
	return writeJSON(writer, http.StatusOK, status)
}

func (server *Server) handleLoopStatus(writer http.ResponseWriter, request *http.Request) error {
	status, err := scheduler.Status(server.config.Root, scheduler.ClosedLoopPolicy())
	if err != nil {
		return err
	}
	return writeJSON(writer, http.StatusOK, status)
}
