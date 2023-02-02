package mock

import (
	"github.com/Levap123/trello-clone/internal/entity"
	errs "github.com/Levap123/trello-clone/pkg/errors"
)

type AuthMockRepo struct{}

func NewAuthRepo() *AuthMockRepo {
	return &AuthMockRepo{}
}

var users = []entity.User{
	{
		Id:       1,
		Email:    "test@mail.ru",
		Password: "$2a$10$3DQNEbOrgqSl4AIrwBhqO.D.fwD.2RwP7p.reQemGDbeqbgTqEwem",
	},
	{
		Id:       2,
		Email:    "arturpidor@mail.ru",
		Password: "arturPiDoR!@#43254",
	},
	{
		Id:       2,
		Email:    "pavel@gmail.ru",
		Password: "qwqwe21DSA!@#ADiu9382",
	},
}

func (ar *AuthMockRepo) CreateUser(email, username, password string) (int, error) {
	return 1, nil
}

func (ar *AuthMockRepo) GetUser(email string) (entity.User, error) {
	for _, user := range users {
		if user.Email == email {
			return user, nil
		}
	}
	return entity.User{}, errs.ErrInvalidEmail
}
