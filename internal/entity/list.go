package entity

type List struct {
	Id      int    `db:"id"`
	Title   string `db:"title"`
	BoardId int    `db:"board_id"`
}
