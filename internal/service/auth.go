package service

import (
	"github.com/Levap123/trello-clone/internal/repository"
	errs "github.com/Levap123/trello-clone/pkg/errors"
	pass "github.com/Levap123/trello-clone/pkg/password"
)

type AuthService struct {
	repo repository.Auth
}

func NewAuthService(repo repository.Auth) *AuthService {
	return &AuthService{
		repo: repo,
	}
}

func (as *AuthService) CreateUser(email, username, password string) (int, error) {
	password, err := pass.GeneratePasswordHash(password)
	if err != nil {
		return 0, errs.Fail(err, "CreateUser")
	}
	return as.repo.CreateUser(email, username, password)
}
