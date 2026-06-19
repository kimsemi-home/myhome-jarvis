package daemon

import (
	"net"
	"net/http"
	"strconv"

	"github.com/kimsemi-home/myhome-jarvis/internal/supervisor"
)

func (server *Server) ListenAndServe() error {
	mux := server.Routes()
	address := net.JoinHostPort(server.config.Host, strconv.Itoa(server.config.Port))
	httpServer := server.httpServer(address, mux)
	listener, err := net.Listen("tcp", address)
	if err != nil {
		return err
	}
	state, err := supervisor.NewDaemonState(
		server.config.Root,
		server.config.Host,
		server.config.Port,
		server.config.Version,
		server.config.Execute,
		server.config.AllowLANBind,
	)
	if err != nil {
		_ = listener.Close()
		return err
	}
	if _, err := supervisor.WriteDaemonState(server.config.Root, state); err != nil {
		_ = listener.Close()
		return err
	}
	return httpServer.Serve(listener)
}

func (server *Server) httpServer(address string, handler http.Handler) *http.Server {
	return &http.Server{
		Addr:              address,
		Handler:           handler,
		ReadHeaderTimeout: server.config.ReadHeaderTimeout,
		ReadTimeout:       server.config.ReadTimeout,
		WriteTimeout:      server.config.WriteTimeout,
		IdleTimeout:       server.config.IdleTimeout,
		MaxHeaderBytes:    server.config.MaxHeaderBytes,
	}
}
