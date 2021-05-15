package notion

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/ketion-so/go-notion/notion/object"
)

const (
	usersPath = "users"
)

// UsersService handles communication to Notion Users API.
//
// API doc: https://developers.notion.com/reference/user
type UsersService service

// Users object represents Notion User.
//
// API doc: https://developers.notion.com/reference/user
//go:generate gomodifytags --file $GOFILE --struct User -add-tags json -w -transform snakecase
type User struct {
	ID        string      `json:"id"`
	Type      object.Type `json:"type"`
	Name      string      `json:"name"`
	AvatarURL string      `json:"avatar_url"`
	Person    *People     `json:"person,omitempty"`
	Bot       *Bot        `json:"bot,omitempty"`
}

// People object represents Notion human account.
//
//go:generate gomodifytags -file $GOFILE -struct People -clear-tags -w
//go:generate gomodifytags --file $GOFILE --struct People -add-tags json -w -transform snakecase
type People struct {
	Email string `json:"email"`
}

// Bot object represents Notion bot account.
//
//go:generate gomodifytags -file $GOFILE -struct Bot -clear-tags -w
//go:generate gomodifytags --file $GOFILE --struct Bot -add-tags json -w -transform snakecase
type Bot struct{}

// Get gets user by user ID.
//
// API doc: https://developers.notion.com/reference/get-user
func (s *UsersService) Get(ctx context.Context, userID string) (*User, error) {
	req, err := s.client.NewGetRequest(fmt.Sprintf("%s/%s", usersPath, userID))
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

	user := &User{}
	if err := json.NewDecoder(resp.Body).Decode(user); err != nil {
		return nil, err
	}

	return user, nil
}

// ListUserResponse represents the response from the list User API
//
//go:generate gomodifytags -file $GOFILE -struct ListUserResponse -clear-tags -w
//go:generate gomodifytags --file $GOFILE --struct ListUserResponse -add-tags json -w -transform snakecase
type ListUserResponse struct {
	Object     object.Type `json:"object"`
	Results    []User      `json:"results"`
	NextCursor string      `json:"next_cursor"`
	HasMore    bool        `json:"has_more"`
}

// List gets the list of users.
//
// API doc: https://developers.notion.com/reference/get-users
func (s *UsersService) List(ctx context.Context) (*ListUserResponse, error) {
	req, err := s.client.NewGetRequest(usersPath)
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

	luResp := &ListUserResponse{}
	if err := json.NewDecoder(resp.Body).Decode(luResp); err != nil {
		return nil, err
	}

	return luResp, nil
}
