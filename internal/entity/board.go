package entity

type Board struct {
	Id          int    `db:"id"`
	Title       string `db:"title"`
	Background  string `db:"background"`
	WorkspaceId int    `db:"workspace_id"`
}
