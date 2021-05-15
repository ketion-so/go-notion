package notion

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
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
//go:generate gomodifytags -file $GOFILE -struct Page -clear-tags -w
//go:generate gomodifytags --file $GOFILE --struct Page -add-tags json -w -transform snakecase
type Page struct {
	Object         string      `json:"object"`
	ID             string      `json:"id"`
	CreatedTime    string      `json:"created_time"`
	LastEditedTime string      `json:"last_edited_time"`
	Properties     interface{} `json:"properties"`
}

// List page list.
//
// API doc: https://developers.notion.com/reference/get-page
func (s *PagesService) Get(ctx context.Context, pageID string) (*Page, error) {
	req, err := s.client.NewGetRequest(fmt.Sprintf("%s/%s", pagesPath, pageID))
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

	page := &Page{}
	if err := json.NewDecoder(resp.Body).Decode(page); err != nil {
		return nil, err
	}

	return page, nil
}

// CreatePageRequest object represents the retrieve page.
//go:generate gomodifytags -file $GOFILE -struct CreatePageRequest -clear-tags -w
//go:generate gomodifytags --file $GOFILE --struct CreatePageRequest -add-tags json -w -transform snakecase
type CreatePageRequest struct {
	Parent     *Parent       `json:"parent"`
	Properties interface{}   `json:"properties"`
	Children   []interface{} `json:"children"`
}

// Parent represents database parent or page parent.
//go:generate gomodifytags --file $GOFILE --struct Parent -add-tags json -w -transform snakecase
type Parent struct {
	Type       string `json:"type"`
	DatabaseID string `json:"database_id,omitempty"`
	PageID     string `json:"page_id,omitempty"`
}

// Create page.
//
// API doc: https://developers.notion.com/reference/post-page
func (s *PagesService) Create(ctx context.Context, pageID string, preq *CreatePageRequest) (*Page, error) {
	req, err := s.client.NewPostRequest(fmt.Sprintf("%s/%s", pagesPath, pageID), preq)
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

	page := &Page{}
	if err := json.NewDecoder(resp.Body).Decode(page); err != nil {
		return nil, err
	}

	return page, nil
}

// Updateupdates page properties.
//
// API doc: https://developers.notion.com/reference/patch-page
func (s *PagesService) Update(ctx context.Context, pageID string, properties interface{}) (*Page, error) {
	req, err := s.client.NewPostRequest(fmt.Sprintf("%s/%s", pagesPath, pageID), properties)
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

	page := &Page{}
	if err := json.NewDecoder(resp.Body).Decode(page); err != nil {
		return nil, err
	}

	return page, nil
}
