package notion

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/ketion-so/go-notion/notion/object"
	"github.com/mitchellh/mapstructure"
)

const (
	pagesPath = "pages"
)

// PagesService handles communication to Notion Pages API.
//
// API doc: https://developers.notion.com/reference/page
type PagesService service

// Page object represents the retrieve page.
//
// API doc: https://developers.notion.com/reference/get-page
//go:generate gomodifytags --file $GOFILE --struct Page -add-tags json,mapstructure -w -transform snakecase
type Page struct {
	Object         object.Type         `json:"object" mapstructure:"object"`
	ID             string              `json:"id" mapstructure:"id"`
	CreatedTime    string              `json:"created_time" mapstructure:"created_time"`
	LastEditedTime string              `json:"last_edited_time" mapstructure:"last_edited_time"`
	Parent         Parent              `json:"parent" mapstructure:"parent"`
	Properties     map[string]Property `json:"properties" mapstructure:"properties"`
}

func (p *Page) GetObject() object.Type {
	return p.Object
}

type page struct {
	Object         object.Type            `json:"object"`
	ID             string                 `json:"id"`
	CreatedTime    string                 `json:"created_time"`
	LastEditedTime string                 `json:"last_edited_time"`
	Parent         map[string]interface{} `json:"parent"`
	Properties     map[string]interface{} `json:"properties"`
}

// Parent represens the interface for all parents of the page.
type Parent interface {
	GetType() object.ParentType
}

// DatabaseParent object represents the retrieve parent.
//go:generate gomodifytags --file $GOFILE --struct DatabaseParent -add-tags json,mapstructure -w -transform snakecase
type DatabaseParent struct {
	Type       object.ParentType `json:"type,omitempty" mapstructure:"type"`
	DatabaseID string            `json:"database_id" mapstructure:"database_id"`
}

// GetType returns the ty
// GetType returns the type of the parent.pe of the parent.
func (p *DatabaseParent) GetType() object.ParentType {
	return object.ParentType(p.Type)
}

// PageParent object represents the retrieve parent.
//go:generate gomodifytags -file $GOFILE -struct PageParent -clear-tags -w
//go:generate gomodifytags --file $GOFILE --struct PageParent -add-tags json,mapstructure -w -transform snakecase
type PageParent struct {
	Type   object.ParentType `json:"type,omitempty" mapstructure:"type"`
	PageID string            `json:"page_id" mapstructure:"page_id"`
}

// GetType returns the type of the parent.
func (p *PageParent) GetType() object.ParentType {
	return object.ParentType(p.Type)
}

// WorkspaceParent object represents the retrieve parent.
//go:generate gomodifytags -file $GOFILE -struct WorkspaceParent -clear-tags -w
//go:generate gomodifytags --file $GOFILE --struct WorkspaceParent -add-tags json,mapstructure -w -transform snakecase
type WorkspaceParent struct {
	Type      object.ParentType `json:"type,omitempty" mapstructure:"type"`
	Workspace bool              `json:"workspace" mapstructure:"workspace"`
}

// GetType returns the type of the parent.
func (p *WorkspaceParent) GetType() object.ParentType {
	return object.ParentType(p.Type)
}

// Get retrieves a page.
//
// API doc: https://developers.notion.com/reference/get-page
func (s *PagesService) Get(ctx context.Context, pageID string) (*Page, error) {
	resp, err := s.client.get(ctx, fmt.Sprintf("%s/%s", pagesPath, pageID))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	data := page{}
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, err
	}

	return convPage(&data)
}

// CreatePageRequest object represents the retrieve page.
//go:generate gomodifytags --file $GOFILE --struct CreatePageRequest -add-tags json,mapstructure -w -transform snakecase
type CreatePageRequest struct {
	Parent     Parent              `json:"parent" mapstructure:"parent"`
	Properties map[string]Property `json:"properties" mapstructure:"properties"`
	Children   []Block             `json:"children,omitempty" mapstructure:"children"`
}

// Create page.
//
// API doc: https://developers.notion.com/reference/post-page
func (s *PagesService) Create(ctx context.Context, preq *CreatePageRequest) (*Page, error) {
	resp, err := s.client.post(ctx, pagesPath, preq)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	data := page{}
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, err
	}

	return convPage(&data)
}

// UpdatePageRequest object represents the update request
type UpdatePageRequest struct {
	Properties map[string]Property `json:"properties" mapstructure:"properties"`
}

// UpdateProperties page properties.
//
// API doc: https://developers.notion.com/reference/patch-page
func (s *PagesService) UpdateProperties(ctx context.Context, pageID string, ureq *UpdatePageRequest) (*Page, error) {
	resp, err := s.client.patch(ctx, fmt.Sprintf("%s/%s", pagesPath, pageID), ureq)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	data := page{}
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, err
	}

	return convPage(&data)
}

func convPage(data *page) (*Page, error) {
	var p Parent
	switch object.ParentType(data.Parent["type"].(string)) {
	case object.DatabaseParentType:
		p = &DatabaseParent{}
	case object.PageParentType:
		p = &PageParent{}
	case object.WorkspaceParentType:
		p = &WorkspaceParent{}
	default:
		return nil, errors.New("not type found for parent properties")
	}

	if err := mapstructure.Decode(data.Parent, &p); err != nil {
		return nil, err
	}

	properties, err := convProperties(data.Properties)
	if err != nil {
		return nil, err
	}

	page := &Page{
		Object:         data.Object,
		ID:             data.ID,
		CreatedTime:    data.CreatedTime,
		LastEditedTime: data.LastEditedTime,
		Properties:     properties,
		Parent:         p,
	}

	return page, nil
}
