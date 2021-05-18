package notion

import (
	"fmt"

	"github.com/ketion-so/go-notion/notion/object"
	"github.com/mitchellh/mapstructure"
)

// Property represents database properties.
type Property interface {
	GetType() object.PropertyType
}

// PageTitleProperty object represents Notion title Property.
//go:generate gomodifytags --file $GOFILE --struct PageTitleProperty -add-tags json,mapstructure -w -transform snakecase
type PageTitleProperty struct {
	Type  object.PropertyType `json:"type,omitempty" mapstructure:"type" `
	ID    string              `json:"id,omitempty" mapstructure:"id" `
	Title []TextObject        `json:"title,omitempty" mapstructure:"title" `
}

// GetType returns the type of the property.
func (p *PageTitleProperty) GetType() object.PropertyType {
	return object.PropertyType(p.Type)
}

// DatabaseTitleProperty object represents Notion title Property.
//go:generate gomodifytags --file $GOFILE --struct DatabaseTitleProperty -add-tags json,mapstructure -w -transform snakecase
type DatabaseTitleProperty struct {
	Type  object.PropertyType `json:"type,omitempty" mapstructure:"type" `
	ID    string              `json:"id,omitempty" mapstructure:"id" `
	Title *TextObject         `json:"title,omitempty" mapstructure:"title" `
}

// GetType returns the type of the property.
func (p *DatabaseTitleProperty) GetType() object.PropertyType {
	return object.PropertyType(p.Type)
}

// TextProperty object represents Notion rich text Property.
//go:generate gomodifytags --file $GOFILE --struct TextProperty -add-tags json,mapstructure -w -transform snakecase
type TextProperty struct {
	Type object.PropertyType `json:"type,omitempty" mapstructure:"type" `
	ID   string              `json:"id,omitempty" mapstructure:"id" `
	Text interface{}         `json:"text,omitempty" mapstructure:"text" `
}

// GetType returns the type of the property.
func (p *TextProperty) GetType() object.PropertyType {
	return object.PropertyType(p.Type)
}

// NumberProperty object represents Notion number Property.
//go:generate gomodifytags --file $GOFILE --struct NumberProperty -add-tags json,mapstructure -w -transform snakecase
type NumberProperty struct {
	Type   object.PropertyType `json:"type,omitempty" mapstructure:"type" `
	ID     string              `json:"id,omitempty" mapstructure:"id" `
	Format string              `json:"format,omitempty" mapstructure:"format" `
	Number float64             `json:"number" mapstructure:"format" `
}

// GetType returns the type of the property.
func (p *NumberProperty) GetType() object.PropertyType {
	return object.PropertyType(p.Type)
}

// SelectProperty object represents Notion select Property.
//go:generate gomodifytags --file $GOFILE --struct SelectProperty -add-tags json,mapstructure -w -transform snakecase
type SelectProperty struct {
	Type    object.PropertyType `json:"type,omitempty" mapstructure:"type" `
	ID      string              `json:"id,omitempty" mapstructure:"id" `
	Options []SelectOption      `json:"options" mapstructure:"options" `
}

// SelectOption object represents Notion select Property.
//go:generate gomodifytags --file $GOFILE --struct SelectOption -add-tags json,mapstructure -w -transform snakecase
type SelectOption struct {
	Name  string `json:"name" mapstructure:"name" `
	ID    string `json:"id,omitempty" mapstructure:"id" `
	Color Color  `json:"color,omitempty" mapstructure:"color" `
}

// GetType returns the type of the property.
func (p *SelectProperty) GetType() object.PropertyType {
	return object.PropertyType(p.Type)
}

// MultiSelectProperty object represents Notion multi select Property.
//go:generate gomodifytags --file $GOFILE --struct MultiSelectProperty -add-tags json,mapstructure -w -transform snakecase
type MultiSelectProperty struct {
	Type    object.PropertyType `json:"type,omitempty" mapstructure:"type" `
	ID      string              `json:"id,omitempty" mapstructure:"id" `
	Options []MultiSelectOption `json:"options,omitempty" mapstructure:"options" `
}

