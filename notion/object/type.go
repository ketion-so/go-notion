package object

// Type is the object type.
type Type string

const (
	Bot      Type = "bot"
	Database Type = "database"
	Error    Type = "error"
	Page     Type = "page"
	Person   Type = "person"
	User     Type = "user"
	List     Type = "list"
)

type BlockType string

const (
	ParagraphBlockType        BlockType = "paragraph"
	HeadingOneBlockType       BlockType = "heading_1"
	HeadingTwoBlockType       BlockType = "heading_2"
	HeadingThreeBlockType     BlockType = "heading_3"
	BulletedListItemBlockType BlockType = "bulleted_list_item"
	NumberListItemBlockType   BlockType = "numbered_list_item"
	ToggleBlockType           BlockType = "toggle"
	ToDoBlockType             BlockType = "to_do"
	ChildPageBlockType        BlockType = "child_page"
	UnsupportedBlockType      BlockType = "unsupported"
)

type Object interface {
	GetObject() Type
}
