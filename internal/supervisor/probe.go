package supervisor

import (
	"net"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func probeHealth(url string, client *http.Client) (bool, int) {
	if strings.TrimSpace(url) == "" {
		return false, 0
	}
	if client == nil {
		client = &http.Client{Timeout: 500 * time.Millisecond}
	}
	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return false, 0
	}
	response, err := client.Do(request)
	if err != nil {
		return false, 0
	}
	defer response.Body.Close()
	return response.StatusCode >= 200 && response.StatusCode < 300, response.StatusCode
}

func healthURL(state DaemonState) string {
	host := strings.TrimSpace(state.Host)
	switch host {
	case "", "0.0.0.0":
		host = "127.0.0.1"
	case "::":
		host = "::1"
	}
	if state.Port <= 0 {
		return ""
	}
	return "http://" + net.JoinHostPort(host, strconv.Itoa(state.Port)) + "/health"
}
