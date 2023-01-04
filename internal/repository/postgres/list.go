package postgres

import (
	"fmt"

	"github.com/Levap123/trello-clone/internal/entity"
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

func (lr *ListRepo) checkAllConstraints(userId, workspaceId, boardId int) (bool, error) {
	ok, err := IsWorkspaceAssignedToUser(lr.db, userId, workspaceId)
	if err != nil {
		return false, err
	}
	if !ok {
		return false, errs.ErrInvalidWorkspace
	}
	ok, err = IsBoardAssignedToWorkspace(lr.db, workspaceId, boardId)
	if err != nil {
		return false, err
	}
	if !ok {
		return false, errs.ErrInvalidWorkspace
	}
	return true, nil
}
func (lr *ListRepo) Create(title string, userId, workspaceId, boardId int) (int, error) {
	tx, err := withTx(lr.db)
	if err != nil {
		return 0, err
	}
	defer tx.Rollback()
	ok, err := lr.checkAllConstraints(userId, workspaceId, boardId)
	if err != nil {
		return 0, err
	}
	if !ok {
		return 0, err
	}
	var listId int
	query := fmt.Sprintf("INSERT INTO %s (title, board_id) VALUES ($1, $2) RETURNING id", listTable)
	if err := tx.Get(&listId, query, title, boardId); err != nil {
		return 0, err
	}
	return listId, tx.Commit()
}

func (lr *ListRepo) GetByBoardId(userId, workspaceId, boardId int) ([]entity.List, error) {
	tx, err := withTx(lr.db)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()
	ok, err := lr.checkAllConstraints(userId, workspaceId, boardId)
	if err != nil {
		return nil, err
	}
	if !ok {
		return nil, err
	}
	var lists []entity.List
	query := fmt.Sprintf("SELECT * FROM %s WHERE board_id = $1", listTable)
	if err := tx.Select(&lists, query, boardId); err != nil {
		return nil, err
	}
	return lists, nil
}
