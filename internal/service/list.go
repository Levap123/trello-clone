package service

import (
	"github.com/Levap123/trello-clone/internal/entity"
	"github.com/Levap123/trello-clone/internal/repository"
)

type ListService struct {
	repo repository.List
}

func NewListService(repo repository.List) *ListService {
	return &ListService{
		repo: repo,
	}
}

func (ls *ListService) Create(title string, userId, workspaceId, boardId int) (int, error) {
	return ls.repo.Create(title, userId, workspaceId, boardId)
}

func (ls *ListService) GetByBoardId(userId, workspaceId, boardId int) ([]entity.List, error) {
	return ls.repo.GetByBoardId(userId, workspaceId, boardId)
}

func (ls *ListService) GetById(userId, workspaceId, boardId, listId int) (entity.List, error) {
	return ls.repo.GetById(userId, workspaceId, boardId, listId)
}
