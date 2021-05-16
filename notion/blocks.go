package notion

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/ketion-so/go-notion/notion/object"
	"github.com/mitchellh/mapstructure"
)

const (
	blocksPath = "blocks"
)

// BlocksService handles communication to Notion Blocks API.
//
// API doc: https://developers.notion.com/reference/database
type BlocksService service

// ListBlockChildrenResult object represents the retrieve block children.
//go:generate gomodifytags -file $GOFILE -struct ListBlockChildrenResult -clear-tags -w
//go:generate gomodifytags --file $GOFILE --struct ListBlockChildrenResult -add-tags json,mapstructure -w -transform snakecase
type ListBlockChildrenResult struct {
	Object  object.Type `json:"object" mapstructure:"object"`
	Results []Block     `json:"results" mapstructure:"results"`
}

// Block represents a block.
type Block interface {
	GetType() object.BlockType
}

// ParagraphBlock object represents the retrieve block children.
//go:generate gomodifytags -file $GOFILE -struct ParagraphBlock -clear-tags -w
//go:generate gomodifytags --file $GOFILE --struct ParagraphBlock -add-tags json,mapstructure -w -transform snakecase
type ParagraphBlock struct {
	Object         object.Type      `json:"object" mapstructure:"object"`
	ID             string           `json:"id" mapstructure:"id"`
	Type           object.BlockType `json:"type" mapstructure:"type"`
	CreatedTime    string           `json:"created_time" mapstructure:"created_time"`
	LastEditedTime string           `json:"last_edited_time" mapstructure:"last_edited_time"`
	HasChildren    bool             `json:"has_children" mapstructure:"has_children"`
	Text           []RichTextType   `json:"text" mapstructure:"text"`
	Children       []Block          `json:"children" mapstructure:"children"`
}

func (b *ParagraphBlock) GetType() object.BlockType {
	return b.Type
}

// HeadingOneBlock object represents the retrieve block children.
//go:generate gomodifytags -file $GOFILE -struct HeadingOneBlock -clear-tags -w
//go:generate gomodifytags --file $GOFILE --struct HeadingOneBlock -add-tags json,mapstructure -w -transform snakecase
type HeadingOneBlock struct {
	Object         object.Type      `json:"object" mapstructure:"object"`
	ID             string           `json:"id" mapstructure:"id"`
	Type           object.BlockType `json:"type" mapstructure:"type"`
	CreatedTime    string           `json:"created_time" mapstructure:"created_time"`
	LastEditedTime string           `json:"last_edited_time" mapstructure:"last_edited_time"`
	HasChildren    bool             `json:"has_children" mapstructure:"has_children"`
	Text           []RichTextType   `json:"text" mapstructure:"text"`
}

func (b *HeadingOneBlock) GetType() object.BlockType {
	return b.Type
}

// HeadingTwoBlock object represents the retrieve block children.
//go:generate gomodifytags -file $GOFILE -struct HeadingTwoBlock -clear-tags -w
//go:generate gomodifytags --file $GOFILE --struct HeadingTwoBlock -add-tags json,mapstructure -w -transform snakecase
type HeadingTwoBlock struct {
	Object         object.Type      `json:"object" mapstructure:"object"`
	ID             string           `json:"id" mapstructure:"id"`
	Type           object.BlockType `json:"type" mapstructure:"type"`
	CreatedTime    string           `json:"created_time" mapstructure:"created_time"`
	LastEditedTime string           `json:"last_edited_time" mapstructure:"last_edited_time"`
	HasChildren    bool             `json:"has_children" mapstructure:"has_children"`
	Text           []RichTextType   `json:"text" mapstructure:"text"`
}

func (b *HeadingTwoBlock) GetType() object.BlockType {
	return b.Type
}

// HeadingThreeBlock object represents the retrieve block children.
//go:generate gomodifytags -file $GOFILE -struct HeadingThreeBlock -clear-tags -w
//go:generate gomodifytags --file $GOFILE --struct HeadingThreeBlock -add-tags json,mapstructure -w -transform snakecase
type HeadingThreeBlock struct {
	Object         object.Type      `json:"object" mapstructure:"object"`
	ID             string           `json:"id" mapstructure:"id"`
	Type           object.BlockType `json:"type" mapstructure:"type"`
	CreatedTime    string           `json:"created_time" mapstructure:"created_time"`
	LastEditedTime string           `json:"last_edited_time" mapstructure:"last_edited_time"`
	HasChildren    bool             `json:"has_children" mapstructure:"has_children"`
	Text           []RichTextType   `json:"text" mapstructure:"text"`
}

