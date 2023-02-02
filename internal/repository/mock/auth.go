package mock

import "github.com/Levap123/trello-clone/internal/entity"

type AuthMockRepo struct{}

func NewAuthRepo() *AuthMockRepo {
	return &AuthMockRepo{}
}

func (ar *AuthMockRepo) CreateUser(email, username, password string) (int, error) {
	return 1, nil
}

func (ar *AuthMockRepo) GetUser(email string) (entity.User, error) {
	return entity.User{
		Id:       1,
		Email:    "test@mail.ru",
		Password: "TEST123!@#",
	}, nil
}
