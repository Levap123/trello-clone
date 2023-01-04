package repository

import (
	"github.com/Levap123/trello-clone/internal/entity"
	"github.com/Levap123/trello-clone/internal/repository/postgres"
	"github.com/jmoiron/sqlx"
)

type Repository struct {
	Auth
	Workspace
	Board
	List
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

type Board interface {
	Create(title, background string, userId, workspaceId int) (int, error)
	GetByWorkspaceId(userId, worskspaceId int) ([]entity.Board, error)
	GetById(userId, boardId, workspaceId int) (entity.Board, error)
	DeleteById(userId, workspace, boardId int) (int, error)
}

type List interface {
	Create(title string, userId, workspaceId, boardId int) (int, error)
	GetByBoardId(userId, workspaceId, boardId int) ([]entity.List, error)
	GetById(userId, workspaceId, boardId, listId int) (entity.List, error)
	DeleteById(userId, workspaceId, boardId, listId int) (int, error)
}

func NewRepoPostgres(db *sqlx.DB) *Repository {
	return &Repository{
		Auth:      postgres.NewAuthRepo(db),
		Workspace: postgres.NewWorkspaceRepo(db),
		Board:     postgres.NewBoardRepo(db),
		List:      postgres.NewListRepo(db),
	}
}
