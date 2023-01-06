package postgres

import (
	"fmt"

	"github.com/Levap123/trello-clone/internal/entity"
	"github.com/jmoiron/sqlx"
)

type CardRepo struct {
	db *sqlx.DB
}

func NewCardRepo(db *sqlx.DB) *CardRepo {
	return &CardRepo{
		db: db,
	}
}

func (cr *CardRepo) Create(title string, userId, workspaceId, boardId, listId int) (int, error) {
	tx, err := withTx(cr.db)
	if err != nil {
		return 0, err
	}
	defer tx.Rollback()
	if err := checkAllConstraints(cr.db, userId, workspaceId, boardId); err != nil {
		return 0, err
	}
	var cardId int
	query := fmt.Sprintf("INSERT INTO %s (title, list_id) VALUES ($1, $2) RETURNING id", cardTable)
	if err := tx.Get(&cardId, query, title, listId); err != nil {
		return 0, err
	}
	return cardId, tx.Commit()
}

func (cr *CardRepo) GetByListId(userId, workspaceId, boardId, listId int) ([]entity.Card, error) {
	tx, err := withTx(cr.db)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()
	if err := checkAllConstraints(cr.db, userId, workspaceId, boardId); err != nil {
		return nil, err
	}
	var cards []entity.Card
	query := fmt.Sprintf("SELECT * FROM %s WHERE list_id = $1", cardTable)
	if err := tx.Select(&cards, query, listId); err != nil {
		return nil, err
	}
	return cards, nil
}
