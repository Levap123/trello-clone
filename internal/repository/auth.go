package repository

import (
	"fmt"

	errs "github.com/Levap123/trello-clone/pkg/errors"
	"github.com/jmoiron/sqlx"
)

type AuthRepo struct {
	db *sqlx.DB
}

func NewAuthRepo(db *sqlx.DB) *AuthRepo {
	return &AuthRepo{
		db: db,
	}
}

func (ar *AuthRepo) CreateUser(email, username, password string) (int, error) {
	tx, err := withTx(ar.db)
	if err != nil {
		return 0, errs.Fail(err, "Create user")
	}
	defer tx.Rollback()

	var id int
	query := fmt.Sprintf("INSERT INTO users (email, name, password) VALUES ($1, $2, $3) RETURNING id")
	if err := tx.Get(&id, query, email, username, password); err != nil {
		return 0, errs.UniqueErr
	}
	return id, tx.Commit()
}
