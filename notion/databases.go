package notion

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/ketion-so/go-notion/notion/object"
	"github.com/mitchellh/mapstructure"
)

const (
	databasesPath = "databases"
)

// Databaseservice handles communication to Notion Databases API.
//
// API doc: https://developers.notion.com/reference/database
type DatabasesService service

// Database object represents Notion Database.
//
// API doc: https://developers.notion.com/reference/database
//go:generate gomodifytags -file $GOFILE -struct Database -clear-tags -w
//go:generate gomodifytags --file $GOFILE --struct Database -add-tags json,mapstructure -w -transform snakecase
type Database struct {
	Object         object.Type         `json:"object" mapstructure:"object"`
	ID             string              `json:"id" mapstructure:"id"`
	CreatedTime    string              `json:"created_time" mapstructure:"created_time"`
	LastEditedTime string              `json:"last_edited_time" mapstructure:"last_edited_time"`
	Title          []RichText          `json:"title" mapstructure:"title"`
	Properties     map[string]Property `json:"properties" mapstructure:"properties"`
}

//go:generate gomodifytags -file $GOFILE -struct database -clear-tags -w
//go:generate gomodifytags --file $GOFILE --struct database -add-tags json,mapstructure -w -transform snakecase
type database struct {
	Object         object.Type            `json:"object" mapstructure:"object"`
	ID             string                 `json:"id" mapstructure:"id"`
	CreatedTime    string                 `json:"created_time" mapstructure:"created_time"`
	LastEditedTime string                 `json:"last_edited_time" mapstructure:"last_edited_time"`
	Title          []RichText             `json:"title" mapstructure:"title"`
	Properties     map[string]interface{} `json:"properties" mapstructure:"properties"`
}

type Property interface {
	GetType() object.PropertyType
}

// TitleProperty object represents Notion title Property.
//go:generate gomodifytags -file $GOFILE -struct TitleProperty -clear-tags -w
//go:generate gomodifytags --file $GOFILE --struct TitleProperty -add-tags json,mapstructure -w -transform snakecase
type TitleProperty struct {
	Type  object.PropertyType `json:"type" mapstructure:"type"`
	ID    string              `json:"id" mapstructure:"id"`
	Title interface{}         `json:"title" mapstructure:"title"`
}

// GetType returns the type of the property.
func (p *TitleProperty) GetType() object.PropertyType {
	return object.PropertyType(p.Type)
}

// TextProperty object represents Notion text Property.
//go:generate gomodifytags -file $GOFILE -struct TextProperty -clear-tags -w
//go:generate gomodifytags --file $GOFILE --struct TextProperty -add-tags json,mapstructure -w -transform snakecase
type TextProperty struct {
	Type  object.PropertyType `json:"type" mapstructure:"type"`
	ID    string              `json:"id" mapstructure:"id"`
	Title *RichText           `json:"title" mapstructure:"title"`
}

// GetType returns the type of the property.
func (p *TextProperty) GetType() object.PropertyType {
	return object.PropertyType(p.Type)
}

// NumberProperty object represents Notion number Property.
//go:generate gomodifytags -file $GOFILE -struct NumberProperty -clear-tags -w
//go:generate gomodifytags --file $GOFILE --struct NumberProperty -add-tags json,mapstructure -w -transform snakecase
type NumberProperty struct {
	Type   object.PropertyType `json:"type" mapstructure:"type"`
	ID     string              `json:"id" mapstructure:"id"`
	Format string              `json:"format" mapstructure:"format"`
}

// GetType returns the type of the property.
func (p *NumberProperty) GetType() object.PropertyType {
	return object.PropertyType(p.Type)
}

// SelectProperty object represents Notion select Property.
//go:generate gomodifytags -file $GOFILE -struct SelectProperty -clear-tags -w
//go:generate gomodifytags --file $GOFILE --struct SelectProperty -add-tags json,mapstructure -w -transform snakecase
type SelectProperty struct {
	Type    object.PropertyType `json:"type" mapstructure:"type"`
	ID      string              `json:"id" mapstructure:"id"`
	Options []SelectOption      `json:"options" mapstructure:"options"`
}

