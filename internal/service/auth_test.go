package service_test

import (
	"testing"

	"github.com/Levap123/trello-clone/internal/repository/mock"
	"github.com/Levap123/trello-clone/internal/service"
	"github.com/Levap123/trello-clone/pkg/jwt"
	pass "github.com/Levap123/trello-clone/pkg/password"
)

func TestGetUser(t *testing.T) {
	emails := []string{"test@mail.ru", "arturpidor@mail.ru", "pavel@mail.ru"}
	
	mockAuthRepo := mock.NewAuthRepo()
	testAuthSvc := service.NewAuthService(mockAuthRepo)
	t.Run("should compare password and generate valid jwt", func(t *testing.T) {
		user, err := testAuthSvc.Repo.GetUser(emails[0])
		if err != nil {
			t.Fatal("error in gettng user from mock repo")
		}
		if err := pass.ComparePassword(user.Password, "TEST123!@#"); err != nil {
			t.Fatal("error in compare password and hash")
		}
		token, err := jwt.GenerateJwt(user.Id)
		if err != nil {
			t.Fatal("error in generating jwt token")
		}
		id, err := jwt.ParseToken(token)
		if err != nil {
			t.Fatal("error in parsinj jwt token")
		}
		if id != user.Id {
			t.Fatal("token parsed incorrect")
		}
	})
}