// MultiSelectOption object represents Notion select Property.
//go:generate gomodifytags --file $GOFILE --struct MultiSelectOption -add-tags json,mapstructure -w -transform snakecase
type MultiSelectOption struct {
	Name  string `json:"name" mapstructure:"name" `
	ID    string `json:"id,omitempty" mapstructure:"id" `
	Color Color  `json:"color,omitempty" mapstructure:"color" `
}

// GetType returns the type of the property.
func (p *MultiSelectProperty) GetType() object.PropertyType {
	return object.PropertyType(p.Type)
}

// DateProperty object represents Notion date Property.
//go:generate gomodifytags --file $GOFILE --struct DateProperty -add-tags json,mapstructure -w -transform snakecase
type DateProperty struct {
	Type object.PropertyType `json:"type,omitempty" mapstructure:"type" `
	ID   string              `json:"id,omitempty" mapstructure:"id" `
	Date *Date               `json:"date" mapstructure:"date"`
}

// Date represents data object's date
type Date struct {
	Start string `json:"start" mapstructure:"start"`
	End   string `json:"end,omitempty" mapstructure:"end"`
}

// GetType returns the type of the property.
func (p *DateProperty) GetType() object.PropertyType {
	return object.PropertyType(p.Type)
}

// PersonProperty object represents Notion people Property.
//go:generate gomodifytags --file $GOFILE --struct PersonProperty -add-tags json,mapstructure -w -transform snakecase
type PersonProperty struct {
	Type   object.PropertyType `json:"type,omitempty" mapstructure:"type" `
	ID     string              `json:"id,omitempty" mapstructure:"id" `
	People *User               `json:"people,omitempty" mapstructure:"people" `
}

// GetType returns the type of the property.
func (p *PersonProperty) GetType() object.PropertyType {
	return object.PropertyType(p.Type)
}

// FilesProperty object represents Notion file Property.
//go:generate gomodifytags --file $GOFILE --struct FilesProperty -add-tags json,mapstructure -w -transform snakecase
type FilesProperty struct {
	Type object.PropertyType `json:"type,omitempty" mapstructure:"type" `
	ID   string              `json:"id,omitempty" mapstructure:"id" `
	File interface{}         `json:"file,omitempty" mapstructure:"file" `
}

// GetType returns the type of the property.
func (p *FilesProperty) GetType() object.PropertyType {
	return object.PropertyType(p.Type)
}

// CheckboxProperty object represents Notion CheckboxProperty Property.
//go:generate gomodifytags --file $GOFILE --struct CheckboxProperty -add-tags json,mapstructure -w -transform snakecase
type CheckboxProperty struct {
	Type     object.PropertyType `json:"type,omitempty" mapstructure:"type" `
	ID       string              `json:"id,omitempty" mapstructure:"id" `
	Checkbox interface{}         `json:"checkbox" mapstructure:"checkbox" `
}

// GetType returns the type of the property.
func (p *CheckboxProperty) GetType() object.PropertyType {
	return object.PropertyType(p.Type)
}

// URLProperty object represents Notion text Property.
//go:generate gomodifytags --file $GOFILE --struct URLProperty -add-tags json,mapstructure -w -transform snakecase
type URLProperty struct {
	Type object.PropertyType `json:"type,omitempty" mapstructure:"type" `
	ID   string              `json:"id,omitempty" mapstructure:"id" `
	URL  interface{}         `json:"url,omitempty" mapstructure:"url" `
}

// GetType returns the type of the property.
func (p *URLProperty) GetType() object.PropertyType {
	return object.PropertyType(p.Type)
}

