package object

// ParentType is a type for Parents.
type ParentType string

const (
	DatabaseParentType  ParentType = "database_id"
	PageParentType      ParentType = "page_id"
	WorkspaceParentType ParentType = "workspace"
)
