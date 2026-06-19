package daemon

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/kimsemi-home/myhome-jarvis/internal/commands"
)

type harnessRequest struct {
	Name string `json:"name"`
}

func (server *Server) handleHarnessRun(writer http.ResponseWriter, request *http.Request) error {
	var body harnessRequest
	if err := decodeBody(request, &body); err != nil {
		return err
	}
	switch strings.TrimSpace(strings.ToLower(body.Name)) {
	case "", "home":
		return server.writeHarnessReport(writer, commands.RunHomeHarness())
	case "finance":
		return server.writeHarnessReport(writer, commands.RunFinanceHarness(server.config.Root))
	case "commerce":
		return server.writeHarnessReport(writer, commands.RunCommerceHarness(server.config.Root))
	default:
		return fmt.Errorf("unknown harness %q", body.Name)
	}
}

func (server *Server) writeHarnessReport(writer http.ResponseWriter, report commands.HarnessReport) error {
	status := http.StatusOK
	if !report.Passed {
		status = http.StatusBadRequest
	}
	return writeJSON(writer, status, report)
}
