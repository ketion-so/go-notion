package notion

import (
	"context"
	"fmt"
	"net/http"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
)

func getDatabaseSON() string {
	return `{
		"object": "database",
		"id": "668d797c-76fa-4934-9b05-ad288df2d136",
		"created_time": "2020-03-17T19:10:04.968Z",
		"last_edited_time": "2020-03-17T21:49:37.913Z",
		"title": [
		  {
			"type": "text",
			"text": {
			  "content": "Grocery List",
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
			"plain_text": "Grocery List",
			"href": null
		  }
		],
		"properties": {
		  "Name": {
			"id": "title",
			"type": "title",
			"title": {}
		  },
		  "In stock": {
			"id": "{xY",
			"type": "checkbox",
			"checkbox": {}
		  },
		  "Food group": {
			"id": "TJmr",
			"type": "select",
			"select": {
			  "options": [
				{
				  "id": "96eb622f-4b88-4283-919d-ece2fbed3841",
				  "name": "🥦Vegetable",
				  "color": "green"
				},
				{
				  "id": "bb443819-81dc-46fb-882d-ebee6e22c432",
				  "name": "🍎Fruit",
				  "color": "red"
				},
				{
				  "id": "7da9d1b9-8685-472e-9da3-3af57bdb221e",
				  "name": "💪Protein",
				  "color": "yellow"
				}
			  ]
			}
		  },
		  "Price": {
			"id": "cU^N",
			"type": "number",
			"number": {
			  "format": "dollar"
			}
		  },
		  "Cost of next trip": {
			"id": "p:sC",
			"type": "formula",
			"formula": {
			  "value": "if(prop(\"In stock\"), 0, prop(\"Price\"))"
			}
		  },
		  "Last ordered": {
			"id": "]\\R[",
			"type": "date",
			"date": {}
		  },
		  "Number of meals": {
			"id": "Z\\Eh",
			"type": "rollup",
			"rollup": {
			  "rollup_property_name": "Name",
			  "relation_property_name": "Meals",
			  "rollup_property_id": "title",
			  "relation_property_id": "mxp^",
			  "function": "count"
			}
		  },
		  "Store availability": {
			"type": "multi_select",
			"multi_select": {
			  "options": [
				[
				  {
					"id": "d209b920-212c-4040-9d4a-bdf349dd8b2a",
					"name": "Duc Loi Market",
					"color": "blue"
				  },
				  {
					"id": "70104074-0f91-467b-9787-00d59e6e1e41",
					"name": "Rainbow Grocery",
					"color": "gray"
				  },
				  {
					"id": "e6fd4f04-894d-4fa7-8d8b-e92d08ebb604",
					"name": "Nijiya Market",
					"color": "purple"
				  },
				  {
					"id": "6c3867c5-d542-4f84-b6e9-a420c43094e7",
					"name": "Gus's Community Market",
					"color": "yellow"
				  }
				]
			  ]
			}
		  },
		  "+1": {
			"id": "aGut",
			"type": "people",
			"people": {}
		  },
		  "Photo": {
			"id": "aTIT",
			"type": "files",
			"files": {}
		  }
		}
	  }`
}

func TestDatabasesService_Get(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	tcs := map[string]struct {
		id   string
		want *Database
	}{
		"ok": {
			"668d797c-76fa-4934-9b05-ad288df2d136",
			&Database{
				ID:             "668d797c-76fa-4934-9b05-ad288df2d136",
				Object:         "database",
				CreatedTime:    "2020-03-17T19:10:04.968Z",
				LastEditedTime: "2020-03-17T21:49:37.913Z",
				Title: []TextObject{
					{
						PlainText:   "Grocery List",
						Annotations: &Annotations{Color: DefaultColor},
						Type:        Text,
					},
				},
			},
		},
	}

	for n, tc := range tcs {
		t.Run(n, func(t *testing.T) {
			mux.HandleFunc(fmt.Sprintf("/%s/%s", databasesPath, tc.id), func(w http.ResponseWriter, r *http.Request) {
				if r.Header.Get(notionVersionHeader) == "" {
					t.Fatalf("no notion version header to request")
				}

				fmt.Fprint(w, getDatabaseSON())
			})

			got, err := client.Databases.Get(context.Background(), tc.id)
			if err != nil {
				t.Fatalf("Failed: %v", err)
			}

			if diff := cmp.Diff(got, tc.want, cmpopts.IgnoreFields(*got, "Properties")); diff != "" {
				t.Fatalf("Diff: %s(-got +want)", diff)
			}
		})
	}
}

func getListDatabaseJSON() string {
	return `{
		"results": [
		  {
			"object": "database",
			"id": "668d797c-76fa-4934-9b05-ad288df2d136",
			"title": "Grocery list",
			"properties": {
			  "Name": {
				"type": "title",
				"title": {}
			  }
			}
		  },
		  {
			"object": "database",
			"id": "74ba0cb2-732c-4d2f-954a-fcaa0d93a898",
			"title": "Pantry",
			"properties": {
			  "Name": {
				"type": "title",
				"title": {}
			  }
			}
		  }
		],
		"next_cursor": "MTY3NDE4NGYtZTdiYy00NzFlLWE0NjctODcxOTIyYWU3ZmM3",
		"has_more": false
	  }`
}

func TestDatabasesService_List(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	tcs := map[string]struct {
		id   string
		want *ListDatabaseResponse
	}{
		"ok": {
			"668d797c-76fa-4934-9b05-ad288df2d136",
			&ListDatabaseResponse{
				HasMore:    false,
				NextCursor: "MTY3NDE4NGYtZTdiYy00NzFlLWE0NjctODcxOTIyYWU3ZmM3",
			},
		},
	}

	for n, tc := range tcs {
		t.Run(n, func(t *testing.T) {
			mux.HandleFunc(fmt.Sprintf("/%s", databasesPath), func(w http.ResponseWriter, r *http.Request) {
				if r.Header.Get(notionVersionHeader) == "" {
					t.Fatalf("no notion version header to request")
				}

				fmt.Fprint(w, getListDatabaseJSON())
			})

			got, err := client.Databases.List(context.Background())
			if err != nil {
				t.Fatalf("Failed: %v", err)
			}

			if diff := cmp.Diff(got, tc.want, cmpopts.IgnoreFields(*got, "Results")); diff != "" {
				t.Fatalf("Diff: %s(-got +want)", diff)
			}
		})
	}
}
