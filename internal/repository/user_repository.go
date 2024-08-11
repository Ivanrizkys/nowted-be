package repository

import (
	"context"
	"database/sql"
	"notes/internal/model"
)

type UserRepository interface {
	AddUserTx(ctx context.Context, tx *sql.Tx, data model.User) error
	GetUserCountWhereEmail(ctx context.Context, db *sql.DB, email string) (int, error)
}

type UserRepositoryImpl struct {
	Pg *sql.DB
}

func NewUserRepository(pg *sql.DB) UserRepository {
	return &UserRepositoryImpl{
		Pg: pg,
	}
}

func (r *UserRepositoryImpl) AddUserTx(ctx context.Context, tx *sql.Tx, data model.User) error {
	SQL := "INSERT INTO public.users (user_id, name, email, created_at, updated_at) VALUES ($1, $2, $3, $4, $5)"
	_, err := tx.ExecContext(ctx, SQL, data.UserId, data.Name, data.Email, data.CreatedAt, data.UpdatedAt)
	if err != nil {
		return err
	}
	return nil
}

func (r *UserRepositoryImpl) GetUserCountWhereEmail(ctx context.Context, db *sql.DB, email string) (int, error) {
	var total int

	SQL := "SELECT COUNT(*) FROM public.users WHERE email = $1"
	err := db.QueryRowContext(ctx, SQL, email).Scan(&total)
	if err != nil {
		return total, err
	}

	return total, nil
}
