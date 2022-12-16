package repository

import "github.com/jmoiron/sqlx"

type Repository struct {
	Auth
}

type Auth interface {
	CreateUser(email, username, password string) (int, error)
}

func NewRepo(db *sqlx.DB) *Repository {
	return &Repository{
		Auth: NewAuthRepo(db),
	}
}
func withTx(db *sqlx.DB) (*sqlx.Tx, error) {
	return db.Beginx()
}
