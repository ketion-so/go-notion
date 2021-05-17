package notion

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	"github.com/ketion-so/go-notion/notion/object"
)

const (
	baseURLPath   = "/v1"
	testAccessKey = "notion-test"
)

const (
	invalidJSON = "{jo,"
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

func TestClient_Do(t *testing.T) {
	type testCase struct {
		rateLimitResetStr     string
		rateLimitRemainingStr string
		rateLimiLimitStr      string
		shouldPass            bool
	}

	tcs := map[string]testCase{
		"ok": {
			"5000",
			"4987",
			"1350085394",
			true,
		},
		"invalid_rate_limit_reset": {
			"hello",
			"4987",
			"1350085394",
			false,
		},
		"invalid_rate_limit_remaning": {
			"5000",
			"good morning",
			"1350085394",
			false,
		},
		"invalid_rate_limit_limit": {
			"5000",
			"4987",
			"Back to the future",
			false,
		},
	}

	for n, tc := range tcs {
		t.Run(n, func(t *testing.T) {
			client, mux, _, teardown := setup()
			defer teardown()

			mux.HandleFunc(fmt.Sprintf("/%s", usersPath), func(w http.ResponseWriter, r *http.Request) {
				if r.Header.Get(notionVersionHeader) == "" {
					t.Fatalf("no notion version header to request")
				}

				if tc.rateLimiLimitStr != "0" && tc.rateLimitRemainingStr != "0" && tc.rateLimitResetStr != "0" {
					w.Header().Set(rateLimitResetHeader, tc.rateLimitResetStr)
					w.Header().Set(rateLimitRemainingHeader, tc.rateLimitRemainingStr)
					w.Header().Set(rateLimitLimitHeader, tc.rateLimiLimitStr)
				}
				w.Write([]byte("{}"))
			})

			_, err := client.Users.List(context.Background())
			if err != nil {
				if tc.shouldPass {
					t.Fatalf("failed: %v", err)
				}

				return
			}

			if fmt.Sprint(client.RateLimit.Limit) != tc.rateLimiLimitStr {
				t.Fatalf("rate limit has not been configured got:%d, want:%s", client.RateLimit.Limit, tc.rateLimiLimitStr)
			}

			if fmt.Sprint(client.RateLimit.Remaining) != tc.rateLimitRemainingStr {
				t.Fatalf("rate remaning has not been configured got:%d, want:%s", client.RateLimit.Remaining, tc.rateLimitRemainingStr)
			}

			if fmt.Sprint(client.RateLimit.Reset.Unix()) != tc.rateLimitResetStr {
				t.Fatalf("rate reset timestamp has not been configured got:%d, want:%s", client.RateLimit.Reset.Unix(), tc.rateLimitResetStr)
			}
		})
	}
}

func TestClient_Do_Error(t *testing.T) {
	type testCase struct {
		want       interface{}
		shouldPass bool
	}

	tcs := map[string]testCase{
		"ok": {
			&Error{
				object.Error,
				http.StatusInternalServerError,
				object.ErrInternalServer,
				"internal server error",
			},
			true,
		},
		"invalid json": {
			invalidJSON,
			false,
		},
	}

	for n, tc := range tcs {
		tc := tc
		t.Run(n, func(t *testing.T) {
			t.Parallel()

			client, mux, _, teardown := setup()
			defer teardown()

			mux.HandleFunc(fmt.Sprintf("/%s", usersPath), func(w http.ResponseWriter, r *http.Request) {
				if r.Header.Get(notionVersionHeader) == "" {
					t.Fatalf("no notion version header to request")
				}

				w.WriteHeader(http.StatusInternalServerError)
				_ = json.NewEncoder(w).Encode(tc.want)
			})

			_, err := client.Users.List(context.Background())
			switch v := err.(type) {
			case *Error:
				if diff := cmp.Diff(v, tc.want); diff != "" {
					t.Fatalf("Diff: %s(-got +want)", diff)
				}
			default:
				if tc.shouldPass {
					t.Fatalf("failed: %v", v)
				}
			}
		})
	}
}
