package notion

import (
	"context"
	"fmt"
	"net/http"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
)

func getPageJSON() string {
	return `{
		"object": "page",
		"id": "b55c9c91-384d-452b-81db-d1ef79372b75",
		"created_time": "2020-03-17T19:10:04.968Z",
		"last_edited_time": "2020-03-17T21:49:37.913Z",
		"properties": {
		  "Name": [
			{
			  "id": "some-property-id",
			  "text": "Avocado", 
			  "annotations": {
				"formatting": [],
				"color": "default",
				"link": null
			  }, 
			  "inline_object": null
			}
		  ],
		  "Description": [
			{
			  "text": "Persea americana", 
			  "annotations": {
				"formatting": [],
				"color": "default",
				"link": null
			  }, 
			  "inline_object": null
			}
		  ],
		  "In stock": false,
		  "Food group": {
			"name": "🍎Fruit",
			"color": "red"
		  },
		  "Price": 2,
		  "Cost of next trip": 2,
		  "Last ordered": "2020-03-10",
		  "Meals": [
			"a91e35b0-5c4e-4018-83e8-584988caee1c",
			"f5051efa-a7d9-4075-97f3-8ce9af14b1a7"
		  ],
		  "Number of meals": 2,
		  "Store availability": [
			{
			  "name": "Rainbow Grocery",
			  "color": "purple"
			},
			{
			  "name": "Gus's Community Market",
			  "color": "green"
			}
		  ],
		  "+1": [
			{
			  "object": "user",
			  "id": "01da9b00-e400-4959-91ce-af55307647e5",
			  "type": "person",
			  "name": "Avocado Lovelace",
			  "person": {
				"email": "avo@example.org"
			  },
			  "avatar_url": "https://secure.notion-static.com/e6a352a8-8381-44d0-a1dc-9ed80e62b53d.jpg"
			}
		  ],
		  "Photos": [
			{
			  "url": "https://s3.us-west-2.amazonaws.com/secure.notion-static.com/e6a352a8-8381-44d0-a1dc-9ed80e62b53d/avocado.jpg",
			  "name": "avocado",
			  "mime_type": "image/jpg"
			}
		  ]
		}
	  }`
}

func TestPagesService_Get(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	tcs := map[string]struct {
		id   string
		want *Page
	}{
		"ok": {
			"d40e767c-d7af-4b18-a86d-55c61f1e39a4",
			&Page{
				Object:         "page",
				ID:             "b55c9c91-384d-452b-81db-d1ef79372b75",
				CreatedTime:    "2020-03-17T19:10:04.968Z",
				LastEditedTime: "2020-03-17T21:49:37.913Z",
			},
		},
	}

	for n, tc := range tcs {
		t.Run(n, func(t *testing.T) {
			mux.HandleFunc(fmt.Sprintf("/%s/%s", pagesPath, tc.id), func(w http.ResponseWriter, r *http.Request) {
				if r.Header.Get(notionVersionHeader) == "" {
					t.Fatalf("no notion version header to request")
				}

				fmt.Fprint(w, getPageJSON())
			})

			got, err := client.Pages.Get(context.Background(), tc.id)
			if err != nil {
				t.Fatalf("Failed: %v", err)
			}

			if diff := cmp.Diff(got, tc.want, cmpopts.IgnoreFields(*got, "Properties")); diff != "" {
				t.Fatalf("Diff: %s(-got +want)", diff)
			}
		})
	}
}