func (b *HeadingThreeBlock) GetType() object.BlockType {
	return b.Type
}

// BulletedListItemBlock object represents the retrieve block children.
//go:generate gomodifytags -file $GOFILE -struct BulletedListItemBlock -clear-tags -w
//go:generate gomodifytags --file $GOFILE --struct BulletedListItemBlock -add-tags json,mapstructure -w -transform snakecase
type BulletedListItemBlock struct {
	Object         object.Type      `json:"object" mapstructure:"object"`
	ID             string           `json:"id" mapstructure:"id"`
	Type           object.BlockType `json:"type" mapstructure:"type"`
	CreatedTime    string           `json:"created_time" mapstructure:"created_time"`
	LastEditedTime string           `json:"last_edited_time" mapstructure:"last_edited_time"`
	HasChildren    bool             `json:"has_children" mapstructure:"has_children"`
	Text           []RichTextType   `json:"text" mapstructure:"text"`
	Children       []Block          `json:"children" mapstructure:"children"`
}

func (b *BulletedListItemBlock) GetType() object.BlockType {
	return b.Type
}

// NumberedListItemBlock object represents the retrieve block children.
//go:generate gomodifytags -file $GOFILE -struct NumberedListItemBlock -clear-tags -w
//go:generate gomodifytags --file $GOFILE --struct NumberedListItemBlock -add-tags json,mapstructure -w -transform snakecase
type NumberedListItemBlock struct {
	Object         object.Type      `json:"object" mapstructure:"object"`
	ID             string           `json:"id" mapstructure:"id"`
	Type           object.BlockType `json:"type" mapstructure:"type"`
	CreatedTime    string           `json:"created_time" mapstructure:"created_time"`
	LastEditedTime string           `json:"last_edited_time" mapstructure:"last_edited_time"`
	HasChildren    bool             `json:"has_children" mapstructure:"has_children"`
	Text           []RichTextType   `json:"text" mapstructure:"text"`
	Children       []Block          `json:"children" mapstructure:"children"`
}

func (b *NumberedListItemBlock) GetType() object.BlockType {
	return b.Type
}

// NumberListItemBlock object represents the retrieve block children.
//go:generate gomodifytags -file $GOFILE -struct NumberListItemBlock -clear-tags -w
//go:generate gomodifytags --file $GOFILE --struct NumberListItemBlock -add-tags json,mapstructure -w -transform snakecase
type NumberListItemBlock struct {
	Object         object.Type      `json:"object" mapstructure:"object"`
	ID             string           `json:"id" mapstructure:"id"`
	Type           object.BlockType `json:"type" mapstructure:"type"`
	CreatedTime    string           `json:"created_time" mapstructure:"created_time"`
	LastEditedTime string           `json:"last_edited_time" mapstructure:"last_edited_time"`
	HasChildren    bool             `json:"has_children" mapstructure:"has_children"`
	Text           []RichTextType   `json:"text" mapstructure:"text"`
	Checked        bool             `json:"checked" mapstructure:"checked"`
	Children       []Block          `json:"children" mapstructure:"children"`
}

func (b *NumberListItemBlock) GetType() object.BlockType {
	return b.Type
}

// ToDoBlock object represents the retrieve block children.
//go:generate gomodifytags -file $GOFILE -struct ToDoBlock -clear-tags -w
//go:generate gomodifytags --file $GOFILE --struct ToDoBlock -add-tags json,mapstructure -w -transform snakecase
type ToDoBlock struct {
	Object         object.Type      `json:"object" mapstructure:"object"`
	ID             string           `json:"id" mapstructure:"id"`
	Type           object.BlockType `json:"type" mapstructure:"type"`
	CreatedTime    string           `json:"created_time" mapstructure:"created_time"`
	LastEditedTime string           `json:"last_edited_time" mapstructure:"last_edited_time"`
	HasChildren    bool             `json:"has_children" mapstructure:"has_children"`
	Text           []RichTextType   `json:"text" mapstructure:"text"`
	Checked        bool             `json:"checked" mapstructure:"checked"`
	Children       []Block          `json:"children" mapstructure:"children"`
}

func (b *ToDoBlock) GetType() object.BlockType {
	return b.Type
}

