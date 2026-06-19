package daemon

import (
	"errors"
	"net"
	"net/http"
	"strings"
	"time"

	"github.com/kimsemi-home/myhome-jarvis/internal/auth"
)

type handlerFunc func(http.ResponseWriter, *http.Request) error

func (server *Server) wrap(next handlerFunc) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		started := time.Now()
		recorder := &statusRecorder{ResponseWriter: writer}
		var handlerErr error
		defer func() {
			server.recordRequestEvent(request, recorder.statusCode(), started, handlerErr)
		}()

		server.requests.Add(1)
		if err := server.authorize(request); err != nil {
			handlerErr = err
			writeError(recorder, http.StatusUnauthorized, err)
			return
		}
		if err := next(recorder, request); err != nil {
			handlerErr = err
			writeError(recorder, http.StatusBadRequest, err)
		}
	}
}

func (server *Server) authorize(request *http.Request) error {
	host, _, err := net.SplitHostPort(request.RemoteAddr)
	if err != nil {
		host = request.RemoteAddr
	}
	parsed := net.ParseIP(host)
	if parsed == nil || parsed.IsLoopback() {
		return nil
	}
	token, err := readLocalToken(server.config.Root)
	if err != nil {
		return errors.New("local token is required for non-localhost clients")
	}
	header := strings.TrimSpace(request.Header.Get("Authorization"))
	if header != "Bearer "+token {
		return errors.New("invalid local token")
	}
	return nil
}

func readLocalToken(root string) (string, error) {
	return auth.Read(root)
}
