package notion

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/ketion-so/go-notion/notion/object"
)

const (
	usersPath = "users"
)

// UsersService handles communication to Notion Users API.
//
// API doc: https://developers.notion.com/reference/user
type UsersService service

// User object represents Notion User.
//
// API doc: https://developers.notion.com/reference/user
//go:generate gomodifytags --file $GOFILE --struct User -add-tags json,mapstructure -w -transform snakecase
type User struct {
	ID        string      `json:"id" mapstructure:"id"`
	Type      object.Type `json:"type,omitempty" mapstructure:"type"`
	Name      string      `json:"name,omitempty" mapstructure:"name"`
	AvatarURL string      `json:"avatar_url,omitempty" mapstructure:"avatar_url"`
	Person    *People     `json:"person,omitempty" mapstructure:"person"`
	Bot       *Bot        `json:"bot,omitempty" mapstructure:"bot"`
}

// People object represents Notion human account.
//
//go:generate gomodifytags -file $GOFILE -struct People -clear-tags -w
//go:generate gomodifytags --file $GOFILE --struct People -add-tags json,mapstructure -w -transform snakecase
type People struct {
	Email string `json:"email" mapstructure:"email"`
}

// Bot object represents Notion bot account.
//
//go:generate gomodifytags -file $GOFILE -struct Bot -clear-tags -w
//go:generate gomodifytags --file $GOFILE --struct Bot -add-tags json,mapstructure -w -transform snakecase
type Bot struct{}

// Get gets user by user ID.
//
// API doc: https://developers.notion.com/reference/get-user
func (s *UsersService) Get(ctx context.Context, userID string) (*User, error) {
	resp, err := s.client.get(ctx, fmt.Sprintf("%s/%s", usersPath, userID))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	user := &User{}
	if err := json.NewDecoder(resp.Body).Decode(user); err != nil {
		return nil, err
	}

	return user, nil
}

// ListUserResponse represents the response from the list User API
//
//go:generate gomodifytags -file $GOFILE -struct ListUserResponse -clear-tags -w
//go:generate gomodifytags --file $GOFILE --struct ListUserResponse -add-tags json,mapstructure -w -transform snakecase
type ListUserResponse struct {
	Object     object.Type `json:"object" mapstructure:"object"`
	Results    []User      `json:"results" mapstructure:"results"`
	NextCursor string      `json:"next_cursor" mapstructure:"next_cursor"`
	HasMore    bool        `json:"has_more" mapstructure:"has_more"`
}

// List gets the list of users.
//
// API doc: https://developers.notion.com/reference/get-users
func (s *UsersService) List(ctx context.Context) (*ListUserResponse, error) {
	resp, err := s.client.get(ctx, usersPath)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	luResp := &ListUserResponse{}
	if err := json.NewDecoder(resp.Body).Decode(luResp); err != nil {
		return nil, err
	}

	return luResp, nil
}
