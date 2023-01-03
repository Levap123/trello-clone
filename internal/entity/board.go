package entity

type Board struct {
	Id          int    `db:"id" json:"id,omitempty"`
	Title       string `db:"title" json:"title,omitempty"`
	Background  string `db:"background" json:"background,omitempty"`
	WorkspaceId int    `db:"workspace_id" json:"workspace_id,omitempty"`
}