// SelectOption object represents Notion select Property.
//go:generate gomodifytags -file $GOFILE -struct SelectOption -clear-tags -w
//go:generate gomodifytags --file $GOFILE --struct SelectOption -add-tags json,mapstructure -w -transform snakecase
type SelectOption struct {
	Name  string `json:"name" mapstructure:"name"`
	ID    string `json:"id" mapstructure:"id"`
	Color Color  `json:"color" mapstructure:"color"`
}

// GetType returns the type of the property.
func (p *SelectProperty) GetType() object.PropertyType {
	return object.PropertyType(p.Type)
}

// MultiSelectProperty object represents Notion multi select Property.
//go:generate gomodifytags -file $GOFILE -struct MultiSelectProperty -clear-tags -w
//go:generate gomodifytags --file $GOFILE --struct MultiSelectProperty -add-tags json,mapstructure -w -transform snakecase
type MultiSelectProperty struct {
	Type    object.PropertyType `json:"type" mapstructure:"type"`
	ID      string              `json:"id" mapstructure:"id"`
	Options []MultiSelectOption `json:"options" mapstructure:"options"`
}

// MultiSelectOption object represents Notion select Property.
//go:generate gomodifytags -file $GOFILE -struct MultiSelectOption -clear-tags -w
//go:generate gomodifytags --file $GOFILE --struct MultiSelectOption -add-tags json,mapstructure -w -transform snakecase
type MultiSelectOption struct {
	Name  string `json:"name" mapstructure:"name"`
	ID    string `json:"id" mapstructure:"id"`
	Color Color  `json:"color" mapstructure:"color"`
}

// GetType returns the type of the property.
func (p *MultiSelectProperty) GetType() object.PropertyType {
	return object.PropertyType(p.Type)
}

// DateProperty object represents Notion date Property.
//go:generate gomodifytags -file $GOFILE -struct DateProperty -clear-tags -w
//go:generate gomodifytags --file $GOFILE --struct DateProperty -add-tags json,mapstructure -w -transform snakecase
type DateProperty struct {
	Type object.PropertyType `json:"type" mapstructure:"type"`
	ID   string              `json:"id" mapstructure:"id"`
	Date interface{}         `json:"date" mapstructure:"date"`
}

// GetType returns the type of the property.
func (p *DateProperty) GetType() object.PropertyType {
	return object.PropertyType(p.Type)
}

// PeopleProperty object represents Notion people Property.
//go:generate gomodifytags -file $GOFILE -struct PeopleProperty -clear-tags -w
//go:generate gomodifytags --file $GOFILE --struct PeopleProperty -add-tags json,mapstructure -w -transform snakecase
type PeopleProperty struct {
	Type   object.PropertyType `json:"type" mapstructure:"type"`
	ID     string              `json:"id" mapstructure:"id"`
	People interface{}         `json:"people" mapstructure:"people"`
}

// GetType returns the type of the property.
func (p *PeopleProperty) GetType() object.PropertyType {
	return object.PropertyType(p.Type)
}

// FilesProperty object represents Notion file Property.
//go:generate gomodifytags -file $GOFILE -struct FilesProperty -clear-tags -w
//go:generate gomodifytags --file $GOFILE --struct FilesProperty -add-tags json,mapstructure -w -transform snakecase
type FilesProperty struct {
	Type object.PropertyType `json:"type" mapstructure:"type"`
	ID   string              `json:"id" mapstructure:"id"`
	File interface{}         `json:"file" mapstructure:"file"`
}

// GetType returns the type of the property.
func (p *FilesProperty) GetType() object.PropertyType {
	return object.PropertyType(p.Type)
}

// CheckboxProperty object represents Notion CheckboxProperty Property.
//go:generate gomodifytags -file $GOFILE -struct CheckboxProperty -clear-tags -w
//go:generate gomodifytags --file $GOFILE --struct CheckboxProperty -add-tags json,mapstructure -w -transform snakecase
type CheckboxProperty struct {
	Type     object.PropertyType `json:"type" mapstructure:"type"`
	ID       string              `json:"id" mapstructure:"id"`
	Checkbox interface{}         `json:"checkbox" mapstructure:"checkbox"`
}

