package repository

import (
	"context"
	"database/sql"
	"errors"

	"github.com/Alfian57/belajar-golang/internal/config"
	"github.com/Alfian57/belajar-golang/internal/errs"
	"github.com/Alfian57/belajar-golang/internal/model"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type UserRepository struct {
	db *sqlx.DB
}

func NewUserRepository() *UserRepository {
	return &UserRepository{db: config.DB}
}

func (r *UserRepository) GetAll(ctx context.Context) ([]model.User, error) {
	users := []model.User{}
	query := "SELECT id, username, created_at, updated_at FROM users"

	err := r.db.SelectContext(ctx, &users, query)

	return users, err
}

func (r *UserRepository) Create(ctx context.Context, user *model.User) error {
	query := "INSERT INTO users(id, username, password) VALUES (?, ?, ?)"

	_, err := r.db.ExecContext(ctx, query, uuid.New().String(), user.Username, user.Password)

	return err
}

func (r *UserRepository) GetByID(ctx context.Context, id string) (model.User, error) {
	user := model.User{}
	query := "SELECT id, username, created_at, updated_at FROM users WHERE id = ?"

	err := r.db.GetContext(ctx, &user, query, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return user, errs.ErrDataNotFound
		}
		return user, err
	}

	return user, nil
}

func (r *UserRepository) Update(ctx context.Context, user *model.User) error {
	query := "UPDATE users SET username = ? WHERE id = ?"
	_, err := r.db.ExecContext(ctx, query, user.Username, user.ID.String())
	return err
}

func (r *UserRepository) Delete(ctx context.Context, id string) error {
	query := "DELETE FROM users WHERE id = ?"
	result, err := r.db.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if rowsAffected == 0 {
		return errs.ErrDataNotFound
	}

	return err
}
