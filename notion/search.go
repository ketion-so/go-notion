package notion

import (
	"context"
	"encoding/json"

	"github.com/ketion-so/go-notion/notion/object"
	"github.com/mitchellh/mapstructure"
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
	HasMore    bool            `json:"has_more" mapstructure:"has_more"`
	NextCursor string          `json:"next_cursor" mapstructure:"next_cursor"`
	Object     string          `json:"object" mapstructure:"object"`
	Results    []object.Object `json:"results" mapstructure:"results"`
}

type searchResults struct {
	HasMore    bool          `json:"has_more" mapstructure:"has_more"`
	NextCursor string        `json:"next_cursor" mapstructure:"next_cursor"`
	Object     string        `json:"object" mapstructure:"object"`
	Results    []interface{} `json:"results" mapstructure:"results"`
}

// Direction is a type to specify how to sort the search results.
type Direction string

const (
	Ascending  Direction = "ascending"
	Descending Direction = "descending"
)

// Sort object represents Notion User.
//go:generate gomodifytags -file $GOFILE -struct Sort -clear-tags -w
//go:generate gomodifytags --file $GOFILE --struct Sort -add-tags json,mapstructure -w -transform snakecase
type Sort struct {
	Property  string    `json:"property" mapstructure:"property"`
	Direction Direction `json:"direction" mapstructure:"direction"`
	Timestamp string    `json:"timestamp" mapstructure:"timestamp"`
}

// FiterValue is a type for specifying what to filter
type FilterValue string

const (
	Object FilterValue = "object"
)

// FilterPropertyValue is a type for what value to filter.
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
	resp, err := s.client.post(ctx, searchPath, sreq)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	data := &searchResults{}
	if err := json.NewDecoder(resp.Body).Decode(data); err != nil {
		return nil, err
	}

	objects := []object.Object{}
	for _, result := range data.Results {
		objectType := result.(map[string]interface{})["object"].(string)
		switch object.Type(objectType) {
		case object.Database:
			var db database
			if err := mapstructure.Decode(result, &db); err != nil {
				return nil, err
			}

			database, err := convDatabase(&db)
			if err != nil {
				return nil, err
			}
			objects = append(objects, database)
		case object.Page:
			var p page
			if err := mapstructure.Decode(result, &p); err != nil {
				return nil, err
			}

			page, err := convPage(&p)
			if err != nil {
				return nil, err
			}
			objects = append(objects, page)
		}
	}

	return &SearchResults{
		HasMore:    data.HasMore,
		NextCursor: data.NextCursor,
		Object:     data.Object,
		Results:    objects,
	}, nil
}
