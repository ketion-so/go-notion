package notion

import (
	"context"
	"fmt"
	"net/http"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
)

func getListChildrenJSON() string {
	return `{
		"object": "list",
		"results": [
		   {
				"object": "block",
				"id": "9bc30ad4-9373-46a5-84ab-0a7845ee52e6",
				"created_time": "2021-03-16T16:31:00.000Z",
				"last_edited_time": "2021-03-16T16:32:00.000Z",
				"has_children": false,
				"type": "heading_2",
				"heading_2": {
					"text": [
						{
							"type": "text",
							"text": {
								"content": "Lacinato kale",
								"link": null
							},
							"annotations": {
								"bold": false,
								"italic": false,
								"strikethrough": false,
								"underline": false,
								"code": false,
								"color": "default"
							},
							"plain_text": "Lacinato kale",
							"href": null
						}
					]
				}
			},
			{
				"object": "block",
				"id": "7face6fd-3ef4-4b38-b1dc-c5044988eec0",
				"created_time": "2021-03-16T16:34:00.000Z",
				"last_edited_time": "2021-03-16T16:36:00.000Z",
				"has_children": false,
				"type": "paragraph",
				"paragraph": {
					"text": [
						{
							"type": "text",
							"text": {
								"content": "Lacinato kale",
								"link": {
									"url": "https://en.wikipedia.org/wiki/Lacinato_kale"
								}
							},
							"annotations": {
								"bold": false,
								"italic": false,
								"strikethrough": false,
								"underline": false,
								"code": false,
								"color": "default"
							},
							"plain_text": "Lacinato kale",
							"href": "https://en.wikipedia.org/wiki/Lacinato_kale"
						},
						{
							"type": "text",
							"text": {
								"content": " is a variety of kale with a long tradition in Italian cuisine, especially that of Tuscany. It is also known as Tuscan kale, Italian kale, dinosaur kale, kale, flat back kale, palm tree kale, or black Tuscan palm.",
								"link": null
							},
							"annotations": {
								"bold": false,
								"italic": false,
								"strikethrough": false,
								"underline": false,
								"code": false,
								"color": "default"
							},
							"plain_text": " is a variety of kale with a long tradition in Italian cuisine, especially that of Tuscany. It is also known as Tuscan kale, Italian kale, dinosaur kale, kale, flat back kale, palm tree kale, or black Tuscan palm.",
							"href": null
						}
					]
				}
			},
			{
				"object": "block",
				"id": "7636e2c9-b6c1-4df1-aeae-3ebf0073c5cb",
				"created_time": "2021-03-16T16:35:00.000Z",
				"last_edited_time": "2021-03-16T16:36:00.000Z",
				"has_children": true,
				"type": "toggle",
				"toggle": {
					"text": [
						{
							"type": "text",
							"text": {
								"content": "Recipes",
								"link": null
							},
							"annotations": {
								"bold": true,
								"italic": false,
								"strikethrough": false,
								"underline": false,
								"code": false,
								"color": "default"
							},
							"plain_text": "Recipes",
							"href": null
						}
					]
				}
			}
		],
		"next_cursor": null,
		"has_more": false
	}`
}

func TestDatabasesService_ListChildren(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	tcs := map[string]struct {
		id   string
		want *ListBlockChildrenResult
	}{
		"ok": {
			"668d797c-76fa-4934-9b05-ad288df2d136",
			&ListBlockChildrenResult{
				Object: "list",
			},
		},
	}

	for n, tc := range tcs {
		t.Run(n, func(t *testing.T) {
			mux.HandleFunc(fmt.Sprintf("/%s/%s/children", blocksPath, tc.id), func(w http.ResponseWriter, r *http.Request) {
				if r.Header.Get(notionVersionHeader) == "" {
					t.Fatalf("no notion version header to request")
				}

				fmt.Fprint(w, getListChildrenJSON())
			})

			got, err := client.Blocks.ListChildren(context.Background(), tc.id)
			if err != nil {
				t.Fatalf("Failed: %v", err)
			}

			if diff := cmp.Diff(got, tc.want, cmpopts.IgnoreFields(*got, "Results")); diff != "" {
				t.Fatalf("Diff: %s(-got +want)", diff)
			}
		})
	}
}

func getAppendChildrenJSON() string {
	return `{
		"object": "block",
		"id": "9bd15f8d-8082-429b-82db-e6c4ea88413b",
		"created_time": "2020-03-17T19:10:04.968Z",
		"last_edited_time": "2020-03-17T21:49:37.913Z",
		"has_children": true,
		"type": "toggle",
		"toggle": {
		  "text": [
			  {
				  "type": "text",
				  "text": {
					  "content": "Recipes",
					  "link": null
				  },
				  "annotations": {
					  "bold": true,
					  "italic": false,
					  "strikethrough": false,
					  "underline": false,
					  "code": false,
					  "color": "default"
				  },
				  "plain_text": "Recipes",
				  "href": null
			  }
		  ]
		}
	  }`
}

func TestDatabasesService_AppendChildren(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	tcs := map[string]struct {
		id    string
		input Block
		want  Block
	}{
		"ok": {
			"668d797c-76fa-4934-9b05-ad288df2d136",
			nil,
			&ToggleBlock{
				Object:         "block",
				Type:           "toggle",
				ID:             "9bd15f8d-8082-429b-82db-e6c4ea88413b",
				CreatedTime:    "2020-03-17T19:10:04.968Z",
				LastEditedTime: "2020-03-17T21:49:37.913Z",
				HasChildren:    true,
			},
		},
	}

	for n, tc := range tcs {
		t.Run(n, func(t *testing.T) {
			mux.HandleFunc(fmt.Sprintf("/%s/%s/children", databasesPath, tc.id), func(w http.ResponseWriter, r *http.Request) {
				if r.Header.Get(notionVersionHeader) == "" {
					t.Fatalf("no notion version header to request")
				}

				fmt.Fprint(w, getAppendChildrenJSON())
			})

			resp, err := client.Blocks.AppendChildren(context.Background(), tc.id, tc.input)
			if err != nil {
				t.Fatalf("Failed: %v", err)
			}

			if resp.GetType() != "toggle" {
				t.Fatalf("block type not toggle")
			}

			got, ok := resp.(*ToggleBlock)
			if !ok {
				t.Fatalf("failed to cast object, got: %T, want: %T", tc.input, tc.want)
			}

			if diff := cmp.Diff(got, tc.want); diff != "" {
				t.Fatalf("Diff: %s(-got +want)", diff)
			}
		})
	}
}
