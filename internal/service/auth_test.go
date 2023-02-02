package service_test

import (
	"testing"

	"github.com/Levap123/trello-clone/internal/repository/mock"
	"github.com/Levap123/trello-clone/internal/service"
)

func testAuthService(authMockRepo *mock.AuthMockRepo) *service.AuthService {
	return service.NewAuthService(authMockRepo)
}
func testGetUser(t *testing.T) {
	
}