// ToggleBlock object represents the retrieve block children.
//go:generate gomodifytags -file $GOFILE -struct ToggleBlock -clear-tags -w
//go:generate gomodifytags --file $GOFILE --struct ToggleBlock -add-tags json,mapstructure -w -transform snakecase
type ToggleBlock struct {
	Object         object.Type      `json:"object" mapstructure:"object"`
	ID             string           `json:"id" mapstructure:"id"`
	Type           object.BlockType `json:"type" mapstructure:"type"`
	CreatedTime    string           `json:"created_time" mapstructure:"created_time"`
	LastEditedTime string           `json:"last_edited_time" mapstructure:"last_edited_time"`
	HasChildren    bool             `json:"has_children" mapstructure:"has_children"`
	Text           []RichTextType   `json:"text" mapstructure:"text"`
	Children       []Block          `json:"children" mapstructure:"children"`
}

func (b *ToggleBlock) GetType() object.BlockType {
	return b.Type
}

// ChildPageBlock object represents the retrieve block children.
//go:generate gomodifytags -file $GOFILE -struct ChildPageBlock -clear-tags -w
//go:generate gomodifytags --file $GOFILE --struct ChildPageBlock -add-tags json,mapstructure -w -transform snakecase
type ChildPageBlock struct {
	Object         object.Type      `json:"object" mapstructure:"object"`
	ID             string           `json:"id" mapstructure:"id"`
	Type           object.BlockType `json:"type" mapstructure:"type"`
	CreatedTime    string           `json:"created_time" mapstructure:"created_time"`
	LastEditedTime string           `json:"last_edited_time" mapstructure:"last_edited_time"`
	HasChildren    bool             `json:"has_children" mapstructure:"has_children"`
	Title          string           `json:"title" mapstructure:"title"`
}

func (b *ChildPageBlock) GetType() object.BlockType {
	return b.Type
}

// ListChildren blocks list.
//
// API doc: https://developers.notion.com/reference/get-block-children
func (s *BlocksService) ListChildren(ctx context.Context, blockID string) (*ListBlockChildrenResult, error) {
	resp, err := s.client.Get(ctx, fmt.Sprintf("%s/%s/children", blocksPath, blockID))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	data := map[string]interface{}{}
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, err
	}

	v, ok := data["results"]
	if !ok {
		return nil, errors.New("no results returned")
	}

	results := v.([]interface{})
	blocks := []Block{}
	for _, result := range results {
		blockData, ok := result.(map[string]interface{})
		if !ok {
			return nil, errors.New("not block type returns")
		}

		block, err := decodeBlock(blockData, object.BlockType(blockData["type"].(string)))
		if err != nil {
			return nil, err
		}

		blocks = append(blocks, block)
	}

	return &ListBlockChildrenResult{
		Object:  object.Type(data["object"].(string)),
		Results: blocks,
	}, nil
}

// AppendChildren children block.
//
// API doc: https://developers.notion.com/reference/get-block-children
func (s *BlocksService) AppendChildren(ctx context.Context, blockID string, children Block) (Block, error) {
	resp, err := s.client.Patch(ctx, fmt.Sprintf("%s/%s/children", databasesPath, blockID), children)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var data map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, err
	}

	blockType, ok := data["type"]
	if !ok {
		return nil, errors.New("not block type returns")
	}

	return decodeBlock(data, object.BlockType(blockType.(string)))
}

func decodeBlock(data map[string]interface{}, blockType object.BlockType) (Block, error) {
	var b Block

	switch blockType {
	case object.ParagraphBlockType:
		b = &ParagraphBlock{}
	case object.HeadingOneBlockType:
		b = &HeadingOneBlock{}
	case object.HeadingTwoBlockType:
		b = &HeadingTwoBlock{}
	case object.HeadingThreeBlockType:
		b = &HeadingThreeBlock{}
	case object.BulletedListItemBlockType:
		b = &BulletedListItemBlock{}
	case object.NumberListItemBlockType:
		b = &NumberListItemBlock{}
	case object.ToggleBlockType:
		b = &ToggleBlock{}
	case object.ToDoBlockType:
		b = &ToDoBlock{}
	case object.ChildPageBlockType:
		b = &ChildPageBlock{}
	case object.UnsupportedBlockType:
		return nil, fmt.Errorf("%s block type not supported", blockType)
	default:
		return nil, fmt.Errorf("%s block type not supported", blockType)
	}

	if err := mapstructure.Decode(data, &b); err != nil {
		return nil, err
	}

	return b, nil
}
