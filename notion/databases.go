package notion

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/ketion-so/go-notion/notion/object"
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
