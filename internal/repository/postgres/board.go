package postgres

import (
	"fmt"

	"github.com/Levap123/trello-clone/internal/entity"
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

	if err := IsWorkspaceAssignedToUser(br.db, userId, workspaceId); err != nil {
		return 0, err
	}

	query := fmt.Sprintf("INSERT INTO %s (title, background, workspace_id) VALUES ($1, $2, $3) RETURNING id", boardTable)
	if err := tx.Get(&id, query, title, background, workspaceId); err != nil {
		return 0, errs.Fail(err, "Create board")
	}
	return id, tx.Commit()
}

func (br *BoardRepo) GetById(userId, boardId, workspaceId int) (entity.Board, error) {
	tx, err := withTx(br.db)
	if err != nil {
		return entity.Board{}, err
	}
	defer tx.Rollback()
	if err := checkAllConstraints(br.db, userId, workspaceId, boardId); err != nil {
		return entity.Board{}, err
	}
	var board entity.Board
	query := fmt.Sprintf("SELECT * FROM %s WHERE workspace_id = $1 AND id = $2", boardTable)
	if err := tx.Get(&board, query, workspaceId, boardId); err != nil {
		return entity.Board{}, err
	}
	return board, tx.Commit()
}

func (br *BoardRepo) GetByWorkspaceId(userId, workspaceId int) ([]entity.Board, error) {
	tx, err := withTx(br.db)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()
	if err := IsWorkspaceAssignedToUser(br.db, userId, workspaceId); err != nil {
		return nil, err
	}
	var boards []entity.Board
	query := fmt.Sprintf("SELECT * FROM %s WHERE workspace_id = $1", boardTable)
	if err := tx.Select(&boards, query, workspaceId); err != nil {
		return nil, err
	}
	return boards, tx.Commit()
}

func (br *BoardRepo) DeleteById(userId, workspaceId, boardId int) (int, error) {
	tx, err := withTx(br.db)
	if err != nil {
		return 0, err
	}
	defer tx.Rollback()

	if err := checkAllConstraints(br.db, userId, workspaceId, boardId); err != nil {
		return 0, err
	}
	query := fmt.Sprintf("DELETE FROM %s WHERE workspace_id = $1 and id = $2 RETURNING id", boardTable)
	if err := tx.Get(&boardId, query, workspaceId, boardId); err != nil {
		return 0, errs.ErrInvalidBoard
	}
	return boardId, tx.Commit()
}
