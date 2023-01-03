package service

import (
	"github.com/Levap123/trello-clone/internal/entity"
	"github.com/Levap123/trello-clone/internal/repository"
)

type Service struct {
	Auth
	Workspace
	Board
}

type Auth interface {
	CreateUser(email, username, password string) (int, error)
	GetUser(email, password string) (string, error)
}

type Workspace interface {
	Create(title, logo string, userId int) (int, error)
	GetAll(userId int) ([]entity.Workspace, error)
	GetById(userId, workspaceId int) (entity.Workspace, error)
	DeleteById(userId, workspaceId int) (int, error)
}

type Board interface {
	Create(title, background string, userId, workspaceId int) (int, error)
	GetByWorkspaceId(userId, workspaceId int) ([]entity.Board, error)
	GetById(userId, boardId, workspaceId int) (entity.Board, error)
	// DeleteById(userId, workspace, boardId int) (int, error)
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		Auth:      NewAuthService(repo.Auth),
		Workspace: NewWorkspaceService(repo.Workspace),
		Board:     NewBoardService(repo.Board),
	}
}
