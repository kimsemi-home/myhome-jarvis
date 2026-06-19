package linear

import (
	"io"
	"net/http"
	"strconv"
	"strings"
)

func linearResponse(remaining int, body string) *http.Response {
	return &http.Response{
		StatusCode: http.StatusOK,
		Header:     http.Header{"X-RateLimit-Remaining": []string{strconv.Itoa(remaining)}},
		Body:       io.NopCloser(strings.NewReader(body)),
	}
}
