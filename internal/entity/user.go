package entity

type User struct {
	Id       int    `db:"id,omitempty"`
	Email    string `db:"email,omitempty"`
	Name     string `db:"name,omitempty"`
	Password string `db:"password,omitempty"`
}
