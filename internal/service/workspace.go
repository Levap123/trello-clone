package service

import (
	"github.com/Levap123/trello-clone/internal/entity"
	"github.com/Levap123/trello-clone/internal/repository"
)

type WorkspaceService struct {
	repo repository.Workspace
}

func NewWorkspaceService(repo repository.Workspace) *WorkspaceService {
	return &WorkspaceService{
		repo: repo,
	}
}

func (ws *WorkspaceService) Create(title, logo string, userId int) (int, error) {
	return ws.repo.Create(title, logo, userId)
}

func (ws *WorkspaceService) GetAll(userId int) ([]entity.Workspace, error) {
	return ws.repo.GetAll(userId)
}

func (ws *WorkspaceService) GetById(userId, workspaceId int) (entity.Workspace, error) {
	return ws.repo.GetById(userId, workspaceId)
}

func (ws *WorkspaceService) DeleteById(userId, workspaceId int) (int, error) {
	return ws.repo.DeleteById(userId, workspaceId)
}
