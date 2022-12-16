package configs

import "os"

type DbConfigs struct {
	Name     string
	Password string
	Sslmode  string
	Host     string
	Port     string
	DbName   string
}

func NewDbConfigs() *DbConfigs {
	return &DbConfigs{
		Name:     os.Getenv("USER_NAME"),
		DbName: os.Getenv("DB_NAME"),
		Password: os.Getenv("DB_PASSWORD"),
		Sslmode:  "disable",
		Host:     os.Getenv("HOST"),
		Port:     os.Getenv("PORT"),
	}
}
