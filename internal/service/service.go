package service

import (
	"github.com/Levap123/trello-clone/internal/repository"
)

type Service struct {
	Auth
}

type Auth interface {
	CreateUser(email, username, password string) (int, error)
	GetUser(emil, password string) (string, error)
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		Auth: NewAuthService(repo.Auth),
	}
}
