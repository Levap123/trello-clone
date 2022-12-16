package repository

import (
	"fmt"

	"github.com/Levap123/trello-clone/configs"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
)

func InitDb() (*sqlx.DB, error) {
	dbCfg := configs.NewDbConfigs()
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		dbCfg.Host, dbCfg.Port, dbCfg.Name, dbCfg.Password, dbCfg.DbName)

	db, err := sqlx.Open("pgx", psqlInfo)
	if err != nil {
		return nil, err
	}
	return db, nil
}
