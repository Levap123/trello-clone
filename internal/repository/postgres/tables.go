package postgres

import (
	"io/ioutil"
	"os"

	errs "github.com/Levap123/trello-clone/pkg/errors"
	"github.com/jmoiron/sqlx"
)

const (
	tableSchemas           = "up.sql"
	workspacesTable        = "workspaces"
	usersTable             = "users"
	workspaceRelationTable = "users_workspaces"
	boardTable             = "boards"
	listTable              = "lists"
	cardTable              = "cards"
)

func createTables(db *sqlx.DB) error {
	f, err := os.OpenFile(tableSchemas, os.O_RDONLY, 0o755)
	if err != nil {
		return errs.Fail(err, "Create tables")
	}
	defer f.Close()
	tables, err := ioutil.ReadAll(f)
	if err != nil {
		return err
	}
	tx, _ := db.Begin()
	_, err = tx.Exec(string(tables))
	if err != nil {
		return errs.Fail(err, "Create tables")
	}
	return tx.Commit()
}