// GetType returns the type of the property.
func (p *CheckboxProperty) GetType() object.PropertyType {
	return object.PropertyType(p.Type)
}

// URLProperty object represents Notion text Property.
//go:generate gomodifytags -file $GOFILE -struct URLProperty -clear-tags -w
//go:generate gomodifytags --file $GOFILE --struct URLProperty -add-tags json,mapstructure -w -transform snakecase
type URLProperty struct {
	Type object.PropertyType `json:"type" mapstructure:"type"`
	ID   string              `json:"id" mapstructure:"id"`
	URL  interface{}         `json:"url" mapstructure:"url"`
}

// GetType returns the type of the property.
func (p *URLProperty) GetType() object.PropertyType {
	return object.PropertyType(p.Type)
}

// EmailProperty object represents Notion emailProperty Property.
//go:generate gomodifytags -file $GOFILE -struct EmailProperty -clear-tags -w
//go:generate gomodifytags --file $GOFILE --struct EmailProperty -add-tags json,mapstructure -w -transform snakecase
type EmailProperty struct {
	Type  object.PropertyType `json:"type" mapstructure:"type"`
	ID    string              `json:"id" mapstructure:"id"`
	Email interface{}         `json:"email" mapstructure:"email"`
}

// GetType returns the type of the property.
func (p *EmailProperty) GetType() object.PropertyType {
	return object.PropertyType(p.Type)
}

// PhoneNumberProperty object represents Notion phone number Property.
//go:generate gomodifytags -file $GOFILE -struct PhoneNumberProperty -clear-tags -w
//go:generate gomodifytags --file $GOFILE --struct PhoneNumberProperty -add-tags json,mapstructure -w -transform snakecase
type PhoneNumberProperty struct {
	Type        object.PropertyType `json:"type" mapstructure:"type"`
	ID          string              `json:"id" mapstructure:"id"`
	PhoneNumber interface{}         `json:"phone_number" mapstructure:"phone_number"`
}

// GetType returns the type of the property.
func (p *PhoneNumberProperty) GetType() object.PropertyType {
	return object.PropertyType(p.Type)
}

// FormulaProperty object represents Notion formula Property.
//go:generate gomodifytags -file $GOFILE -struct FormulaProperty -clear-tags -w
//go:generate gomodifytags --file $GOFILE --struct FormulaProperty -add-tags json,mapstructure -w -transform snakecase
type FormulaProperty struct {
	Type       object.PropertyType `json:"type" mapstructure:"type"`
	ID         string              `json:"id" mapstructure:"id"`
	Expression string              `json:"expression" mapstructure:"expression"`
}

// GetType returns the type of the property.
func (p *FormulaProperty) GetType() object.PropertyType {
	return object.PropertyType(p.Type)
}

// RelationProperty object represents Notion relation Property.
//go:generate gomodifytags -file $GOFILE -struct RelationProperty -clear-tags -w
//go:generate gomodifytags --file $GOFILE --struct RelationProperty -add-tags json,mapstructure -w -transform snakecase
type RelationProperty struct {
	Type     object.PropertyType `json:"type" mapstructure:"type"`
	ID       string              `json:"id" mapstructure:"id"`
	Relation *Relation           `json:"relation" mapstructure:"relation"`
}

// Relation object represents Notion relation.
//go:generate gomodifytags -file $GOFILE -struct Relation -clear-tags -w
//go:generate gomodifytags --file $GOFILE --struct Relation -add-tags json,mapstructure -w -transform snakecase
type Relation struct {
	DatabaseID         string `json:"database_id" mapstructure:"database_id"`
	SyncedPropertyName string `json:"synced_property_name" mapstructure:"synced_property_name"`
	SyncedPropertyID   string `json:"synced_property_id" mapstructure:"synced_property_id"`
}

// GetType returns the type of the property.
func (p *RelationProperty) GetType() object.PropertyType {
	return object.PropertyType(p.Type)
}

