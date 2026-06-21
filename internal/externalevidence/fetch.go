package externalevidence

import (
	"fmt"
	"io"
	"net/http"
)

type fetchResult struct {
	Status int
	Body   []byte
}

func fetchSource(policy Policy, source SourceDescriptor, client *http.Client) (fetchResult, error) {
	request, err := http.NewRequest(http.MethodGet, source.URL, nil)
	if err != nil {
		return fetchResult{}, fmt.Errorf("request_build")
	}
	request.Header.Set("User-Agent", "myhome-jarvis-external-evidence/0.1")
	response, err := client.Do(request)
	if err != nil {
		return fetchResult{}, fmt.Errorf("network")
	}
	defer response.Body.Close()
	if response.StatusCode < 200 || response.StatusCode >= 300 {
		return fetchResult{Status: response.StatusCode}, fmt.Errorf("http_status")
	}
	body, err := readBoundedBody(response.Body, policy.CollectionMaxBytes)
	if err != nil {
		return fetchResult{Status: response.StatusCode}, err
	}
	return fetchResult{Status: response.StatusCode, Body: body}, nil
}

func readBoundedBody(reader io.Reader, maxBytes int64) ([]byte, error) {
	if maxBytes <= 0 {
		return nil, fmt.Errorf("payload_limit")
	}
	body, err := io.ReadAll(io.LimitReader(reader, maxBytes+1))
	if err != nil {
		return nil, fmt.Errorf("read_body")
	}
	if int64(len(body)) > maxBytes {
		return nil, fmt.Errorf("payload_too_large")
	}
	return body, nil
}
