package notion

// RichTextType is type of this rich text object
type RichTextType string

const (
	TextRichTextType    RichTextType = "text"
	MentionRichTextType RichTextType = "mention"
	EquationRichTextTye RichTextType = "equation"
)

// RicchText is descibed in API doc: https://developers.notion.com/reference/rich-text
type RichText interface {
	GetType() RichTextType
}

// Annotations object represents Notion rich text annotation
//go:generate gomodifytags -file $GOFILE -struct Annotations -clear-tags -w
//go:generate gomodifytags --file $GOFILE --struct Annotations -add-tags json,mapstructure -w -transform snakecase
type Annotations struct {
	Bold          bool  `json:"bold" mapstructure:"bold"`
	Italic        bool  `json:"italic" mapstructure:"italic"`
	StrikeThrough bool  `json:"strike_through" mapstructure:"strike_through"`
	Underline     bool  `json:"underline" mapstructure:"underline"`
	Code          bool  `json:"code" mapstructure:"code"`
	Color         Color `json:"color" mapstructure:"color"`
}

// Color is type for text and background colors.
type Color string

const (
	DefaultColor          Color = "default"
	GrayColor             Color = "gray"
	BrownColor            Color = "brown"
	OrangeColor           Color = "orange"
	YellowColor           Color = "yellow"
	GreenColor            Color = "green"
	BlueColor             Color = "blue"
	PurpleColor           Color = "purple"
	PinkColor             Color = "ping"
	RedColor              Color = "red"
	GrayBackGroundColor   Color = "gray_background"
	BrownBackGroundColor  Color = "brown_background"
	OrangeBackGroundColor Color = "orange_background"
	YellowBackGroundColor Color = "yellow_background"
	GreenBackGroundColor  Color = "green_background"
	BlueBackGroundColor   Color = "blue_background"
	PurpleBackGroundColor Color = "purple_background"
	PinkBackGroundColor   Color = "pink_background"
	RedBackGroundColor    Color = "red_background"
)

// TextObject object represents Notion rich text object
//go:generate gomodifytags -file $GOFILE -struct TextObject -clear-tags -w
//go:generate gomodifytags --file $GOFILE --struct TextObject -add-tags json,mapstructure -w -transform snakecase
type TextObject struct {
	PlainText   string       `json:"plain_text" mapstructure:"plain_text"`
	Href        string       `json:"href" mapstructure:"href"`
	Annotations *Annotations `json:"annotations" mapstructure:"annotations"`
	Type        RichTextType `json:"type" mapstructure:"type"`
	Text        *Text        `json:"text" mapstructure:"text"`
	Link        *LinkObject  `json:"link" mapstructure:"link"`
}

type Text struct {
	Type    RichTextType
	Content string
}

// GetType returns the object type
func (obj *TextObject) GetType() RichTextType {
	return obj.Type
}

// LinkObject object represents Notion rich text object
//go:generate gomodifytags -file $GOFILE -struct LinkObject -clear-tags -w
//go:generate gomodifytags --file $GOFILE --struct LinkObject -add-tags json,mapstructure -w -transform snakecase
type LinkObject struct {
	PlainText   string       `json:"plain_text" mapstructure:"plain_text"`
	Href        string       `json:"href" mapstructure:"href"`
	Annotations *Annotations `json:"annotations" mapstructure:"annotations"`
	Type        RichTextType `json:"type" mapstructure:"type"`
	URL         string       `json:"url" mapstructure:"url"`
}

// GetType returns the object type
func (obj *LinkObject) GetType() RichTextType {
	return obj.Type
}

// MentionObjectType is for types of mentions.
type MentionObjectType string

const (
	UserMentionObject     MentionObjectType = "user"
	PageMentionObject     MentionObjectType = "page"
	DatabaseMentionObject MentionObjectType = "database"
	DateionObject         MentionObjectType = "date"
)

// MentionObject object represents Notion rich text object
//go:generate gomodifytags -file $GOFILE -struct MentionObject -clear-tags -w
//go:generate gomodifytags --file $GOFILE --struct MentionObject -add-tags json,mapstructure -w -transform snakecase
type MentionObject struct {
	PlainText   string       `json:"plain_text" mapstructure:"plain_text"`
	Href        string       `json:"href" mapstructure:"href"`
	Annotations *Annotations `json:"annotations" mapstructure:"annotations"`
	Type        RichTextType `json:"type" mapstructure:"type"`
	Database    *Database    `json:"database" mapstructure:"database"`
	User        *User        `json:"user" mapstructure:"user"`
}

// GetType returns the object type
func (obj *MentionObject) GetType() RichTextType {
	return obj.Type
}

// EquationObject object represents Notion rich text object
//go:generate gomodifytags -file $GOFILE -struct EquationObject -clear-tags -w
//go:generate gomodifytags --file $GOFILE --struct EquationObject -add-tags json,mapstructure -w -transform snakecase
type EquationObject struct {
	PlainText   string       `json:"plain_text" mapstructure:"plain_text"`
	Href        string       `json:"href" mapstructure:"href"`
	Annotations *Annotations `json:"annotations" mapstructure:"annotations"`
	Type        RichTextType `json:"type" mapstructure:"type"`
	Expression  string       `json:"expression" mapstructure:"expression"`
}

// GetType returns the object type
func (obj *EquationObject) GetType() RichTextType {
	return obj.Type
}
