package entity

type User struct {
	Id       int    `db:"id,omitempty" json:"id,omitempty"`
	Email    string `db:"email,omitempty" json:"email,omitempty"`
	Name     string `db:"name,omitempty" json:"name,omitempty"`
	Password string `db:"password,omitempty" json:"password,omitempty"`
}