// EmailProperty object represents Notion emailProperty Property.
//go:generate gomodifytags --file $GOFILE --struct EmailProperty -add-tags json,mapstructure -w -transform snakecase
type EmailProperty struct {
	Type  object.PropertyType `json:"type,omitempty" mapstructure:"type" `
	ID    string              `json:"id,omitempty" mapstructure:"id" `
	Email interface{}         `json:"email" mapstructure:"email" `
}

// GetType returns the type of the property.
func (p *EmailProperty) GetType() object.PropertyType {
	return object.PropertyType(p.Type)
}

// PhoneNumberProperty object represents Notion phone number Property.
//go:generate gomodifytags --file $GOFILE --struct PhoneNumberProperty -add-tags json,mapstructure -w -transform snakecase
type PhoneNumberProperty struct {
	Type        object.PropertyType `json:"type,omitempty" mapstructure:"type" `
	ID          string              `json:"id,omitempty" mapstructure:"id" `
	PhoneNumber interface{}         `json:"phone_number,omitempty" mapstructure:"phone_number" `
}

// GetType returns the type of the property.
func (p *PhoneNumberProperty) GetType() object.PropertyType {
	return object.PropertyType(p.Type)
}

// FormulaProperty object represents Notion formula Property.
//go:generate gomodifytags --file $GOFILE --struct FormulaProperty -add-tags json,mapstructure -w -transform snakecase
type FormulaProperty struct {
	Type       object.PropertyType `json:"type,omitempty" mapstructure:"type" `
	ID         string              `json:"id,omitempty" mapstructure:"id" `
	Expression interface{}         `json:"expression,omitempty" mapstructure:"expression" `
}

// GetType returns the type of the property.
func (p *FormulaProperty) GetType() object.PropertyType {
	return object.PropertyType(p.Type)
}

// RelationProperty object represents Notion relation Property.
//go:generate gomodifytags --file $GOFILE --struct RelationProperty -add-tags json,mapstructure -w -transform snakecase
type RelationProperty struct {
	Type     object.PropertyType `json:"type,omitempty" mapstructure:"type" `
	ID       string              `json:"id,omitempty" mapstructure:"id" `
	Relation []Relation          `json:"relation,omitempty" mapstructure:"relation" `
}

// Relation object represents Notion relation.
//go:generate gomodifytags --file $GOFILE --struct Relation -add-tags json,mapstructure -w -transform snakecase
type Relation struct {
	DatabaseID         string `json:"database_id,omitempty" mapstructure:"database_id" `
	SyncedPropertyName string `json:"synced_property_name,omitempty" mapstructure:"synced_property_name" `
	SyncedPropertyID   string `json:"synced_property_id,omitempty" mapstructure:"synced_property_id" `
}

// GetType returns the type of the property.
func (p *RelationProperty) GetType() object.PropertyType {
	return object.PropertyType(p.Type)
}

// RollupProperty object represents Notion rollup Property.
//go:generate gomodifytags --file $GOFILE --struct RollupProperty -add-tags json,mapstructure -w -transform snakecase
type RollupProperty struct {
	Type   object.PropertyType `json:"type,omitempty" mapstructure:"type" `
	ID     string              `json:"id,omitempty" mapstructure:"id" `
	Rollup *Rollup             `json:"rollup,omitempty" mapstructure:"rollup" `
}

// GetType returns the type of the property.
func (p *RollupProperty) GetType() object.PropertyType {
	return object.PropertyType(p.Type)
}

// Rollup object represents Notion rollup.
//go:generate gomodifytags --file $GOFILE --struct Rollup -add-tags json,mapstructure -w -transform snakecase
type Rollup struct {
	RelationPropertyName string `json:"relation_property_name,omitempty" mapstructure:"relation_property_name" `
	RelationPropertyID   string `json:"relation_property_id,omitempty" mapstructure:"relation_property_id" `
	RollupPropertyName   string `json:"rollup_property_name,omitempty" mapstructure:"rollup_property_name" `
	RollupPropertyID     string `json:"rollup_property_id,omitempty" mapstructure:"rollup_property_id" `
	Function             string `json:"function,omitempty" mapstructure:"function" `
}

