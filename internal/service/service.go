package service

type Service struct {
	Auth
}

type Auth interface {
	CreateUser(email, username, password string) (int, error)
}
