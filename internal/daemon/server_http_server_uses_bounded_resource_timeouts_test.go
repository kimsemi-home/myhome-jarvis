package daemon

import (
	"testing"
	"time"
)

func TestHTTPServerUsesBoundedResourceTimeouts(t *testing.T) {
	server, err := New(Config{Root: t.TempDir(), Port: 3888, Version: "test"})
	if err != nil {
		t.Fatal(err)
	}
	httpServer := server.httpServer("127.0.0.1:0", server.Routes())

	if httpServer.ReadHeaderTimeout != 5*time.Second {
		t.Fatalf("read header timeout = %s", httpServer.ReadHeaderTimeout)
	}
	if httpServer.ReadTimeout != 15*time.Second {
		t.Fatalf("read timeout = %s", httpServer.ReadTimeout)
	}
	if httpServer.WriteTimeout != 30*time.Second {
		t.Fatalf("write timeout = %s", httpServer.WriteTimeout)
	}
	if httpServer.IdleTimeout != 60*time.Second {
		t.Fatalf("idle timeout = %s", httpServer.IdleTimeout)
	}
	if httpServer.MaxHeaderBytes != 1<<20 {
		t.Fatalf("max header bytes = %d", httpServer.MaxHeaderBytes)
	}
}
