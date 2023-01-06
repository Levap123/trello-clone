package entity

type Card struct {
	Id     int    `db:"id,omitempty"`
	Title  string `db:"title,omitempty"`
	ListId string `db:"list_id,omitempty"`
}
