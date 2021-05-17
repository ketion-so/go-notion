package notion

import (
	"context"
	"fmt"
	"net/http"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/ketion-so/go-notion/notion/object"
)

func getPageJSON() string {
	return `{
		"object": "page",
		"id": "b55c9c91-384d-452b-81db-d1ef79372b75",
		"created_time": "2020-03-17T19:10:04.968Z",
		"last_edited_time": "2020-03-17T21:49:37.913Z",
		"parent": {
			"type": "workspace",
			"workspace": true
		},
		"properties": {
			"Tags": {
				"id": "G~UH",
				"type": "multi_select",
				"multi_select": [
				  {
					"id": "4b9d25a5-a0ee-49ff-a7f1-9485f4773abf",
					"name": "tags2",
					"color": "green"
				  }
				]
			  },
			  "Text": {
				"id": "Me;J",
				"type": "text",
				"text": [
				  {
					"type": "text",
					"text": {
					  "content": "uuu",
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
					"plain_text": "uuu",
					"href": null
				  }
				]
			  },
			  "Date": {
				"id": "YnD",
				"type": "date",
				"date": {
				  "start": "2021-05-12",
				  "end": null
				}
			  },
			  "Name": {
				"id": "title",
				"type": "title",
				"title": [
				  {
					"type": "text",
					"text": {
					  "content": "Hoho",
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
					"plain_text": "Hoho",
					"href": null
				  }
				]
			  }
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
			"b55c9c91-384d-452b-81db-d1ef79372b75",
			&Page{
				Object:         "page",
				ID:             "b55c9c91-384d-452b-81db-d1ef79372b75",
				CreatedTime:    "2020-03-17T19:10:04.968Z",
				LastEditedTime: "2020-03-17T21:49:37.913Z",
				Parent: &WorkspaceParent{
					Type:      object.WorkspaceParentType,
					Workspace: true,
				},
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

func createPageJSON() string {
	return `{
  "object": "page",
  "id": "251d2b5f-268c-4de2-afe9-c71ff92ca95c",
  "created_time": "2020-03-17T19:10:04.968Z",
  "last_edited_time": "2020-03-17T21:49:37.913Z",
  "parent": {
    "type": "database_id",
    "database_id": "48f8fee9-cd79-4180-bc2f-ec0398253067"
  },
  "archived": false,
  "properties": {
    "Recipes": {
      "id": "AiL",
      "type": "relation",
      "relation": []
    },
    "Cost of next trip": {
      "id": "R}wl",
      "type": "formula",
      "formula": {
        "type": "number",
        "number": null
      }
    },
    "Photos": {
      "id": "d:Cb",
      "type": "files",
      "files": []
    },
    "Store availability": {
      "id": "jrFQ",
      "type": "multi_select",
      "multi_select": []
    },
    "+1": {
      "id": "k?CE",
      "type": "people",
      "person": []
    },
    "In stock": {
      "id": "{>U;",
      "type": "checkbox",
      "checkbox": false
    },
    "Name": {
      "id": "title",
      "type": "title",
      "title": [
        {
          "type": "text",
          "text": {
            "content": "Tuscan Kale",
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
          "plain_text": "Tuscan Kale",
          "href": null
        }
      ]
    }
  }
}`
}

func TestPagesService_Create(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	tcs := map[string]struct {
		input *CreatePageRequest
		want  *Page
	}{
		"ok": {
			&CreatePageRequest{
				Parent: &DatabaseParent{},
			},
			&Page{
				Object:         "page",
				ID:             "251d2b5f-268c-4de2-afe9-c71ff92ca95c",
				CreatedTime:    "2020-03-17T19:10:04.968Z",
				LastEditedTime: "2020-03-17T21:49:37.913Z",
				Parent: &DatabaseParent{
					Type:       object.DatabaseParentType,
					DatabaseID: "48f8fee9-cd79-4180-bc2f-ec0398253067",
				},
			},
		},
	}

	for n, tc := range tcs {
		t.Run(n, func(t *testing.T) {
			mux.HandleFunc(fmt.Sprintf("/%s", pagesPath), func(w http.ResponseWriter, r *http.Request) {
				if r.Header.Get(notionVersionHeader) == "" {
					t.Fatalf("no notion version header to request")
				}

				fmt.Fprint(w, createPageJSON())
			})

			got, err := client.Pages.Create(context.Background(), tc.input)
			if err != nil {
				t.Fatalf("Failed: %v", err)
			}

			if diff := cmp.Diff(got, tc.want, cmpopts.IgnoreFields(*got, "Properties")); diff != "" {
				t.Fatalf("Diff: %s(-got +want)", diff)
			}
		})
	}
}

func updatePageJSON() string {
	return `{
		"object": "page",
		"id": "60bdc8bd-3880-44b8-a9cd-8a145b3ffbd7",
		  "created_time": "2020-03-17T19:10:04.968Z",
		  "last_edited_time": "2020-03-17T21:49:37.913Z",
		"parent": {
		  "type": "database_id",
		  "database_id": "48f8fee9-cd79-4180-bc2f-ec0398253067"
		},
		"archived": false,
		"properties": {
		  "In stock": {
			"id": "{>U;",
			"type": "checkbox",
			"checkbox": true
		  },
		  "Name": {
			"id": "title",
			"type": "title",
			"title": [
			  {
				"type": "text",
				"text": {
				  "content": "Avocado",
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
				"plain_text": "Avocado",
				"href": null
			  }
			]
		  }
		}
	  }`
}

func TestPagesService_UpdateProperties(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	tcs := map[string]struct {
		id    string
		input interface{}
		want  *Page
	}{
		"ok": {
			"60bdc8bd-3880-44b8-a9cd-8a145b3ffbd7",
			nil,
			&Page{
				Object:         "page",
				ID:             "60bdc8bd-3880-44b8-a9cd-8a145b3ffbd7",
				CreatedTime:    "2020-03-17T19:10:04.968Z",
				LastEditedTime: "2020-03-17T21:49:37.913Z",
				Parent: &DatabaseParent{
					Type:       object.DatabaseParentType,
					DatabaseID: "48f8fee9-cd79-4180-bc2f-ec0398253067",
				},
				Properties: map[string]Property{
					"In stock": &CheckboxProperty{Type: "checkbox", ID: "{>U;", Checkbox: true},
					"Name": &PageTitleProperty{Type: "title", ID: "title", Title: []TextObject{
						{
							PlainText:   "Avocado",
							Annotations: &Annotations{Color: "default"},
							Type:        "text",
							Text:        &Text{Content: "Avocado"},
						},
					}},
				},
			},
		},
	}

	for n, tc := range tcs {
		t.Run(n, func(t *testing.T) {
			mux.HandleFunc(fmt.Sprintf("/%s/%s", pagesPath, tc.id), func(w http.ResponseWriter, r *http.Request) {
				if r.Header.Get(notionVersionHeader) == "" {
					t.Fatalf("no notion version header to request")
				}

				fmt.Fprint(w, updatePageJSON())
			})

			got, err := client.Pages.UpdateProperties(context.Background(), tc.id, tc.input)
			if err != nil {
				t.Fatalf("Failed: %v", err)
			}

			if diff := cmp.Diff(got, tc.want); diff != "" {
				t.Fatalf("Diff: %s(-got +want)", diff)
			}
		})
	}
}
