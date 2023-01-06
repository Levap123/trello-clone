package service

import (
	"github.com/Levap123/trello-clone/internal/entity"
	"github.com/Levap123/trello-clone/internal/repository"
)

type CardService struct {
	repo repository.Card
}

func NewCardService(repo repository.Card) *CardService {
	return &CardService{
		repo: repo,
	}
}

func (cs *CardService) Create(title string, userId, workspaceId, boardId, ListId int) (int, error) {
	return cs.repo.Create(title, userId, workspaceId, boardId, ListId)
}

func (cs *CardService) GetByListId(userId, workspaceId, boardId, ListId int) ([]entity.Cards, error) {
	return cs.repo.GetByListId(userId, workspaceId, boardId, ListId)
}
