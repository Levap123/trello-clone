package service

import (
	"github.com/Levap123/trello-clone/internal/entity"
	"github.com/Levap123/trello-clone/internal/repository"
)

type BoardService struct {
	repo repository.Board
}

func NewBoardService(repo repository.Board) *BoardService {
	return &BoardService{
		repo: repo,
	}
}

func (bs *BoardService) Create(title, background string, userId, workspaceId int) (int, error) {
	return bs.repo.Create(title, background, userId, workspaceId)
}

func (bs *BoardService) GetById(userId, boardId, workspaceId int) (entity.Board, error) {
	return bs.repo.GetById(userId, boardId, workspaceId)
}

func (bs *BoardService) GetByWorkspaceId(userId, workspaceId int) ([]entity.Board, error) {
	return bs.repo.GetByWorkspaceId(userId, workspaceId)
}
