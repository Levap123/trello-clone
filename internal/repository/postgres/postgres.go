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

func IsWorkspaceAssignedToUser(db *sqlx.DB, userId, workspaceId int) (bool, error) {
	tx, err := withTx(db)
	if err != nil {
		return false, err
	}
	defer tx.Rollback()
	var counter int
	query := fmt.Sprintf("SELECT COUNT(*) FROM %s WHERE user_id = $1 AND workspace_id = $2", workspaceRelationTable)
	if err := tx.Get(&counter, query, userId, workspaceId); err != nil {
		return false, err
	}
	return counter > 0, tx.Commit()
}

func IsBoardAssignedToWorkspace(db *sqlx.DB, workspaceId, boardId int) (bool, error) {
	tx, err := withTx(db)
	if err != nil {
		return false, err
	}
	defer tx.Rollback()
	var counter int
	query := fmt.Sprintf("SELECT COUNT(*) FROM %s WHERE workspace_id = $1 AND id = $2", boardTable)
	if err := tx.Get(&counter, query, workspaceId, boardId); err != nil {
		return false, err
	}
	return counter > 0, tx.Commit()
}