// RollupProperty object represents Notion rollup Property.
//go:generate gomodifytags -file $GOFILE -struct RollupProperty -clear-tags -w
//go:generate gomodifytags --file $GOFILE --struct RollupProperty -add-tags json,mapstructure -w -transform snakecase
type RollupProperty struct {
	Type   object.PropertyType `json:"type" mapstructure:"type"`
	ID     string              `json:"id" mapstructure:"id"`
	Rollup *Rollup             `json:"rollup" mapstructure:"rollup"`
}

// GetType returns the type of the property.
func (p *RollupProperty) GetType() object.PropertyType {
	return object.PropertyType(p.Type)
}

// Rollup object represents Notion rollup.
//go:generate gomodifytags -file $GOFILE -struct Rollup -clear-tags -w
//go:generate gomodifytags --file $GOFILE --struct Rollup -add-tags json,mapstructure -w -transform snakecase
type Rollup struct {
	RelationPropertyName string `json:"relation_property_name" mapstructure:"relation_property_name"`
	RelationPropertyID   string `json:"relation_property_id" mapstructure:"relation_property_id"`
	RollupPropertyName   string `json:"rollup_property_name" mapstructure:"rollup_property_name"`
	RollupPropertyID     string `json:"rollup_property_id" mapstructure:"rollup_property_id"`
	Function             string `json:"function" mapstructure:"function"`
}

// CreatedTimeProperty object represents Notion created time Property.
//go:generate gomodifytags -file $GOFILE -struct CreatedTimeProperty -clear-tags -w
//go:generate gomodifytags --file $GOFILE --struct CreatedTimeProperty -add-tags json,mapstructure -w -transform snakecase
type CreatedTimeProperty struct {
	Type        object.PropertyType `json:"type" mapstructure:"type"`
	ID          string              `json:"id" mapstructure:"id"`
	CreatedTime interface{}         `json:"created_time" mapstructure:"created_time"`
}

// GetType returns the type of the property.
func (p *CreatedTimeProperty) GetType() object.PropertyType {
	return object.PropertyType(p.Type)
}

// CreatedByProperty object represents Notion created by Property.
//go:generate gomodifytags -file $GOFILE -struct CreatedByProperty -clear-tags -w
//go:generate gomodifytags --file $GOFILE --struct CreatedByProperty -add-tags json,mapstructure -w -transform snakecase
type CreatedByProperty struct {
	Type      object.PropertyType `json:"type" mapstructure:"type"`
	ID        string              `json:"id" mapstructure:"id"`
	CreatedBy interface{}         `json:"created_by" mapstructure:"created_by"`
}

// GetType returns the type of the property.
func (p *CreatedByProperty) GetType() object.PropertyType {
	return object.PropertyType(p.Type)
}

// LastEditedTimeProperty object represents Notion last edited time Property.
//go:generate gomodifytags -file $GOFILE -struct LastEditedTimeProperty -clear-tags -w
//go:generate gomodifytags --file $GOFILE --struct LastEditedTimeProperty -add-tags json,mapstructure -w -transform snakecase
type LastEditedTimeProperty struct {
	Type           object.PropertyType `json:"type" mapstructure:"type"`
	ID             string              `json:"id" mapstructure:"id"`
	LastEditedTime interface{}         `json:"last_edited_time" mapstructure:"last_edited_time"`
}

// GetType returns the type of the property.
func (p *LastEditedTimeProperty) GetType() object.PropertyType {
	return object.PropertyType(p.Type)
}

// LastEditedByProperty object represents Notion last edited by Property.
//go:generate gomodifytags -file $GOFILE -struct LastEditedByProperty -clear-tags -w
//go:generate gomodifytags --file $GOFILE --struct LastEditedByProperty -add-tags json,mapstructure -w -transform snakecase
type LastEditedByProperty struct {
	Type         object.PropertyType `json:"type" mapstructure:"type"`
	ID           string              `json:"id" mapstructure:"id"`
	LastEditedBy interface{}         `json:"last_edited_by" mapstructure:"last_edited_by"`
}

// GetType returns the type of the property.
func (p *LastEditedByProperty) GetType() object.PropertyType {
	return object.PropertyType(p.Type)
}

// Get retrieves database by database ID.
//
// API doc: https://developers.notion.com/reference/get-database
func (s *DatabasesService) Get(ctx context.Context, databaseID string) (*Database, error) {
	resp, err := s.client.get(ctx, fmt.Sprintf("%s/%s", databasesPath, databaseID))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	data := database{}
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, err
	}

	return convDatabase(&data)
}

