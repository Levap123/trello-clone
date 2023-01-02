package postgres

import (
	"fmt"

	errs "github.com/Levap123/trello-clone/pkg/errors"
	"github.com/jmoiron/sqlx"
)

type BoardRepo struct {
	db *sqlx.DB
}

func NewBoardRepo(db *sqlx.DB) *BoardRepo {
	return &BoardRepo{
		db: db,
	}
}

func (br *BoardRepo) Create(title, background string, userId, workspaceId int) (int, error) {
	tx, err := withTx(br.db)
	if err != nil {
		return 0, err
	}
	defer tx.Rollback()
	var id int
	ok, err := br.IsWorkspaceAssignToUser(userId, workspaceId)
	if err != nil {
		return 0, err
	}
	if !ok {
		return 0, errs.ErrInvalidWorkspace
	}
	query := fmt.Sprintf("INSERT INTO %s (title, background, workspace_id) VALUES ($1, $2, $3) RETURNING id", boardTable)
	if err := tx.Get(&id, query, title, background, workspaceId); err != nil {
		return 0, errs.Fail(err, "Create board")
	}
	return id, tx.Commit()
}

func (br *BoardRepo) IsWorkspaceAssignToUser(userId, workspaceId int) (bool, error) {
	tx, err := withTx(br.db)
	if err != nil {
		return false, err
	}
	defer tx.Rollback()
	var counter int
	query := fmt.Sprintf("SELECT COUNT(*) FROM %s WHERE user_id = $1 AND workspace_id = $2", workspaceRelationTable)
	if err := tx.Get(&counter, query, userId, workspaceId); err != nil {
		return false, err
	}
	return counter > 0, nil
}
