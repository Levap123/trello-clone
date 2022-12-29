package repository

import (
	"github.com/Levap123/trello-clone/internal/entity"
	"github.com/jmoiron/sqlx"
)

type Repository struct {
	Auth
	Workspace
}

type Auth interface {
	CreateUser(email, username, password string) (int, error)
	GetUser(email string) (entity.User, error)
}

type Workspace interface {
	Create(title, logo string, userId int) (int, error)
	GetAll(userId int) ([]entity.Workspace, error)
	GetById(userId, workspaceId int) (entity.Workspace, error)
	DeleteById(userId, workspaceId int) (int, error)
}

func NewRepo(db *sqlx.DB) *Repository {
	return &Repository{
		Auth:      NewAuthRepo(db),
		Workspace: NewWorkspaceRepo(db),
	}
}
func withTx(db *sqlx.DB) (*sqlx.Tx, error) {
	return db.Beginx()
}
