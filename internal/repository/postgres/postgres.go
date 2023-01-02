package postgres

import (
	"fmt"

	"github.com/Levap123/trello-clone/configs"
	errs "github.com/Levap123/trello-clone/pkg/errors"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
)

func InitDb(dbCfg *configs.DbConfigs) (*sqlx.DB, error) {
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		dbCfg.Host, dbCfg.Port, dbCfg.Name, dbCfg.Password, dbCfg.DbName)

	db, err := sqlx.Open("pgx", psqlInfo)
	if err := createTables(db); err != nil {
		return nil, errs.Fail(err, "Init db")
	}
	if err != nil {
		return nil, errs.Fail(err, "Init db")
	}
	return db, nil
}

func withTx(db *sqlx.DB) (*sqlx.Tx, error) {
	return db.Beginx()
}
