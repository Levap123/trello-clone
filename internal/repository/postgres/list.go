package postgres

import (
	"fmt"

	errs "github.com/Levap123/trello-clone/pkg/errors"
	"github.com/jmoiron/sqlx"
)

type ListRepo struct {
	db *sqlx.DB
}

func NewListRepo(db *sqlx.DB) *ListRepo {
	return &ListRepo{
		db: db,
	}
}

func (lr *ListRepo) Create(title string, userId, workspaceId, boardId int) (int, error) {
	tx, err := withTx(lr.db)
	if err != nil {
		return 0, err
	}
	defer tx.Rollback()
	ok, err := IsWorkspaceAssignedToUser(lr.db, userId, workspaceId)
	if err != nil {
		return 0, err
	}
	if !ok {
		return 0, errs.ErrInvalidWorkspace
	}
	ok, err = IsBoardAssignedToWorkspace(lr.db, workspaceId, boardId)
	if err != nil {
		return 0, err
	}
	if !ok {
		return 0, errs.ErrInvalidWorkspace
	}
	var listId int
	query := fmt.Sprintf("INSERT INTO %s (title, board_id) VALUES ($1, $2) RETURNING id", listTable)
	if err := tx.Get(&listId, query, title, boardId); err != nil {
		return 0, err
	}
	return listId, tx.Commit()
}
