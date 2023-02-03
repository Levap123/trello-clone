package service_test

import (
	"testing"

	"github.com/Levap123/trello-clone/internal/repository/mock"
	"github.com/Levap123/trello-clone/internal/service"
	errs "github.com/Levap123/trello-clone/pkg/errors"
	"github.com/Levap123/trello-clone/pkg/jwt"
	pass "github.com/Levap123/trello-clone/pkg/password"
	"golang.org/x/crypto/bcrypt"
)

func TestGetUserUtils(t *testing.T) {
	emails := []string{"test@mail.ru", "arturpidor@mail.ru", "pavel@gmail.ru"}
	passwords := []string{"TEST123!@#", "arturPiDoR!@#43254", "qwqwe21DSA!@#ADiu9382"}
	mockAuthRepo := mock.NewAuthRepo()
	testAuthSvc := service.NewAuthService(mockAuthRepo)
	t.Run("should compare password and generate valid jwt", func(t *testing.T) {
		for ind, email := range emails {
			user, err := testAuthSvc.Repo.GetUser(email)
			if err != nil {
				t.Fatal("error in gettng user from mock repo")
			}
			if err := pass.ComparePassword(user.Password, passwords[ind]); err != nil {
				t.Fatal("error in compare password and hash")
			}
			token, err := jwt.GenerateJwt(user.Id)
			if err != nil {
				t.Fatal("error in generating jwt token")
			}
			id, err := jwt.ParseToken(token)
			if err != nil {
				t.Fatal("error in parsing jwt token")
			}
			if id != user.Id {
				t.Fatalf("expected id = %d, but got %d", user.Id, id)
			}
		}
	})
	t.Run("should throw err in password comparing", func(t *testing.T) {
		user, err := testAuthSvc.Repo.GetUser(emails[0])
		if err != nil {
			t.Fatal("error in gettng user from mock repo")
		}
		if err := pass.ComparePassword(user.Password, "testest"); err == nil {
			t.Fatalf("expected %v, but got nil", bcrypt.ErrMismatchedHashAndPassword)
		}
	})
}

func TestGetUser(t *testing.T) {
	emails := []string{"test@mail.ru", "arturpidor@mail.ru", "pavel@gmail.ru"}
	passwords := []string{"TEST123!@#", "arturPiDoR!@#43254", "qwqwe21DSA!@#ADiu9382"}
	mockAuthRepo := mock.NewAuthRepo()
	testAuthSvc := service.NewAuthService(mockAuthRepo)
	t.Run("should be working without any errors", func(t *testing.T) {
		for ind, email := range emails {
			_, err := testAuthSvc.GetUser(email, passwords[ind])
			if err != nil {
				t.Errorf("error should be nil, but got %v", err)
			}
		}
	})
	t.Run("should throw error in getting user", func(t *testing.T) {
		_, err := testAuthSvc.Repo.GetUser("dafijsdfouh(*datygf")
		if err == nil {
			t.Fatalf("expected %v, but got nil", errs.ErrInvalidEmail)
		}
	})
}
