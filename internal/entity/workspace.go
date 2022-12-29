package entity

type Workspace struct {
	Id     int    `db:"id" json:"id,omitempty"`
	Title  string `db:"title" json:"title,omitempty"`
	Logo   string `db:"logo" json:"logo,omitempty"`
	UserId string `db:"user_id" json:"user_id,omitempty"`
}
