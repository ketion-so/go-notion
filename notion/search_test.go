package notion

import (
	"context"
	"fmt"
	"net/http"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
)

func getSearchJSON() string {
	return `{
		"has_more": false,
		"next_cursor": null,
		"object": "list",
		"results": [
			{
				"created_time": "2021-04-22T22:23:26.080Z",
				"id": "e6c6f8ff-c70e-4970-91ba-98f03e0d7fc6",
				"last_edited_time": "2021-04-23T04:21:00.000Z",
				"object": "database",
				"properties": {
					"Name": {
						"id": "title",
						"title": {},
						"type": "title"
					},
					"Task Type": {
						"id": "vd@l",
						"multi_select": {
							"options": []
						},
						"type": "multi_select"
					}
				},
				"title": [
					{
						"annotations": {
							"bold": false,
							"code": false,
							"color": "default",
							"italic": false,
							"strikethrough": false,
							"underline": false
						},
						"href": null,
						"plain_text": "Tasks",
						"text": {
							"content": "Tasks",
							"link": null
						},
						"type": "text"
					}
				]
			},
			{
				"archived": false,
				"created_time": "2021-04-23T04:21:00.000Z",
				"id": "4f555b50-3a9b-49cb-924c-3746f4ca5522",
				"last_edited_time": "2021-04-23T04:21:00.000Z",
				"object": "page",
				"parent": {
					"database_id": "e6c6f8ff-c70e-4970-91ba-98f03e0d7fc6",
					"type": "database_id"
				},
				"properties": {
					"Name": {
						"id": "title",
						"title": [
							{
								"annotations": {
									"bold": false,
									"code": false,
									"color": "default",
									"italic": false,
									"strikethrough": false,
									"underline": false
								},
								"href": null,
								"plain_text": "Task 1",
								"text": {
									"content": "Task1 1",
									"link": null
								},
								"type": "text"
							}
						],
						"type": "title"
					}
				}
			}
		]
	}
}`
}

func TestSearchService_Search(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	tcs := map[string]struct {
		input *SearchRequest
		want  *SearchResult
	}{
		"ok": {
			&SearchRequest{
				Query: "External tasks",
				Sort: &Sort{
					Direction: Ascending,
					Timestamp: "last_edited_time",
				},
			},
			&SearchResult{
				HasMore:    false,
				NextCursor: "",
				Object:     "list",
			},
		},
	}

	for n, tc := range tcs {
		t.Run(n, func(t *testing.T) {
			mux.HandleFunc(fmt.Sprintf("/%s", searchPath), func(w http.ResponseWriter, r *http.Request) {
				if r.Header.Get(notionVersionHeader) == "" {
					t.Fatalf("no notion version header to request")
				}

				fmt.Fprint(w, getSearchJSON())
			})

			got, err := client.Search.Search(context.Background(), tc.input)
			if err != nil {
				t.Fatalf("Failed: %v", err)
			}

			if diff := cmp.Diff(got, tc.want, cmpopts.IgnoreFields(*got, "Results")); diff != "" {
				t.Fatalf("Diff: %s(-got +want)", diff)
			}
		})
	}
}
