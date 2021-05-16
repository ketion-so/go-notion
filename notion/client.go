package notion

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"sync"
	"time"

	"github.com/ketion-so/go-notion/notion/object"
)

var (
	defaultRateLimit = &RateLimit{
		Limit:     10000,
		Remaining: 10000,
	}
)

const (
	rateLimitResetHeader     = "X-RateLimit-Reset"
	rateLimitRemainingHeader = "X-RateLimit-Remaining"
	rateLimitLimitHeader     = "X-RateLimit-Limit"

	baseURL          = "https://api.notion.com"
	defaultUserAgent = "go-notion"
)

type service struct {
	client *Client
}

// Client represents the API client for Notion
type Client struct {
	accessKey string
	common    service
	client    *http.Client

	mu sync.RWMutex

	RateLimit   *RateLimit
	UserAgent   string
	AccessToken string
	BaseURL     *url.URL
	version     string

	Blocks    *BlocksService
	Databases *DatabasesService
	Pages     *PagesService
	Search    *SearchService
	Users     *UsersService
}

// RateLimit represents the rate limit info for the API
type RateLimit struct {
	Limit     int
	Remaining int
	Reset     time.Time
}

// ClientOption represents options to configure this Notion API client.
type ClientOption func(c *Client)

// WithHTTPClient overrides the default http.Client.
func WithHTTPClient(client *http.Client) ClientOption {
	return func(c *Client) {
		c.client = client
	}
}

// WithVersion overrides the Notion API version to communicate.
func WithVersion(version string) ClientOption {
	return func(c *Client) {
		c.version = version
	}
}

// NewClient returns the API client for Notion
func NewClient(accessKey string, opts ...ClientOption) *Client {
	baseURL, _ := url.Parse(baseURL)
	c := &Client{
		BaseURL:   baseURL,
		accessKey: accessKey,
		UserAgent: defaultUserAgent,
		version:   defaultVersion,
		client:    http.DefaultClient,
	}

	for _, opt := range opts {
		opt(c)
	}

	c.common.client = c
	c.RateLimit = defaultRateLimit

	c.Blocks = (*BlocksService)(&c.common)
	c.Databases = (*DatabasesService)(&c.common)
	c.Pages = (*PagesService)(&c.common)
	c.Search = (*SearchService)(&c.common)
	c.Users = (*UsersService)(&c.common)
	return c
}

func (c *Client) request(ctx context.Context, method, urlStr string, body interface{}) (*http.Response, error) {
	u, err := c.BaseURL.Parse(fmt.Sprintf("v1/%s", urlStr))
	if err != nil {
		return nil, err
	}

	var buf io.ReadWriter
	if body != nil {
		buf = &bytes.Buffer{}
		enc := json.NewEncoder(buf)
		enc.SetEscapeHTML(false)
		err := enc.Encode(body)
		if err != nil {
			return nil, err
		}
	}

	req, err := http.NewRequest(method, u.String(), buf)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.accessKey))
	req.Header.Add(notionVersionHeader, c.version)

	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}

	if c.UserAgent != "" {
		req.Header.Set("User-Agent", c.UserAgent)
	}

	return c.Do(ctx, req)
}

// Get requests API GET request.
func (c *Client) Get(ctx context.Context, urlStr string) (*http.Response, error) {
	return c.request(ctx, "GET", urlStr, nil)
}

// Post requests API POST request.
func (c *Client) Post(ctx context.Context, urlStr string, body interface{}) (*http.Response, error) {
	return c.request(ctx, "POST", urlStr, body)
}

// Patch requests API Patch request.
func (c *Client) Patch(ctx context.Context, urlStr string, body interface{}) (*http.Response, error) {
	return c.request(ctx, "PATCH", urlStr, body)
}

// Delete requests API Delete request.
func (c *Client) Delete(ctx context.Context, urlStr string) (*http.Response, error) {
	return c.request(ctx, "DELETE", urlStr, nil)
}

func (c *Client) Do(ctx context.Context, req *http.Request) (*http.Response, error) {
	resp, err := c.client.Do(req.WithContext(ctx))
	if err != nil {
		return nil, err
	}

	if l := resp.Header.Get(rateLimitLimitHeader); l != "" {
		c.mu.Lock()
		limit, err := strconv.Atoi(l)
		if err != nil {
			return nil, err
		}
		c.RateLimit.Limit = limit
		c.mu.Unlock()
	}

	if r := resp.Header.Get(rateLimitRemainingHeader); r != "" {
		c.mu.Lock()
		remaining, err := strconv.Atoi(r)
		if err != nil {
			return nil, err
		}
		c.RateLimit.Remaining = remaining
		c.mu.Unlock()
	}

	if r := resp.Header.Get(rateLimitResetHeader); r != "" {
		c.mu.Lock()
		r, err := strconv.Atoi(r)
		if err != nil {
			return nil, err
		}
		c.RateLimit.Reset = time.Unix(int64(r), 0)
		c.mu.Unlock()
	}

	if resp.StatusCode != http.StatusOK {
		apiErr := &Error{}
		if err := json.NewDecoder(resp.Body).Decode(apiErr); err != nil {
			return nil, err
		}
		return nil, apiErr
	}

	return resp, nil
}

// Error represents error response from Notion
//go:generate gomodifytags -file $GOFILE -struct Error -clear-tags -w
//go:generate gomodifytags --file $GOFILE --struct Error -add-tags json,mapstructure -w -transform snakecase
type Error struct {
	Object  string           `json:"object" mapstructure:"object"`
	Status  int              `json:"status" mapstructure:"status"`
	Code    object.ErrorCode `json:"code" mapstructure:"code"`
	Message string           `json:"message" mapstructure:"message"`
}

// Error implements the error interface
func (e *Error) Error() string {
	return e.Message
}
