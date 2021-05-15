package notion

const (
	databasesPath = "databases"
)

// Databaseservice handles communication to Notion Databases API.
//
// API doc: https://developers.notion.com/reference/database
type Databaseservice service

// Database object represents Notion Database.
//
// API doc: https://developers.notion.com/reference/database
//go:generate gomodifytags -file $GOFILE -struct Database -clear-tags -w
//go:generate gomodifytags --file $GOFILE --struct Database -add-tags json -w -transform snakecase
type Database struct {
	Object         string             `json:"object"`
	ID             string             `json:"id"`
	CreatedTime    string             `json:"created_time"`
	LastEditedTime string             `json:"last_edited_time"`
	Title          []RichText         `json:"title"`
	Properties     []DatabaseProperty `json:"properties"`
}

type DatabasePropertyType string

const (
	TitleDatabase          DatabasePropertyType = "title"
	RichTextDatabase       DatabasePropertyType = "rich_text"
	NumberDatabase         DatabasePropertyType = "number"
	SelectDatabase         DatabasePropertyType = "select"
	MultiSelectDatabase    DatabasePropertyType = "multi_select"
	DateDatabase           DatabasePropertyType = "date"
	PeopleDatabase         DatabasePropertyType = "people"
	FileDatabase           DatabasePropertyType = "file"
	CheckboxDatabase       DatabasePropertyType = "checkbox"
	URLDatabase            DatabasePropertyType = "url"
	EmailDatabase          DatabasePropertyType = "email"
	PhoneNumberDatabase    DatabasePropertyType = "phone_number"
	FormulaDatabase        DatabasePropertyType = "formula"
	RelationDatabase       DatabasePropertyType = "relation"
	RollupDatabase         DatabasePropertyType = "rollup"
	CreatedTypeDatabase    DatabasePropertyType = "created_time"
	CreatedByDatabase      DatabasePropertyType = "created_by"
	LastEditedTimeDatabase DatabasePropertyType = "last_edited_time"
	LastEditedByDatabase   DatabasePropertyType = "last_edited_by"
)

type DatabaseProperty struct {
	ID   string
	Type DatabasePropertyType
}
