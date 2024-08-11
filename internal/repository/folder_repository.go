package repository

import (
	"context"
	"database/sql"
	"fmt"
	"notes/internal/model"
	"strings"
)

type FolderRepository interface {
	AddMultipleFolderTx(ctx context.Context, tx *sql.Tx, folders []model.Folder) error
}

type FolderRepositoryImpl struct {
	Pg *sql.DB
}

func NewFolderRepository(pg *sql.DB) FolderRepository {
	return &FolderRepositoryImpl{
		Pg: pg,
	}
}

func (r *FolderRepositoryImpl) AddMultipleFolderTx(ctx context.Context, tx *sql.Tx, folders []model.Folder) error {
	var values []any
	var placeholders []string

	SQL := "INSERT INTO public.folders (folder_id, user_id, name, created_at, updated_at) VALUES "

	for i, folder := range folders {
		placeholder := fmt.Sprintf("($%d, $%d, $%d, $%d, $%d)", i*5+1, i*5+2, i*5+3, i*5+4, i*5+5)
		placeholders = append(placeholders, placeholder)

		values = append(values, folder.FolderId, folder.UserId, folder.Name, folder.CreatedAt, folder.UpdatedAt)
	}

	SQL += strings.Join(placeholders, ", ")

	_, err := tx.ExecContext(ctx, SQL, values...)
	if err != nil {
		return err
	}
	return nil
}
