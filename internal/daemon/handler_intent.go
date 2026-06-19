package daemon

import (
	"encoding/json"
	"net/http"

	"github.com/kimsemi-home/myhome-jarvis/internal/audit"
	"github.com/kimsemi-home/myhome-jarvis/internal/commands"
)

type intentRequest struct {
	Command string          `json:"command"`
	Payload json.RawMessage `json:"payload"`
	Execute bool            `json:"execute"`
}

func (server *Server) handleIntent(writer http.ResponseWriter, request *http.Request) error {
	var body intentRequest
	if err := decodeBody(request, &body); err != nil {
		_ = audit.AppendCommandIntent(server.config.Root, audit.CommandIntentFromPlan("daemon", "", false, commands.Plan{}, err))
		return err
	}
	if len(body.Payload) == 0 {
		body.Payload = []byte("{}")
	}
	plan, err := commands.Build(body.Command, body.Payload)
	if err != nil {
		_ = audit.AppendCommandIntent(server.config.Root, audit.CommandIntentFromPlan("daemon", body.Command, body.Execute, plan, err))
		return err
	}
	plan, err = server.finalizeIntentPlan(request, body, plan)
	if err != nil {
		return err
	}
	if err := audit.AppendCommandIntent(server.config.Root, audit.CommandIntentFromPlan("daemon", body.Command, body.Execute, plan, nil)); err != nil {
		return err
	}
	return writeJSON(writer, http.StatusOK, plan)
}

func (server *Server) finalizeIntentPlan(request *http.Request, body intentRequest, plan commands.Plan) (commands.Plan, error) {
	executeAllowed := body.Execute && server.config.Execute
	plan = commands.WithExecuteAllowed(plan, executeAllowed)
	if body.Execute && !server.config.Execute {
		plan.Warnings = append(plan.Warnings, "execute was requested but daemon execute mode is disabled")
	}
	if !executeAllowed {
		return plan, nil
	}
	plan, err := commands.Execute(request.Context(), plan, commands.ExecuteOptions{
		Platform: server.config.CommandPlatform,
		Runner:   server.config.CommandRunner,
	})
	if err != nil {
		_ = audit.AppendCommandIntent(server.config.Root, audit.CommandIntentFromPlan("daemon", body.Command, body.Execute, plan, err))
	}
	return plan, err
}
