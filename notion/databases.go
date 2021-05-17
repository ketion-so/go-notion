package notion

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/ketion-so/go-notion/notion/object"
	"github.com/mitchellh/mapstructure"
)

const (
	databasesPath = "databases"
)

// DatabasesService handles communication to Notion Databases API.
//
// API doc: https://developers.notion.com/reference/database
type DatabasesService service

// Database object represents Notion Database.
//
// API doc: https://developers.notion.com/reference/database
//go:generate gomodifytags -file $GOFILE -struct Database -clear-tags -w
//go:generate gomodifytags --file $GOFILE --struct Database -add-tags json,mapstructure -w -transform snakecase
type Database struct {
	Object         object.Type         `json:"object" mapstructure:"object"`
	ID             string              `json:"id" mapstructure:"id"`
	CreatedTime    string              `json:"created_time" mapstructure:"created_time"`
	LastEditedTime string              `json:"last_edited_time" mapstructure:"last_edited_time"`
	Title          []TextObject        `json:"title" mapstructure:"title"`
	Properties     map[string]Property `json:"properties" mapstructure:"properties"`
}

func (db *Database) GetObject() object.Type {
	return db.Object
}

//go:generate gomodifytags -file $GOFILE -struct database -clear-tags -w
//go:generate gomodifytags --file $GOFILE --struct database -add-tags json,mapstructure -w -transform snakecase
type database struct {
	Object         object.Type            `json:"object" mapstructure:"object"`
	ID             string                 `json:"id" mapstructure:"id"`
	CreatedTime    string                 `json:"created_time" mapstructure:"created_time"`
	LastEditedTime string                 `json:"last_edited_time" mapstructure:"last_edited_time"`
	Title          []TextObject           `json:"title" mapstructure:"title"`
	Properties     map[string]interface{} `json:"properties" mapstructure:"properties"`
}

// Get retrieves database by database ID.
//
// API doc: https://developers.notion.com/reference/get-database
func (s *DatabasesService) Get(ctx context.Context, databaseID string) (*Database, error) {
	resp, err := s.client.get(ctx, fmt.Sprintf("%s/%s", databasesPath, databaseID))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	data := database{}
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, err
	}

	return convDatabase(&data)
}

// DatabaseQuery is a query for database
type DatabaseQuery struct {
	Filter      map[CompoundFilterType]FilterObject `json:"filter,omitempty" mapstructure:"filter"`
	Sorts       []Sort                              `json:"sort,omitempty" mapstructure:"sort"`
	StartCursor string                              `json:"start_cursor,omitempty" mapstructure:"start_cursor"`
	PageSize    int32                               `json:"page_size,omitempty" mapstructure:"page_size"`
}

type FilterObject interface{}

// TextFilter filters text properties.
//go:generate gomodifytags --file $GOFILE --struct database -add-tags json,mapstructure -w -transform snakecase -add-options json=omitempty
type TextFilter struct {
	Property       string `json:"property" mapstructure:"property"`
	Equals         string `json:"equals,omitempty" mapstructure:"equals"`
	DoesNotEqual   string `json:"does_not_equal,omitempty" mapstructure:"does_not_equal"`
	Contains       string `json:"contains,omitempty" mapstructure:"contains"`
	DoesNotContain string `json:"does_not_contain,omitempty" mapstructure:"does_not_contain"`
	StartsWith     string `json:"starts_with,omitempty" mapstructure:"starts_with"`
	EndsWith       string `json:"ends_with,omitempty" mapstructure:"ends_with"`
	IsEmpty        bool   `json:"is_empty,omitempty" mapstructure:"is_empty"`
	IsNotEmpty     bool   `json:"is_not_empty,omitempty" mapstructure:"is_not_empty"`
}

// NumberFilter filters number properties.
type NumberFilter struct {
	Property             string  `json:"property" mapstructure:"property"`
	Equals               float64 `json:"equals,omitempty" mapstructure:"equals"`
	GreaterThan          float64 `json:"greater_than,omitempty" mapstructure:"greater_than"`
	LessThan             float64 `json:"less_than,omitempty" mapstructure:"less_than"`
	GreaterThanOrEqualTo float64 `json:"greater_than_or_equal_to,omitempty" mapstructure:"greater_than_or_equal_to"`
	LessThanOrEqualTo    float64 `json:"less_than_or_equal_to,omitempty" mapstructure:"less_than_or_equal_to"`
	IsEmpty              bool    `json:"is_empty,omitempty" mapstructure:"is_empty"`
	IsNotEmpty           bool    `json:"is_not_empty,omitempty" mapstructure:"is_not_empty"`
}

