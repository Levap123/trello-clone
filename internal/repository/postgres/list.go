package postgres

import (
	"fmt"

	"github.com/Levap123/trello-clone/internal/entity"
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
	if err := checkAllConstraints(lr.db, userId, workspaceId, boardId); err != nil {
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
	if err := checkAllConstraints(lr.db, userId, workspaceId, boardId); err != nil {
		return nil, err
	}
	var lists []entity.List
	query := fmt.Sprintf("SELECT * FROM %s WHERE board_id = $1", listTable)
	if err := tx.Select(&lists, query, boardId); err != nil {
		return nil, err
	}
	return lists, nil
}

func (lr *ListRepo) GetById(userId, workspaceId, boardId, listId int) (entity.List, error) {
	tx, err := withTx(lr.db)
	if err != nil {
		return entity.List{}, err
	}
	defer tx.Rollback()
	if err := checkAllConstraints(lr.db, userId, workspaceId, boardId); err != nil {
		return entity.List{}, err
	}
	var list entity.List
	query := fmt.Sprintf("SELECT * FROM %s WHERE  id = $1", listTable)
	if err := tx.Get(&list, query, listId); err != nil {
		return entity.List{}, err
	}
	return list, nil
}
func (lr *ListRepo) DeleteById(userId, workspaceId, boardId, listId int) (int, error) {
	tx, err := withTx(lr.db)
	if err != nil {
		return 0, err
	}
	defer tx.Rollback()
	if err := checkAllConstraints(lr.db, userId, workspaceId, boardId); err != nil {
		return 0, err
	}

	query := fmt.Sprintf("DELETE FROM %s WHERE id = $1 RETURNING id", listTable)
	if err := tx.Get(&listId, query, listId); err != nil {
		return 0, err
	}
	return listId, tx.Commit()
}
