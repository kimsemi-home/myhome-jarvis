package daemon

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

func decodeBody(request *http.Request, target any) error {
	defer request.Body.Close()
	decoder := json.NewDecoder(io.LimitReader(request.Body, 1<<20))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(target); err != nil {
		if errors.Is(err, io.EOF) {
			return nil
		}
		return fmt.Errorf("invalid json body: %w", err)
	}
	return nil
}

func writeJSON(writer http.ResponseWriter, status int, value any) error {
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(status)
	encoder := json.NewEncoder(writer)
	encoder.SetIndent("", "  ")
	return encoder.Encode(value)
}

func writeError(writer http.ResponseWriter, status int, err error) {
	_ = writeJSON(writer, status, map[string]any{
		"ok":    false,
		"error": err.Error(),
	})
}
