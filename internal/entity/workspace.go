package entity

type Workspace struct {
	Id     int    `db:"id"`
	Title  string `db:"title"`
	Logo   string `db:"logo"`
	UserId int    `db:"user_id"`
}