// CheckboxFilter filters object properties.
type CheckboxFilter struct {
	Property    string `json:"property" mapstructure:"property"`
	Equals      bool   `json:"equals,omitempty" mapstructure:"equals"`
	DoeNotEqual bool   `json:"does_not_equal,omitempty" mapstructure:"does_not_equal"`
}

// SelectFilter filters select properties.
type SelectFilter struct {
	Property    string `json:"property" mapstructure:"property"`
	Equals      string `json:"equals,omitempty" mapstructure:"equals"`
	DoeNotEqual string `json:"does_not_equal,omitempty" mapstructure:"does_not_equal"`
	IsEmpty     bool   `json:"is_empty,omitempty" mapstructure:"is_empty"`
	IsNotEmpty  bool   `json:"is_not_empty,omitempty" mapstructure:"is_not_empty"`
}

// MultiSelectFilter filters select properties.
type MultiSelectFilter struct {
	Property    string `json:"property" mapstructure:"property"`
	Equals      string `json:"equals,omitempty" mapstructure:"equals"`
	DoeNotEqual string `json:"does_not_equal,omitempty" mapstructure:"does_not_equal"`
	IsEmpty     bool   `json:"is_empty,omitempty" mapstructure:"is_empty"`
	IsNotEmpty  bool   `json:"is_not_empty,omitempty" mapstructure:"is_not_empty"`
}

// DataFilter filters data properties.
type DataFilter struct {
	Property   string      `json:"property" mapstructure:"property"`
	Equals     string      `json:"equals,omitempty" mapstructure:"equals"`
	Before     string      `json:"before,omitempty" mapstructure:"before"`
	After      string      `json:"after,omitempty" mapstructure:"after"`
	OnOrBefore string      `json:"on_or_before,omitempty" mapstructure:"on_or_before"`
	IsEmpty    bool        `json:"is_empty,omitempty" mapstructure:"is_empty"`
	IsNotEmpty bool        `json:"is_not_empty,omitempty" mapstructure:"is_not_empty"`
	OnOrAfter  string      `json:"on_or_after,omitempty" mapstructure:"on_or_after"`
	PassWeek   interface{} `json:"pass_week,omitempty" mapstructure:"pass_week"`
	PassMonth  interface{} `json:"pass_month,omitempty" mapstructure:"pass_month"`
	PassYear   interface{} `json:"pass_year,omitempty" mapstructure:"pass_year"`
	NextWeek   interface{} `json:"next_weeb,omitempty" mapstructure:"next_weeb"`
	NextMonth  interface{} `json:"next_month,omitempty" mapstructure:"next_month"`
	NextYear   interface{} `json:"next_year,omitempty" mapstructure:"next_year"`
}

// PeopleFilter filters select properties.
type PeopleFilter struct {
	Property    string `json:"property" mapstructure:"property"`
	Equals      string `json:"equals,omitempty" mapstructure:"equals"`
	DoeNotEqual string `json:"does_not_equal,omitempty" mapstructure:"does_not_equal"`
	IsEmpty     bool   `json:"is_empty,omitempty" mapstructure:"is_empty"`
	IsNotEmpty  bool   `json:"is_not_empty,omitempty" mapstructure:"is_not_empty"`
}

// FilesFilter filters select properties.
type FilesFilter struct {
	Property   string `json:"property" mapstructure:"property"`
	IsEmpty    bool   `json:"is_empty,omitempty" mapstructure:"is_empty"`
	IsNotEmpty bool   `json:"is_not_empty,omitempty" mapstructure:"is_not_empty"`
}

// RelationFilter filters select properties.
type RelationFilter struct {
	Property    string `json:"property" mapstructure:"property"`
	Equals      string `json:"equals,omitempty" mapstructure:"equals"`
	DoeNotEqual string `json:"does_not_equal,omitempty" mapstructure:"does_not_equal"`
	IsEmpty     bool   `json:"is_empty,omitempty" mapstructure:"is_empty"`
	IsNotEmpty  bool   `json:"is_not_empty,omitempty" mapstructure:"is_not_empty"`
}

