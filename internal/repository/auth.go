package repository

import (
	"fmt"

	"github.com/Levap123/trello-clone/internal/entity"
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
	query := fmt.Sprintf("INSERT INTO %s (email, name, password) VALUES ($1, $2, $3) RETURNING id", usersTable)
	if err := tx.Get(&id, query, email, username, password); err != nil {
		return 0, errs.ErrUniqueUser
	}
	return id, tx.Commit()
}

func (ar *AuthRepo) GetUser(email string) (entity.User, error) {
	tx, err := withTx(ar.db)
	if err != nil {
		return entity.User{}, errs.Fail(err, "Get user")
	}
	defer tx.Rollback()

	var user entity.User
	query := fmt.Sprintf("SELECT * FROM %s WHERE email = $1", usersTable)
	if err := tx.Get(&user, query, email); err != nil {
		return entity.User{}, errs.ErrInvalidEmail
	}
	return user, nil
}
