package notion

// RichTextType is type of this rich text object
type RichTextType string

const (
	Text     RichTextType = "text"
	Mention  RichTextType = "mention"
	Equation RichTextType = "equation"
)

// RichText object represents Notion rich text object
//
// API doc: https://developers.notion.com/reference/rich-text
//go:generate gomodifytags -file $GOFILE -struct RichText -clear-tags -w
//go:generate gomodifytags --file $GOFILE --struct RichText -add-tags json,mapstructure -w -transform snakecase
type RichText struct {
	PlainText   string       `json:"plain_text" mapstructure:"plain_text"`
	Href        string       `json:"href" mapstructure:"href"`
	Annotations *Annotations `json:"annotations" mapstructure:"annotations"`
	Type        RichTextType `json:"type" mapstructure:"type"`
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
	Content string      `json:"content" mapstructure:"content"`
	Link    *LinkObject `json:"link" mapstructure:"link"`
}

// LinkObject object represents Notion rich text object
//go:generate gomodifytags -file $GOFILE -struct LinkObject -clear-tags -w
//go:generate gomodifytags --file $GOFILE --struct LinkObject -add-tags json,mapstructure -w -transform snakecase
type LinkObject struct {
	Type string `json:"type" mapstructure:"type"`
	URL  string `json:"url" mapstructure:"url"`
}

// MentionObjectType is for types of mentions.
type MentionObjectType string

const (
	UserMentionObject     MentionObjectType = "user"
	PageMentionObject     MentionObjectType = "page"
	DatabaseMentionObject MentionObjectType = "database"
	DateionObject         MentionObjectType = "date"
)

// TextObject object represents Notion rich text object
//go:generate gomodifytags -file $GOFILE -struct TextObject -clear-tags -w
//go:generate gomodifytags --file $GOFILE --struct TextObject -add-tags json,mapstructure -w -transform snakecase
type MentionObject struct {
	Type     MentionObjectType
	Database *Database
	User     *User
}

// EquationObject object represents Notion rich text object
//go:generate gomodifytags -file $GOFILE -struct EquationObject -clear-tags -w
//go:generate gomodifytags --file $GOFILE --struct EquationObject -add-tags json,mapstructure -w -transform snakecase
type EquationObject struct {
	Expression string `json:"expression" mapstructure:"expression"`
}
