package notion

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
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

func TestWithHTTPClient(t *testing.T) {
	type testCase struct {
		client *http.Client
	}

	tcs := map[string]testCase{
		"ok": {
			&http.Client{
				Timeout: time.Duration(10),
			},
		},
	}

	for n, tc := range tcs {
		tc := tc
		t.Run(n, func(t *testing.T) {
			t.Parallel()

			got := NewClient(testAccessKey, WithHTTPClient(tc.client))
			if diff := cmp.Diff(got.client, tc.client); diff != "" {
				t.Fatalf("Diff: %s(-got +want)", diff)
			}
		})
	}
}

func TestWithVersion(t *testing.T) {
	type testCase struct {
		version string
	}

	tcs := map[string]testCase{
		"ok": {
			"2021-05-13",
		},
		"empty string": {
			"",
		},
	}

	for n, tc := range tcs {
		tc := tc
		t.Run(n, func(t *testing.T) {
			t.Parallel()

			c := NewClient(testAccessKey, WithVersion(tc.version))
			if c.version != tc.version {
				t.Fatalf("version not set got:%s want:%s", c.version, tc.version)
			}
		})
	}
}

func TestError_Error(t *testing.T) {
	type testCase struct {
		message string
	}

	tcs := map[string]testCase{
		"ok": {
			"Invalid request",
		},
	}

	for n, tc := range tcs {
		tc := tc
		t.Run(n, func(t *testing.T) {
			t.Parallel()

			e := &Error{
				Message: tc.message,
			}

			if e.Error() != tc.message {
				t.Fatalf("failed to get message got:%s want:%s", e.Error(), tc.message)
			}
		})
	}
}

func TesClient_Do(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	type testCase struct {
		rateLimitReset     int
		rateLimitRemaining int
		rateLimiLimit      int
	}

	tcs := map[string]testCase{
		"ok": {
			5000,
			4987,
			1350085394,
		},
		"empty rate limit header": {
			0,
			0,
			0,
		},
	}

	for n, tc := range tcs {
		t.Run(n, func(t *testing.T) {
			mux.HandleFunc(fmt.Sprintf("/%s", usersPath), func(w http.ResponseWriter, r *http.Request) {
				if r.Header.Get(notionVersionHeader) == "" {
					t.Fatalf("no notion version header to request")
				}

				if tc.rateLimiLimit != 0 && tc.rateLimitRemaining != 0 && tc.rateLimitReset != 0 {
					w.Header().Set(rateLimitResetHeader, fmt.Sprint(tc.rateLimitReset))
					w.Header().Set(rateLimitRemainingHeader, fmt.Sprint(tc.rateLimitRemaining))
					w.Header().Set(rateLimitLimitHeader, fmt.Sprint(tc.rateLimiLimit))
				}
				fmt.Fprint(w, "")
			})

			_, err := client.Users.List(context.Background())
			if err != nil {
				t.Fatalf("failed: %v", err)
			}

			if client.RateLimit.Limit != tc.rateLimiLimit {
				t.Fatalf("rate limit has not been configured got:%d, want:%d", client.RateLimit.Limit, tc.rateLimiLimit)
			}

			if client.RateLimit.Remaining != tc.rateLimitRemaining {
				t.Fatalf("rate remaning has not been configured got:%d, want:%d", client.RateLimit.Remaining, tc.rateLimitRemaining)
			}

			if client.RateLimit.Reset.Unix() != int64(tc.rateLimitReset) {
				t.Fatalf("rate reset timestamp has not been configured got:%d, want:%d", client.RateLimit.Reset.Unix(), tc.rateLimitReset)
			}
		})
	}
}
