package notion

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

const (
	searchPath = "search"
)

// SearchService handles communication to Notion Search API.
//
// API doc: https://developers.notion.com/reference/post-search
type SearchService service

// SearchRequest object represents Notion Search params
//go:generate gomodifytags -file $GOFILE -struct SearchRequest -clear-tags -w
//go:generate gomodifytags --file $GOFILE --struct SearchRequest -add-tags json -w -transform snakecase
type SearchRequest struct {
	Query       string `json:"query"`
	Sort        *Sort  `json:"sort"`
	StartCursor string `json:"start_cursor"`
	PageSize    int32  `json:"page_size"`
}

// SearchResult object represents Notion Search params
//go:generate gomodifytags -file $GOFILE -struct SearchResult -clear-tags -w
//go:generate gomodifytags --file $GOFILE --struct SearchResult -add-tags json -w -transform snakecase
type SearchResult struct {
	HasMore    bool          `json:"has_more"`
	NextCursor string        `json:"next_cursor"`
	Object     string        `json:"object"`
	Results    []interface{} `json:"results"`
}

type Direction string

const (
	Ascending  Direction = "ascending"
	Descending Direction = "descending"
)

// Sort object represents Notion User.
//go:generate gomodifytags -file $GOFILE -struct Sort -clear-tags -w
//go:generate gomodifytags --file $GOFILE --struct Sort -add-tags json -w -transform snakecase
type Sort struct {
	Direction Direction `json:"direction"`
	Timestamp string    `json:"timestamp"`
}

type FilterValue string

const (
	Object FilterValue = "object"
)

type FilterPropertyValue string

const (
	ObjectFilterProperty FilterPropertyValue = "object"
)

// Filter object represents Notion User.
//go:generate gomodifytags -file $GOFILE -struct Filter -clear-tags -w
//go:generate gomodifytags --file $GOFILE --struct Filter -add-tags json -w -transform snakecase
type Filter struct {
	Value    FilterValue         `json:"value"`
	Property FilterPropertyValue `json:"property"`
}

// Get gets user by user ID.
//
// API doc: https://developers.notion.com/reference/get-user
func (s *SearchService) Search(ctx context.Context, sreq *SearchRequest) (*SearchResult, error) {
	req, err := s.client.NewPostRequest(searchPath, sreq)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		respErr := &RespError{}
		if err := json.NewDecoder(resp.Body).Decode(respErr); err != nil {
			return nil, err
		}
		return nil, fmt.Errorf("status code not expected, got:%d, message:%s", resp.StatusCode, respErr.Message)
	}

	result := &SearchResult{}
	if err := json.NewDecoder(resp.Body).Decode(result); err != nil {
		return nil, err
	}

	return result, nil
}
