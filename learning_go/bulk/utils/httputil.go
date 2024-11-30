package utils

import (
	"net/http"
	"time"
)

type ContextRoundTripper struct {
	rt http.RoundTripper
}

func HttpClient() (*http.Client, error) {
	client := &http.Client{
		Timeout:   30 * time.Second,
		Transport: &ContextRoundTripper{},
	}

	return client, nil
}
func (crt *ContextRoundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	if crt.rt == nil {
		crt.rt = http.DefaultTransport
	}

	if requestId, ok := req.Context().Value("X-Request-Id").(string); ok {
		req.Header.Set("X-Request-Id", requestId)
	}

	if entityId, ok := req.Context().Value("X-Entity-Id").(string); ok {
		req.Header.Set("X-Entity-Id", entityId)
	}

	return crt.rt.RoundTrip(req)
}
