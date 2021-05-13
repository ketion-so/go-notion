package notion

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
)

const (
	baseURLPath   = "/v1"
	testAccessKey = "notion-test"
)

func getErrorJSON(status int) string {
	return fmt.Sprintf(`{
	"status": %d,
	"message": "error",
	"type": "error"
}`, status)
}

func addHeader(w http.ResponseWriter) {
	w.Header().Add(rateLimitRemainingHeader, "99")
	w.Header().Add(rateLimitLimitHeader, "1000")
	w.Header().Add(rateLimitResetHeader, "1598795193")
}

func setup() (*Client, *http.ServeMux, string, func()) {
	mux := http.NewServeMux()

	apiHandler := http.NewServeMux()
	apiHandler.Handle(baseURLPath+"/", http.StripPrefix(baseURLPath, mux))
	server := httptest.NewServer(apiHandler)
	client := NewClient(testAccessKey)
	url, _ := url.Parse(server.URL + baseURLPath)
	client.BaseURL = url
	return client, mux, server.URL, server.Close
}
