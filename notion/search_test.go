package notion

import (
	"context"
	"fmt"
	"net/http"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/ketion-so/go-notion/notion/object"
)

func getSearchDatabaseJSON() string {
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
			}
		]
	}
}`
}

func getSearchPageJSON() string {
	return `{
		"has_more": false,
		"next_cursor": null,
		"object": "list",
		"results": [
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
	tcs := map[string]struct {
		input *SearchRequest
		want  *SearchResults
		json  func() string
	}{
		"database": {
			&SearchRequest{
				Query: "External tasks",
				Sort: &Sort{
					Direction: Ascending,
					Timestamp: "last_edited_time",
				},
			},
			&SearchResults{
				HasMore:    false,
				NextCursor: "",
				Object:     "list",
				Results: []object.Object{
					&Database{
						Object:         object.Database,
						ID:             "e6c6f8ff-c70e-4970-91ba-98f03e0d7fc6",
						CreatedTime:    "2021-04-22T22:23:26.080Z",
						LastEditedTime: "2021-04-23T04:21:00.000Z",
						Title: []RichText{
							{
								PlainText:   "Tasks",
								Annotations: &Annotations{Color: "default"},
								Type:        "text",
							},
						},

						Properties: map[string]Property{
							"Name":      &DatabaseTitleProperty{Type: "title", ID: "title", Title: &RichText{}},
							"Task Type": &MultiSelectProperty{Type: "multi_select", ID: "vd@l"},
						},
					},
				},
			},
			getSearchDatabaseJSON,
		},
		"page": {
			&SearchRequest{
				Query: "External tasks",
				Sort: &Sort{
					Direction: Ascending,
					Timestamp: "last_edited_time",
				},
			},
			&SearchResults{
				HasMore:    false,
				NextCursor: "",
				Object:     "list",
				Results: []object.Object{
					&Page{
						Object: object.Page,
						ID:     "4f555b50-3a9b-49cb-924c-3746f4ca5522",
						Parent: &DatabaseParent{
							Type:       object.DatabaseParentType,
							DatabaseID: "e6c6f8ff-c70e-4970-91ba-98f03e0d7fc6",
						},
						Properties: map[string]Property{
							"Name": &PageTitleProperty{
								ID: "title",
								Title: []RichText{
									{
										PlainText:   "Task 1",
										Annotations: &Annotations{Color: "default"},
										Type:        "text",
									},
								},
								Type: "title",
							},
						},
					},
				},
			},
			getSearchPageJSON,
		},
	}

	for n, tc := range tcs {
		t.Run(n, func(t *testing.T) {
			client, mux, _, teardown := setup()
			defer teardown()
			mux.HandleFunc(fmt.Sprintf("/%s", searchPath), func(w http.ResponseWriter, r *http.Request) {
				if r.Header.Get(notionVersionHeader) == "" {
					t.Fatalf("no notion version header to request")
				}

				fmt.Fprint(w, tc.json())
			})

			got, err := client.Search.Search(context.Background(), tc.input)
			if err != nil {
				t.Fatalf("Failed: %v", err)
			}

			if diff := cmp.Diff(got, tc.want); diff != "" {
				t.Fatalf("Diff: %s(-got +want)", diff)
			}
		})
	}
}
