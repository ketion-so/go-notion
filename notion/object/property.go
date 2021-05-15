package object

type PropertyType string

const (
	TextPropertyType           PropertyType = "text"
	TitlePropertyType          PropertyType = "title"
	NumberPropertyType         PropertyType = "number"
	SelectPropertyType         PropertyType = "select"
	MultiSelectPropertyType    PropertyType = "multi_select"
	DatePropertyType           PropertyType = "date"
	PeoplePropertyType         PropertyType = "people"
	FilesPropertyType          PropertyType = "files"
	CheckboxPropertyType       PropertyType = "checkbox"
	URLPropertyType            PropertyType = "url"
	EmailPropertyType          PropertyType = "email"
	PhoneNumberPropertyType    PropertyType = "phone_number"
	FormulaPropertyType        PropertyType = "formula"
	RelationPropertyType       PropertyType = "relation"
	RollupPropertyType         PropertyType = "rollup"
	CreatedTimePropertyType    PropertyType = "created_time"
	CreatedByPropertyType      PropertyType = "created_by"
	LastEditedTimePropertyType PropertyType = "last_edited_time"
	LastEditedByPropertyType   PropertyType = "last_edited_by"
)