// ListDatabase object represents Notion ListDatabase.
//
// API doc: https://developers.notion.com/reference/database
//go:generate gomodifytags -file $GOFILE -struct ListDatabase -clear-tags -w
//go:generate gomodifytags --file $GOFILE --struct ListDatabase -add-tags json,mapstructure -w -transform snakecase
type ListDatabase struct {
	Object         object.Type            `json:"object" mapstructure:"object"`
	ID             string                 `json:"id" mapstructure:"id"`
	CreatedTime    string                 `json:"created_time" mapstructure:"created_time"`
	LastEditedTime string                 `json:"last_edited_time" mapstructure:"last_edited_time"`
	Title          string                 `json:"title" mapstructure:"title"`
	Properties     map[string]interface{} `json:"properties" mapstructure:"properties"`
}

// ListDatabaseResponse represents the response from the list User API
//
//go:generate gomodifytags -file $GOFILE -struct ListDatabaseResponse -clear-tags -w
//go:generate gomodifytags --file $GOFILE --struct ListDatabaseResponse -add-tags json,mapstructure -w -transform snakecase
type ListDatabaseResponse struct {
	Object     object.Type    `json:"object" mapstructure:"object"`
	Results    []ListDatabase `json:"results" mapstructure:"results"`
	NextCursor string         `json:"next_cursor" mapstructure:"next_cursor"`
	HasMore    bool           `json:"has_more" mapstructure:"has_more"`
}

// List lists database.
//
// API doc: https://developers.notion.com/reference/get-databases
func (s *DatabasesService) List(ctx context.Context) (*ListDatabaseResponse, error) {
	resp, err := s.client.get(ctx, databasesPath)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	results := &ListDatabaseResponse{}
	if err := json.NewDecoder(resp.Body).Decode(results); err != nil {
		return nil, err
	}

	return results, nil
}

func convDatabase(data *database) (*Database, error) {
	properties := map[string]Property{}
	for k, v := range data.Properties {
		var p Property

		propertyType := object.PropertyType(v.(map[string]interface{})["type"].(string))
		switch object.PropertyType(v.(map[string]interface{})["type"].(string)) {
		case object.TextPropertyType:
			p = &TextProperty{}
		case object.TitlePropertyType:
			p = &TitleProperty{}
		case object.NumberPropertyType:
			p = &NumberProperty{}
		case object.SelectPropertyType:
			p = &SelectProperty{}
		case object.MultiSelectPropertyType:
			p = &MultiSelectProperty{}
		case object.DatePropertyType:
			p = &DateProperty{}
		case object.PeoplePropertyType:
			p = &PeopleProperty{}
		case object.FilesPropertyType:
			p = &FilesProperty{}
		case object.CheckboxPropertyType:
			p = &CheckboxProperty{}
		case object.URLPropertyType:
			p = &URLProperty{}
		case object.EmailPropertyType:
			p = &EmailProperty{}
		case object.PhoneNumberPropertyType:
			p = &PhoneNumberProperty{}
		case object.FormulaPropertyType:
			p = &FormulaProperty{}
		case object.RelationPropertyType:
			p = &RelationProperty{}
		case object.RollupPropertyType:
			p = &RollupProperty{}
		case object.CreatedTimePropertyType:
			p = &CreatedTimeProperty{}
		case object.CreatedByPropertyType:
			p = &CreatedByProperty{}
		case object.LastEditedTimePropertyType:
			p = &LastEditedTimeProperty{}
		case object.LastEditedByPropertyType:
			p = &LastEditedByProperty{}
		default:
			return nil, fmt.Errorf("%v type is not suppported propert type", propertyType)
		}

		if err := mapstructure.Decode(v, &p); err != nil {
			return nil, err
		}

		properties[k] = p
	}

	return &Database{
		Object:         data.Object,
		ID:             data.ID,
		Title:          data.Title,
		CreatedTime:    data.CreatedTime,
		LastEditedTime: data.LastEditedTime,
		Properties:     properties,
	}, nil
}
