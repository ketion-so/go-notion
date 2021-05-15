package notion

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

const (
	databasesPath = "databases"
)

// Databaseservice handles communication to Notion Databases API.
//
// API doc: https://developers.notion.com/reference/database
type DatabasesService service

// Database object represents Notion Database.
//
// API doc: https://developers.notion.com/reference/database
//go:generate gomodifytags -file $GOFILE -struct Database -clear-tags -w
//go:generate gomodifytags --file $GOFILE --struct Database -add-tags json -w -transform snakecase
type Database struct {
	Object         string                 `json:"object"`
	ID             string                 `json:"id"`
	CreatedTime    string                 `json:"created_time"`
	LastEditedTime string                 `json:"last_edited_time"`
	Title          []RichText             `json:"title"`
	Properties     map[string]interface{} `json:"properties"`
}

type DatabasePropertyType string

const (
	TitleDatabaseType          DatabasePropertyType = "title"
	RichTextDatabaseType       DatabasePropertyType = "rich_text"
	NumberDatabaseType         DatabasePropertyType = "number"
	SelectDatabaseType         DatabasePropertyType = "select"
	MultiSelectDatabaseType    DatabasePropertyType = "multi_select"
	DateDatabaseType           DatabasePropertyType = "date"
	PeopleDatabaseType         DatabasePropertyType = "people"
	FileDatabaseType           DatabasePropertyType = "file"
	CheckboxDatabaseType       DatabasePropertyType = "checkbox"
	URLDatabaseType            DatabasePropertyType = "url"
	EmailDatabaseType          DatabasePropertyType = "email"
	PhoneNumberDatabaseType    DatabasePropertyType = "phone_number"
	FormulaDatabaseType        DatabasePropertyType = "formula"
	RelationDatabaseType       DatabasePropertyType = "relation"
	RollupDatabaseType         DatabasePropertyType = "rollup"
	CreatedTypeDatabaseType    DatabasePropertyType = "created_time"
	CreatedByDatabaseType      DatabasePropertyType = "created_by"
	LastEditedTimeDatabaseType DatabasePropertyType = "last_edited_time"
	LastEditedByDatabaseType   DatabasePropertyType = "last_edited_by"
)

// Get retrieves database by database ID.
//
// API doc: https://developers.notion.com/reference/get-database
func (s *DatabasesService) Get(ctx context.Context, databaseID string) (*Database, error) {
	req, err := s.client.NewGetRequest(fmt.Sprintf("%s/%s", databasesPath, databaseID))
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

	db := &Database{}
	if err := json.NewDecoder(resp.Body).Decode(db); err != nil {
		return nil, err
	}

	return db, nil
}
