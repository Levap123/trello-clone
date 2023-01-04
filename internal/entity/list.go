package entity

type List struct {
	Id      int    `db:"id" json:"id,omitempty"`
	Title   string `db:"title" json:"title,omitempty"`
	BoardId int    `db:"board_id" json:"board_id,omitempty"`
}
