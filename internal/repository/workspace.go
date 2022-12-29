package repository

import (
	"fmt"

	"github.com/Levap123/trello-clone/internal/entity"
	errs "github.com/Levap123/trello-clone/pkg/errors"
	"github.com/jmoiron/sqlx"
)

type WorkspaceRepo struct {
	db *sqlx.DB
}

func NewWorkspaceRepo(db *sqlx.DB) *WorkspaceRepo {
	return &WorkspaceRepo{
		db: db,
	}
}

func (wr *WorkspaceRepo) Create(title, logo string, userId int) (int, error) {
	tx, err := withTx(wr.db)
	if err != nil {
		return 0, errs.Fail(err, "Create workspace")
	}
	defer tx.Rollback()
	var workspaceId int
	query := fmt.Sprintf("INSERT INTO %s (title, logo) VALUES ($1,$2) RETURNING id", workspacesTable)
	if err := tx.Get(&workspaceId, query, title, logo); err != nil {
		return 0, errs.Fail(err, "Create workspace")
	}
	query = fmt.Sprintf("INSERT INTO %s (user_id, workspace_id) VALUES ($1,$2)", workspaceRelationTable)
	if _, err := tx.Exec(query, userId, workspaceId); err != nil {
		return 0, errs.Fail(err, "Create relation")
	}
	return workspaceId, tx.Commit()
}

func (wr *WorkspaceRepo) GetAll(userId int) ([]entity.Workspace, error) {
	tx, err := withTx(wr.db)
	if err != nil {
		return nil, errs.Fail(err, "Get all workspace")
	}
	defer tx.Rollback()
	var workspaces []entity.Workspace
	query := fmt.Sprintf("SELECT * FROM %s WHERE user_id = $1", workspacesTable)
	if err := tx.Select(&workspaces, query, userId); err != nil {
		return nil, errs.Fail(err, "Get all workspace")
	}
	return workspaces, tx.Commit()
}

func (wr *WorkspaceRepo) GetById(userId, workspaceId int) (entity.Workspace, error) {
	tx, err := withTx(wr.db)
	if err != nil {
		return entity.Workspace{}, errs.Fail(err, "Get workspace")
	}
	defer tx.Rollback()
	var workspace entity.Workspace
	query := fmt.Sprintf(`SELECT w.id, w.logo, w.title, uw.user_id FROM %s AS uw JOIN workspaces as w ON uw.workspace_id = w.id 
	WHERE uw.user_id = $1 and uw.workspace_id = $2 limit 1`, workspaceRelationTable)
	if err := tx.Get(&workspace, query, userId, workspaceId); err != nil {
		return entity.Workspace{}, errs.ErrInvalidWorkspace
	}
	return workspace, tx.Commit()
}

func (wr *WorkspaceRepo) DeleteById(userId, workspaceId int) (int, error) {
	tx, err := withTx(wr.db)
	if err != nil {
		return 0, errs.Fail(err, "Delete workspace")
	}
	defer tx.Rollback()
	var id int
	query := fmt.Sprintf("DELETE FROM %s WHERE workspace_id = $1 AND user_id = $2 RETURNING workspace_id", workspaceRelationTable)
	if err := tx.Get(&id, query, workspaceId, userId); err != nil {
		return 0, errs.ErrInvalidWorkspace
	}
	query = fmt.Sprintf("DELETE FROM %s WHERE id = $1 RETURNING id", workspacesTable)
	if _, err := tx.Exec(query, workspaceId); err != nil {
		return 0, errs.Fail(err, "Delete workspace")
	}
	return id, tx.Commit()
}
