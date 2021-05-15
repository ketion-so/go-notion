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
//go:generate gomodifytags --file $GOFILE --struct SearchRequest -add-tags json,mapstructure -w -transform snakecase
type SearchRequest struct {
	Query       string `json:"query" mapstructure:"query"`
	Sort        *Sort  `json:"sort" mapstructure:"sort"`
	StartCursor string `json:"start_cursor" mapstructure:"start_cursor"`
	PageSize    int32  `json:"page_size" mapstructure:"page_size"`
}

// SearchResults object represents Notion Search params
//go:generate gomodifytags -file $GOFILE -struct SearchResults -clear-tags -w
//go:generate gomodifytags --file $GOFILE --struct SearchResults -add-tags json,mapstructure -w -transform snakecase
type SearchResults struct {
	HasMore    bool       `json:"has_more" mapstructure:"has_more"`
	NextCursor string     `json:"next_cursor" mapstructure:"next_cursor"`
	Object     string     `json:"object" mapstructure:"object"`
	Results    []Database `json:"results" mapstructure:"results"`
}

type searchResults struct {
	HasMore    bool          `json:"has_more" mapstructure:"has_more"`
	NextCursor string        `json:"next_cursor" mapstructure:"next_cursor"`
	Object     string        `json:"object" mapstructure:"object"`
	Results    []interface{} `json:"results" mapstructure:"results"`
}

type Direction string

const (
	Ascending  Direction = "ascending"
	Descending Direction = "descending"
)

// Sort object represents Notion User.
//go:generate gomodifytags -file $GOFILE -struct Sort -clear-tags -w
//go:generate gomodifytags --file $GOFILE --struct Sort -add-tags json,mapstructure -w -transform snakecase
type Sort struct {
	Direction Direction `json:"direction" mapstructure:"direction"`
	Timestamp string    `json:"timestamp" mapstructure:"timestamp"`
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
//go:generate gomodifytags --file $GOFILE --struct Filter -add-tags json,mapstructure -w -transform snakecase
type Filter struct {
	Value    FilterValue         `json:"value" mapstructure:"value"`
	Property FilterPropertyValue `json:"property" mapstructure:"property"`
}

// Get gets user by user ID.
//
// API doc: https://developers.notion.com/reference/get-user
func (s *SearchService) Search(ctx context.Context, sreq *SearchRequest) (*SearchResults, error) {
	resp, err := s.client.Post(ctx, searchPath, sreq)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		respErr := &Error{}
		if err := json.NewDecoder(resp.Body).Decode(respErr); err != nil {
			return nil, err
		}
		return nil, fmt.Errorf("status code not expected, got:%d, message:%s", resp.StatusCode, respErr.Message)
	}

	data := &searchResults{}
	if err := json.NewDecoder(resp.Body).Decode(data); err != nil {
		return nil, err
	}

	databases := []Database{}
	for _, result := range data.Results {
		switch v := result.(type) {
		case *database:
			database, err := convDatabase(v)
			if err != nil {
				return nil, err
			}
			databases = append(databases, *database)
		}
	}

	return &SearchResults{
		HasMore:    data.HasMore,
		NextCursor: data.NextCursor,
		Object:     data.Object,
		Results:    databases,
	}, nil
}
