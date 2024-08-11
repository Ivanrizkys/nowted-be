package repository

import (
	"context"
	"database/sql"
	"fmt"
	"notes/internal/model"
	"strings"
)

type NoteRepository interface {
	AddMultipleNoteTx(ctx context.Context, tx *sql.Tx, notes []model.Notes) error
}

type NoteRepositoryImpl struct {
	Pg *sql.DB
}

func NewNoteRepository(pg *sql.DB) NoteRepository {
	return &NoteRepositoryImpl{
		Pg: pg,
	}
}

func (r *NoteRepositoryImpl) AddMultipleNoteTx(ctx context.Context, tx *sql.Tx, notes []model.Notes) error {
	var values []any
	var placeholders []string

	SQL := "INSERT INTO public.notes (note_id, user_id, folder_id, title, content, created_at, updated_at) VALUES "

	for i, note := range notes {
		placeholder := fmt.Sprintf("($%d, $%d, $%d, $%d, $%d, $%d, $%d)", i*7+1, i*7+2, i*7+3, i*7+4, i*7+5, i*7+6, i*7+7)
		placeholders = append(placeholders, placeholder)

		values = append(values, note.NoteId, note.UserId, note.FolderId, note.Title, note.Content, note.CreatedAt, note.UpdatedAt)
	}

	SQL += strings.Join(placeholders, ", ")

	_, err := tx.ExecContext(ctx, SQL, values...)
	if err != nil {
		return err
	}
	return nil
}