// CreatedTimeProperty object represents Notion created time Property.
//go:generate gomodifytags --file $GOFILE --struct CreatedTimeProperty -add-tags json,mapstructure -w -transform snakecase
type CreatedTimeProperty struct {
	Type        object.PropertyType `json:"type,omitempty" mapstructure:"type" `
	ID          string              `json:"id,omitempty" mapstructure:"id" `
	CreatedTime interface{}         `json:"created_time,omitempty" mapstructure:"created_time" `
}

// GetType returns the type of the property.
func (p *CreatedTimeProperty) GetType() object.PropertyType {
	return object.PropertyType(p.Type)
}

// CreatedByProperty object represents Notion created by Property.
//go:generate gomodifytags --file $GOFILE --struct CreatedByProperty -add-tags json,mapstructure -w -transform snakecase
type CreatedByProperty struct {
	Type      object.PropertyType `json:"type,omitempty" mapstructure:"type" `
	ID        string              `json:"id,omitempty" mapstructure:"id" `
	CreatedBy interface{}         `json:"created_by,omitempty" mapstructure:"created_by" `
}

// GetType returns the type of the property.
func (p *CreatedByProperty) GetType() object.PropertyType {
	return object.PropertyType(p.Type)
}

// LastEditedTimeProperty object represents Notion last edited time Property.
//go:generate gomodifytags --file $GOFILE --struct LastEditedTimeProperty -add-tags json,mapstructure -w -transform snakecase
type LastEditedTimeProperty struct {
	Type           object.PropertyType `json:"type,omitempty" mapstructure:"type" `
	ID             string              `json:"id,omitempty" mapstructure:"id" `
	LastEditedTime interface{}         `json:"last_edited_time,omitempty" mapstructure:"last_edited_time" `
}

// GetType returns the type of the property.
func (p *LastEditedTimeProperty) GetType() object.PropertyType {
	return object.PropertyType(p.Type)
}

// LastEditedByProperty object represents Notion last edited by Property.
//go:generate gomodifytags --file $GOFILE --struct LastEditedByProperty -add-tags json,mapstructure -w -transform snakecase
type LastEditedByProperty struct {
	Type         object.PropertyType `json:"type,omitempty" mapstructure:"type" `
	ID           string              `json:"id,omitempty" mapstructure:"id" `
	LastEditedBy interface{}         `json:"last_edited_by,omitempty" mapstructure:"last_edited_by" `
}

// GetType returns the type of the property.
func (p *LastEditedByProperty) GetType() object.PropertyType {
	return object.PropertyType(p.Type)
}

func convProperties(input map[string]interface{}) (map[string]Property, error) {
	properties := map[string]Property{}
	for k, v := range input {
		var p Property
		switch obj := v.(type) {
		case map[string]interface{}:
			switch object.PropertyType(obj["type"].(string)) {
			case object.TextPropertyType:
				p = &TextProperty{}
			case object.TitlePropertyType:
				switch v.(map[string]interface{})["title"].(type) {
				case map[string]interface{}:
					p = &DatabaseTitleProperty{}
				default:
					p = &PageTitleProperty{}
				}

			case object.NumberPropertyType:
				p = &NumberProperty{}
			case object.SelectPropertyType:
				p = &SelectProperty{}
			case object.MultiSelectPropertyType:
				p = &MultiSelectProperty{}
			case object.DatePropertyType:
				p = &DateProperty{}
			case object.PeoplePropertyType:
				p = &PersonProperty{}
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
				return nil, fmt.Errorf("%v type is not suppported propert type", obj["type"])
			}

			if err := mapstructure.Decode(v, &p); err != nil {
				return nil, err
			}

			properties[k] = p
		default:
		}
	}

	return properties, nil
}
