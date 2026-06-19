package daemon

import (
	"net/http"
	"path/filepath"
	"runtime"
	"time"
)

func (server *Server) handleMetrics(writer http.ResponseWriter, request *http.Request) error {
	events := server.events.snapshot()
	var memory runtime.MemStats
	runtime.ReadMemStats(&memory)
	return writeJSON(writer, http.StatusOK, map[string]any{
		"started":           server.started.Format(time.RFC3339),
		"uptime_seconds":    int64(time.Since(server.started).Seconds()),
		"requests":          server.requests.Load(),
		"event_count":       len(events),
		"goroutine_count":   runtime.NumGoroutine(),
		"heap_alloc_bytes":  memory.HeapAlloc,
		"heap_sys_bytes":    memory.HeapSys,
		"stack_inuse_bytes": memory.StackInuse,
		"gc_count":          memory.NumGC,
		"execute_enabled":   server.config.Execute,
		"bind_host":         server.config.Host,
		"linear_queue":      filepath.ToSlash(filepath.Join("data", "private", "linear-offline-queue.jsonl")),
		"dry_run_default":   true,
		"lan_bind_allowed":  server.config.AllowLANBind,
	})
}

func (server *Server) handleEvents(writer http.ResponseWriter, request *http.Request) error {
	events := server.events.snapshot()
	return writeJSON(writer, http.StatusOK, map[string]any{
		"count":  len(events),
		"events": events,
	})
}
