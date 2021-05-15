package notion

import (
	"context"
	"fmt"
	"net/http"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/ketion-so/go-notion/notion/object"
)

func getListUserSON() string {
	return `{
"results": [
	{
		"object": "user",
		"id": "d40e767c-d7af-4b18-a86d-55c61f1e39a4",
		"type": "person",
		"person": {
			"email": "avo@example.org"
		},
		"name": "Avocado Lovelace",
		"avatar_url": "https://secure.notion-static.com/e6a352a8-8381-44d0-a1dc-9ed80e62b53d.jpg"
	},
	{
		"object": "user",
		"id": "9a3b5ae0-c6e6-482d-b0e1-ed315ee6dc57",
		"type": "bot",
		"bot": {},
		"name": "Doug Engelbot",
		"avatar_url": "https://secure.notion-static.com/6720d746-3402-4171-8ebb-28d15144923c.jpg"
	}
],
"next_cursor": "fe2cc560-036c-44cd-90e8-294d5a74cebc",
"has_more": true
}`
}

func getUserJSON(id string) string {
	return fmt.Sprintf(`{
	"object": "user",
	"id": "%s",
	"type": "person",
	"person": {
		"email": "avo@example.org"
	},
	"name": "Avocado Lovelace",
	"avatar_url": "https://secure.notion-static.com/e6a352a8-8381-44d0-a1dc-9ed80e62b53d.jpg"
}`, id)
}

func TestUsersService_Get(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	tcs := map[string]struct {
		id   string
		want *User
	}{
		"ok": {
			"d40e767c-d7af-4b18-a86d-55c61f1e39a4",
			&User{
				ID:   "d40e767c-d7af-4b18-a86d-55c61f1e39a4",
				Type: object.Person,
				Person: &People{
					Email: "avo@example.org",
				},
				Name:      "Avocado Lovelace",
				AvatarURL: "https://secure.notion-static.com/e6a352a8-8381-44d0-a1dc-9ed80e62b53d.jpg",
			},
		},
	}

	for n, tc := range tcs {
		t.Run(n, func(t *testing.T) {
			mux.HandleFunc(fmt.Sprintf("/%s/%s", usersPath, tc.id), func(w http.ResponseWriter, r *http.Request) {
				if r.Header.Get(notionVersionHeader) == "" {
					t.Fatalf("no notion version header to request")
				}

				fmt.Fprint(w, getUserJSON(tc.id))
			})

			got, err := client.Users.Get(context.Background(), tc.id)
			if err != nil {
				t.Fatalf("Failed: %v", err)
			}

			if diff := cmp.Diff(got, tc.want); diff != "" {
				t.Fatalf("Diff: %s(-got +want)", diff)
			}
		})
	}
}

func TestUsersService_List(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	tcs := map[string]struct {
		want *ListUserResponse
	}{
		"ok": {
			&ListUserResponse{
				Results: []User{
					{
						ID:   "d40e767c-d7af-4b18-a86d-55c61f1e39a4",
						Type: object.Person,
						Person: &People{
							Email: "avo@example.org",
						},
						Name:      "Avocado Lovelace",
						AvatarURL: "https://secure.notion-static.com/e6a352a8-8381-44d0-a1dc-9ed80e62b53d.jpg",
					},
					{
						ID:        "9a3b5ae0-c6e6-482d-b0e1-ed315ee6dc57",
						Type:      object.Bot,
						Bot:       &Bot{},
						Name:      "Doug Engelbot",
						AvatarURL: "https://secure.notion-static.com/6720d746-3402-4171-8ebb-28d15144923c.jpg",
					},
				},
				NextCursor: "fe2cc560-036c-44cd-90e8-294d5a74cebc",
				HasMore:    true,
			},
		},
	}

	for n, tc := range tcs {
		t.Run(n, func(t *testing.T) {
			mux.HandleFunc("/"+usersPath, func(w http.ResponseWriter, r *http.Request) {
				if r.Header.Get(notionVersionHeader) == "" {
					t.Fatalf("no notion version header to request")
				}

				fmt.Fprint(w, getListUserSON())
			})

			got, err := client.Users.List(context.Background())
			if err != nil {
				t.Fatalf("Failed: %v", err)
			}

			if diff := cmp.Diff(got, tc.want); diff != "" {
				t.Fatalf("Diff: %s(-got +want)", diff)
			}
		})
	}
}