// FormulaFilter filters select properties.
type FormulaFilter struct {
	Property string          `json:"property" mapstructure:"property"`
	Text     *TextFilter     `json:"equals,omitempty" mapstructure:"equals"`
	Checkbox *CheckboxFilter `json:"does_not_equal,omitempty" mapstructure:"does_not_equal"`
	Number   *NumberFilter   `json:"is_empty,omitempty" mapstructure:"is_empty"`
	Date     *DataFilter     `json:"is_not_empty,omitempty" mapstructure:"is_not_empty"`
}

// CompoundFilterType is a type for compound filters.
type CompoundFilterType string

const (
	OrFilter  CompoundFilterType = "or"
	AndFilter CompoundFilterType = "and"
)

// QueryDatabaseResults object represents Notion Search params
//go:generate gomodifytags -file $GOFILE -struct QueryDatabaseResults -clear-tags -w
//go:generate gomodifytags --file $GOFILE --struct QueryDatabaseResults -add-tags json,mapstructure -w -transform snakecase
type QueryDatabaseResults struct {
	HasMore    bool            `json:"has_more" mapstructure:"has_more"`
	NextCursor string          `json:"next_cursor" mapstructure:"next_cursor"`
	Object     object.Type     `json:"object" mapstructure:"object"`
	Results    []object.Object `json:"results" mapstructure:"results"`
}

type queryDatabaseResults struct {
	HasMore    bool          `json:"has_more" mapstructure:"has_more"`
	NextCursor string        `json:"next_cursor" mapstructure:"next_cursor"`
	Object     object.Type   `json:"object" mapstructure:"object"`
	Results    []interface{} `json:"results" mapstructure:"results"`
}

// Query queries a database.
//
// API doc: https://developers.notion.com/reference/post-databases-query
func (s *DatabasesService) Query(ctx context.Context, databaseID string, query *DatabaseQuery) (*QueryDatabaseResults, error) {
	resp, err := s.client.post(ctx, fmt.Sprintf("%s/%s/query", databasesPath, databaseID), query)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	data := &queryDatabaseResults{}
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

	return &QueryDatabaseResults{
		HasMore:    data.HasMore,
		NextCursor: data.NextCursor,
		Object:     data.Object,
		Results:    objects,
	}, nil
}

// ListDatabase object represents Notion ListDatabase.
//
// API doc: https://developers.notion.com/reference/database
//go:generate gomodifytags -file $GOFILE -struct ListDatabase -clear-tags -w
//go:generate gomodifytags --file $GOFILE --struct ListDatabase -add-tags json,mapstructure -w -transform snakecase
type ListDatabase struct {
	Object         object.Type            `json:"object" mapstructure:"object"`
	ID             string                 `json:"id" mapstructure:"id"`
	CreatedTime    string                 `json:"created_time" mapstructure:"created_time"`
	LastEditedTime string                 `json:"last_edited_time" mapstructure:"last_edited_time"`
	Title          string                 `json:"title" mapstructure:"title"`
	Properties     map[string]interface{} `json:"properties" mapstructure:"properties"`
}

// ListDatabaseResponse represents the response from the list User API
//
//go:generate gomodifytags -file $GOFILE -struct ListDatabaseResponse -clear-tags -w
//go:generate gomodifytags --file $GOFILE --struct ListDatabaseResponse -add-tags json,mapstructure -w -transform snakecase
type ListDatabaseResponse struct {
	Object     object.Type    `json:"object" mapstructure:"object"`
	Results    []ListDatabase `json:"results" mapstructure:"results"`
	NextCursor string         `json:"next_cursor" mapstructure:"next_cursor"`
	HasMore    bool           `json:"has_more" mapstructure:"has_more"`
}

// List lists database.
//
// API doc: https://developers.notion.com/reference/get-databases
func (s *DatabasesService) List(ctx context.Context) (*ListDatabaseResponse, error) {
	resp, err := s.client.get(ctx, databasesPath)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	results := &ListDatabaseResponse{}
	if err := json.NewDecoder(resp.Body).Decode(results); err != nil {
		return nil, err
	}

	return results, nil
}

func convDatabase(data *database) (*Database, error) {
	properties, err := convProperties(data.Properties)
	if err != nil {
		return nil, err
	}

	return &Database{
		Object:         data.Object,
		ID:             data.ID,
		Title:          data.Title,
		CreatedTime:    data.CreatedTime,
		LastEditedTime: data.LastEditedTime,
		Properties:     properties,
	}, nil
}
